package merge

import (
	"github.com/Jordation/jsonl/internal/persistance"
	"github.com/Jordation/jsonl/internal/types"
)

// merge is for merging events of our internal type into structured, relational data representing the gamestate.

type MergeManager struct {
	EventStream chan any
	Persistance *persistance.Persistance
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

func NewMerger(p *persistance.Persistance) *MergeManager {
	m := &MergeManager{
		Persistance: p,
		EventStream: make(chan any),
		Matches:     map[string]*matchMerger{},
	}
	m.start()
	return m
}

func (m *MergeManager) Receive(evt any, gameID string) {
	if _, ok := m.Matches[gameID]; !ok {
		m.Matches[gameID], _ = m.newMatchManager(gameID)
	}
	m.Matches[gameID].HandleEvent(evt)
}

func (m *MergeManager) start() {
	go func() {
		for evt := range m.EventStream {
			_ = evt
		}
	}()
}

func (m *MergeManager) newMatchManager(gameID string) (*matchMerger, error) {
	cfg, err := m.Persistance.GetGameConfig(gameID)
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
		MatchState:    &types.MatchState{},
	}, nil
}

func (mm *matchMerger) HandleEvent(incoming any) {
	switch event := incoming.(type) {
	case *types.CombatEvent:
		mm.handleCombatEvent(event)
	case *types.RoundEvent:
		mm.handleRoundEvent(event)
	}
}

func (mm *matchMerger) handleCombatEvent(evt *types.CombatEvent) {

}

func (mm *matchMerger) handleRoundEvent(evt *types.RoundEvent) {

}
