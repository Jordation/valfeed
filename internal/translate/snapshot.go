package translate

import (
	"github.com/Jordation/jsonl/internal/types"
	riotTypes "github.com/Jordation/jsonl/provider/types"
)

type SnapshotTranslator struct {
	PlayerMap   map[int]string
	OutputQueue chan<- *types.PlayerPositionEvents
}

func NewPositionTranslator(playerMap map[int]string, outChan chan<- *types.PlayerPositionEvents) *SnapshotTranslator {
	return &SnapshotTranslator{
		PlayerMap:   playerMap,
		OutputQueue: outChan,
	}
}

func (t *SnapshotTranslator) HandleEvent(event *riotTypes.Event) {
	if event.Snapshot == nil {
		return
	}
	result := &types.PlayerPositionEvents{
		Events:         []*types.PlayerPosition{},
		GameID:         event.Metadata.GameID.Value,
		SequenceNumber: event.Metadata.SequenceNumber,
	}
	for _, player := range event.Snapshot.Players {
		result.Events = append(result.Events, &types.PlayerPosition{
			PlayerID: player.PlayerID.Value,
			GameID:   event.Metadata.GameID.Value,
		})
	}
}
