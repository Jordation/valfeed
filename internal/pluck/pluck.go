package pluck

import (
	"strings"

	"github.com/Jordation/jsonl/internal/types"
	riotTypes "github.com/Jordation/jsonl/provider/types"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func init() {
	log = logrus.New()
}

func GetTeams(events []*riotTypes.Event) map[int]*types.Team {
	res := map[int]*types.Team{}
	for _, event := range events {
		// we can return once we have both teams filled, no need to check all events
		if len(res) == 2 {
			break
		}

		if event.Configuration != nil && event.Configuration.Teams != nil {
			for _, team := range event.Configuration.Teams {
				t := types.Team{}
				for _, player := range team.PlayersInTeam {
					t.Players = append(t.Players, player.Value)
				}
				t.StartSide = teamStartSide(team.Name)
				res[team.TeamID.Value] = &t
			}
		}
	}
	return res
}

func GetPlayers(events []*riotTypes.Event) map[int]string {
	res := map[int]string{}
	for _, event := range events {
		// we can return once we have all the players, no need to check all events
		if len(res) == 10 {
			break
		}

		if event.Configuration != nil && event.Configuration.Players != nil {
			for _, player := range event.Configuration.Players {
				res[player.PlayerID.Value] = player.DisplayName
			}
		}
	}
	return res
}

func GetScores(events []*riotTypes.Event) (map[int]int, int) {
	res := map[int]int{}
	for _, event := range events {
		if event.RoundStarted != nil {
			roundsPlayed := len(event.RoundStarted.SpikeMode.CompletedRounds)
			if roundsPlayed > 0 {
				lastRound := event.RoundStarted.SpikeMode.CompletedRounds[roundsPlayed-1]
				res[lastRound.WinningTeam.Value]++
			}
		}
	}

	var score, winner int

	for teamId, roundsWon := range res {
		if roundsWon > score {
			score = roundsWon
			winner = teamId
		}

	}
	return res, winner
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
