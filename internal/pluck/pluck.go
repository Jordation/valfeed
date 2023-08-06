package pluck

import (
	"github.com/Jordation/jsonl/internal/types"
	riotTypes "github.com/Jordation/jsonl/provider/types"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func init() {
	log = logrus.New()
}

func GetScores(events []*riotTypes.Event) (map[int]int, int) {
	res := map[int]int{}
	for _, event := range events {
		if event.RoundDecided != nil {
			res[event.RoundDecided.Result.WinningTeam.Value]++
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

func InitialiseRounds(events []*riotTypes.Event) []*types.Round {
	res := []*types.Round{}
	roundsRecorded := 0
	for _, event := range events {
		if event.GamePhase != nil {
			switch event.GamePhase.Phase {
			case riotTypes.ROUND_START:
				res = append(res, &types.Round{Meta: &types.RoundMeta{
					RoundStart: event.Metadata.SequenceNumber,
				}})

			case riotTypes.PLAY_START:
				res[roundsRecorded].Meta.PlayStart = event.Metadata.SequenceNumber

			case riotTypes.ROUND_END:
				res[roundsRecorded].Meta.End = event.Metadata.SequenceNumber
				roundsRecorded++
			}
		}
	}

	return res
}
