package translate

import (
	"github.com/Jordation/jsonl/internal/distributor"
	"github.com/Jordation/jsonl/internal/types"
	riotTypes "github.com/Jordation/jsonl/provider/types"
	"github.com/Jordation/jsonl/utils"
)

type Translatorer interface {
	HandleEvent(event *riotTypes.Event)
}

type TranslationManager struct {
	// game id to translator
	Translators    map[string]*matchTranslator
	IncomingEvents chan *riotTypes.Event
	DamageReturn   chan *types.CombatEvent
	RoundReturn    chan *types.RoundEvent
	persistence    *persistence.persistence
}

type matchTranslator struct {
	EventStream chan *riotTypes.Event
	Translators map[string]Translatorer
}

func NewManager(p *persistence.persistence, d *distributor.Distributor) *TranslationManager {
	t := &TranslationManager{
		persistence:    p,
		Translators:    map[string]*matchTranslator{},
		DamageReturn:   d.IncDmgEvents,
		RoundReturn:    d.IncRndEvents,
		IncomingEvents: make(chan *riotTypes.Event),
	}
	t.start()
	return t
}

func (m *TranslationManager) Receive(event *riotTypes.Event) {
	m.IncomingEvents <- event
}

func (m *TranslationManager) start() {
	go func() {
		for event := range m.IncomingEvents {
			if _, ok := m.Translators[event.Metadata.GameID.Value]; !ok && event.Configuration != nil {
				m.Translators[event.Metadata.GameID.Value] = m.newTranslator(event.Configuration, event.Metadata.GameID.Value)
			} else if ok {
				m.Translators[event.Metadata.GameID.Value].EventStream <- event
			}
		}

	}()
}

// a translator is prepared to translate ONE active game
// seq1 on each stream has a config
func (m *TranslationManager) newTranslator(cfg *riotTypes.GameConfig, ID string) *matchTranslator {
	playerMap, playerToTeamMap, sideStartMap := cfg.GetMappings()
	wepMaps := utils.GetWeaponMappings()

	if err := m.persistence.StoreGameConfig(cfg, ID); err != nil {
		panic(err)
	}

	translator := &matchTranslator{
		EventStream: make(chan *riotTypes.Event),
		Translators: m.getTranslators(playerMap, wepMaps, playerToTeamMap, sideStartMap),
	}

	translator.start()
	return translator
}

func (t *matchTranslator) start() {
	go func() {
		for event := range t.EventStream {
			for _, translator := range t.Translators {
				translator.HandleEvent(event)
			}
		}
	}()
}

func (m *TranslationManager) getTranslators(playerMap map[int]string, weapons map[string]string, playerToTeams map[int]int, sideStart map[int]string) map[string]Translatorer {
	res := map[string]Translatorer{}
	res["damage"] = NewDamageTranslator(playerMap, weapons, m.DamageReturn)
	res["round"] = NewRoundTranslator(m.RoundReturn)

	return res
}

func TranslateConfig(cfg *riotTypes.GameConfig) (
	map[int]string,
	map[int]int,
	map[string]int,
	*types.MatchState,
) {
	playerMap := make(map[int]string, 10)
	playerToTeamMap := make(map[int]int, 10)
	sideStartMap := make(map[string]int, 2)
	for _, player := range cfg.Players {
		playerMap[player.PlayerID.Value] = player.DisplayName
	}
	for _, team := range cfg.Teams {
		for _, player := range team.PlayersInTeam {
			playerToTeamMap[player.Value] = team.TeamID.Value
		}
	}
	sideStartMap["atk"] = cfg.SpikeMode.AttackingTeam.Value
	sideStartMap["def"] = cfg.SpikeMode.DefendingTeam.Value

	return playerMap, playerToTeamMap, sideStartMap, &types.MatchState{}
}
