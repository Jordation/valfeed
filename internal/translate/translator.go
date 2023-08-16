package translate

import (
	"github.com/Jordation/jsonl/internal/types"
	riotTypes "github.com/Jordation/jsonl/provider/types"
	"github.com/Jordation/jsonl/utils"
	"github.com/sirupsen/logrus"
)

type Translatorer interface {
	HandleEvent(event *riotTypes.Event)
}

// TODO impl manager
type manager struct {
	// game id to translator
	Translators  map[string]*MatchTranslator
	Stream       chan *riotTypes.Event
	DamageReturn chan *types.DamageEvent
	RoundReturn  chan *types.RoundEvent
}

type MatchTranslator struct {
	EventStream chan *riotTypes.Event
	Translators map[string]Translatorer
}

func (m *manager) Start() {
	for {
		select {
		case event := <-m.Stream:
			if translator, ok := m.Translators[event.Metadata.GameID.Value]; ok {
				translator.EventStream <- event
			} else {
				if event.Configuration != nil {
					translator := m.newTranslator(event.Configuration)
					translator.start()
					m.Translators[event.Metadata.GameID.Value] = translator
				}
			}
		}
	}
}

// a translator is prepared to translate ONE active game
// seq1 on each stream has a config
func (m *manager) newTranslator(cfg *riotTypes.GameConfig) *MatchTranslator {
	playerMap := make(map[int]string, 10)
	playerToTeamMap := make(map[int]int, 10)
	sideStartMap := make(map[int]string, 2)

	for _, player := range cfg.Players {
		playerMap[player.PlayerID.Value] = player.DisplayName
	}

	for _, team := range cfg.Teams {
		for _, player := range team.PlayersInTeam {
			playerToTeamMap[player.Value] = team.TeamID.Value
		}
	}

	sideStartMap[cfg.SpikeMode.AttackingTeam.Value] = "atk"
	sideStartMap[cfg.SpikeMode.DefendingTeam.Value] = "def"
	wepMaps := utils.GetWeaponMappings()
	//abilityMaps := utils.GetAgentAbilityMappings()

	return &MatchTranslator{
		EventStream: make(chan *riotTypes.Event),
		Translators: m.getTranslators(playerMap, wepMaps, playerToTeamMap, sideStartMap),
	}
}

func (t *MatchTranslator) start() {
	for {
		select {
		case event := <-t.EventStream:
			for name, translator := range t.Translators {
				logrus.Infof("%v_handling_event_%v", name, event.Metadata.SequenceNumber)
				translator.HandleEvent(event)
			}
		}
	}
}

func (m *manager) getTranslators(playerMap map[int]string, weapons map[string]string, playerToTeams map[int]int, sideStart map[int]string) map[string]Translatorer {
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
