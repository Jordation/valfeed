package merge

import (
	"github.com/Jordation/jsonl/internal/types"
)

// merge is for merging events of our internal type into structured, relational data representing the gamestate.

/*
	not sure how i feel about this code
	i think it's better if the translators produce simplier outputs with the relationship links
	to be stored persistently

*/

type MergeManager struct {
	EventStream chan *types.TranslatedEvent
	persistence *persistence.persistence
	Matches     map[string]*matchMerger
}

type matchMerger struct {
	ID            string
	PlayerMap     map[int]string
	PlayerTeamMap map[int]int
	SideStartMap  map[int]string

	InEvents   chan any
	MatchState *types.MatchState
}

func NewMerger(p *persistence.persistence) *MergeManager {
	m := &MergeManager{
		persistence: p,
		EventStream: make(chan *types.TranslatedEvent),
		Matches:     map[string]*matchMerger{},
	}
	m.start()
	return m
}

func (m *MergeManager) start() {
	go func() {
		for evt := range m.EventStream {
			if _, ok := m.Matches[evt.ID]; !ok {
				newManager, err := m.newMatchManager(evt.ID)
				if err != nil {
					panic(err)
				}
				m.Matches[evt.ID] = newManager
			}
			m.Matches[evt.ID].HandleEvent(evt)
		}
	}()
}

func (m *MergeManager) newMatchManager(gameID string) (*matchMerger, error) {
	cfg, err := m.persistence.GetGameConfig(gameID)
	if err != nil {
		return nil, err
	}

	playerMap, playersToTeam, sideStart := cfg.GetMappings()

	return &matchMerger{
		ID:            gameID,
		PlayerMap:     playerMap,
		PlayerTeamMap: playersToTeam,
		SideStartMap:  sideStart,
		InEvents:      make(chan any),
	}, nil
}

func (mm *matchMerger) HandleEvent(incoming *types.TranslatedEvent) {
	if mm.MatchState == nil {
		mm.setInitialMatchState()
	}
	switch event := incoming.Event.(type) {
	case *types.CombatEvent:
		mm.handleCombatEvent(event)
	case *types.RoundEvent:
		mm.handleRoundEvent(event)
	}
}

func (mm *matchMerger) handleCombatEvent(evt *types.CombatEvent) {
	targetIndex := len(mm.MatchState.Rounds) - 1
	mm.MatchState.Rounds[targetIndex].Events = append(mm.MatchState.Rounds[targetIndex].Events, evt)
}

func (mm *matchMerger) handleRoundEvent(evt *types.RoundEvent) {
	targetIndex := len(mm.MatchState.Rounds)
	mm.MatchState.Rounds = append(mm.MatchState.Rounds, mm.newRoundState())

	mm.MatchState.Rounds[targetIndex].RoundNum = targetIndex + 1
	mm.MatchState.Rounds[targetIndex].Finished = true
	mm.MatchState.Rounds[targetIndex].WinReason = evt.WinReason
	mm.MatchState.Rounds[targetIndex].Winner = evt.Winner
}

func (mm *matchMerger) setInitialMatchState() {
	mm.MatchState = &types.MatchState{
		Rounds: []*types.RoundState{
			mm.newRoundState(),
		},
	}
}

func (mm *matchMerger) newRoundState() *types.RoundState {
	return &types.RoundState{
		Players: mm.newPlayerStates(),
		Events:  []interface{}{},
	}
}

func (mm *matchMerger) newPlayerStates() []*types.PlayerState {
	res := []*types.PlayerState{}
	for id, name := range mm.PlayerMap {
		res = append(res, &types.PlayerState{
			ID:   id,
			Name: name,
			Team: mm.PlayerTeamMap[id],
		})
	}
	return res
}
