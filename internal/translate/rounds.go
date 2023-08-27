package translate

import (
	"github.com/Jordation/jsonl/internal/types"
	riotTypes "github.com/Jordation/jsonl/provider/types"
)

const (
	R_START   = "ROUND_STARTING"
	R_END     = "ROUND_ENDING"
	R_IN      = "IN_ROUND"
	R_BETWEEN = "BETWEEN_ROUNDS"
)

type RoundTranslator struct {
	OpenRound   *openRound
	InputQueue  <-chan *riotTypes.Event
	OutputQueue chan<- *types.RoundEvent
}

type openRound struct {
	round *types.RoundData
	rNum  int
	phase string
}

func NewRoundTranslator(outChan chan<- *types.RoundEvent) *RoundTranslator {
	return &RoundTranslator{
		InputQueue:  make(chan *riotTypes.Event),
		OutputQueue: outChan,
	}
}

func (t *RoundTranslator) HandleEvent(event *riotTypes.Event) {
	if event.RoundDecided != nil {
		t.translateRoundDecided(event.RoundDecided, event.Metadata.GameID.Value)
	}
	if event.GamePhase != nil {
		t.translateGamePhase(event.GamePhase, event.Metadata.SequenceNumber)
	}
}

func (t *RoundTranslator) translateRoundDecided(event *riotTypes.RoundDecided, ID string) {
	t.OutputQueue <- &types.RoundEvent{
		GameID:      ID,
		SeqInfo:     t.OpenRound.round,
		RoundNumber: t.OpenRound.rNum,
		WinReason:   event.Result.SpikeModeResult.Cause,
		Winner:      event.Result.WinningTeam.Value,
	}
}

func (t *RoundTranslator) translateGamePhase(event *riotTypes.GamePhase, seq int) {
	if t.OpenRound == nil {
		t.OpenRound = &openRound{
			rNum:  0,
			phase: R_BETWEEN,
		}
	}

	switch event.Phase {
	case R_START:
		t.OpenRound.round = &types.RoundData{
			Start: seq,
		}
		t.OpenRound.rNum++
		t.OpenRound.phase = R_START

	case R_END:
		t.OpenRound.phase = R_END
		t.OpenRound.round.End = seq

	case R_IN:
		t.OpenRound.round.Play = seq

	}

}
