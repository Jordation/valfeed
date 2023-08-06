package translation

import (
	"strings"

	"github.com/Jordation/jsonl/internal/types"
	riotTypes "github.com/Jordation/jsonl/provider/types"
	"github.com/sirupsen/logrus"
)

// The pipeline asseses the entities present on the event and sends them off to be translated
func TranslatorPipeline(event *riotTypes.Event) *types.Map {
	mapDelta := &types.Map{}
	if event.Configuration != nil {
		translateGameConfig(event.Configuration, mapDelta)
	}

	if event.InventoryTransaction != nil {
		logrus.Info("inv trans handling not impl.")
	}

	if event.Metadata != nil {
		logrus.Info("metadata handling not impl.")
	}

	if event.RoundDecided != nil {
		translateRoundDecided(event.RoundDecided, mapDelta)
	}

	if event.RoundStarted != nil {
		logrus.Info("round start handling not impl.")
	}

	if event.GamePhase != nil {
		logrus.Info("game phase handling not impl.")
	}

	if event.DamageEvent != nil {
		logrus.Info("damage event handling not impl.")
	}

	if event.Snapshot != nil {
		logrus.Info("snapshot handling not impl.")
	}

	return mapDelta
}

func translateGameConfig(cfg *riotTypes.GameConfig, m *types.Map) {
	teams := map[int]*types.Team{}
	for _, team := range cfg.Teams {
		t := types.Team{}

		for _, player := range team.PlayersInTeam {
			t.Players = append(t.Players, player.Value)
		}

		t.StartSide = teamStartSide(team.Name)
		teams[team.TeamID.Value] = &t
	}
	m.Teams = teams

	players := map[int]string{}
	for _, player := range cfg.Players {
		players[player.PlayerID.Value] = player.DisplayName
	}
	m.Players = players
}

func translateRoundDecided(rd *riotTypes.RoundDecided, m *types.Map) {
	// todo
}

func translateGamePhase(phase *riotTypes.GamePhase, m *types.Map) {
	switch phase.Phase {
	case riotTypes.ROUND_START:
		m.Rounds = []*types.Round{{Meta: &types.RoundMeta{RoundStart: phase.RoundNumber}}}
	case riotTypes.PLAY_START:
	case riotTypes.ROUND_END:
	}
}

func teamStartSide(col string) string {
	switch strings.ToLower(col) {
	case "blue":
		return "def"
	case "red":
		return "atk"
	default:
		return ""
	}
}
