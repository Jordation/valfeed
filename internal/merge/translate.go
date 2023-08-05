package merge

import (
	"fmt"

	"github.com/Jordation/jsonl/internal/pluck"
	"github.com/Jordation/jsonl/internal/types"
	riotTypes "github.com/Jordation/jsonl/provider/provider"
)

func TranslateMatchData(events []*riotTypes.Event) (*types.Map, error) {
	res := &types.Map{}

	teams := pluck.GetTeams(events)
	if len(teams) != 2 {
		return nil, fmt.Errorf("failed to fill teams")
	}
	res.Teams = teams

	players := pluck.GetPlayers(events)
	if len(players) != 10 {
		return nil, fmt.Errorf("failed to fill players")
	}
	res.Players = players

	res.Scores, res.Winner = pluck.GetScores(events)

	log.Info("result ", res)
	return res, nil
}
