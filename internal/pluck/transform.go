package pluck

import (
	riotTypes "github.com/Jordation/jsonl/provider/types"
	"github.com/jinzhu/copier"
)

// transform groups, filters or otherwise modifies a series of events

func GroupInventoryTransactionsByPlayer(events []*riotTypes.Event) map[int][]*riotTypes.InventoryTransaction {
	res := map[int][]*riotTypes.InventoryTransaction{}
	for _, event := range events {
		res[event.InventoryTransaction.Player.Value] = append(res[event.InventoryTransaction.Player.Value], event.InventoryTransaction)
	}
	return res
}

func GetPlayerMovements(events []*riotTypes.Event) map[int][]*riotTypes.Position {
	res := map[int][]*riotTypes.Position{}

	for _, event := range events {
		for _, player := range event.GetSnapshot().GetPlayers() {
			id := player.GetPlayerID().Value
			if id == 0 {
				continue
			}

			for _, ts := range player.GetTimeseries() {
				if ts.Position != nil {
					res[id] = append(res[id], ts.Position)
				}
			}
		}
	}

	return res
}

func SplitDeltasByPlayer(events []*riotTypes.Event) map[int][]*riotTypes.Event {
	res := map[int][]*riotTypes.Event{}

	for _, event := range events {
		for _, player := range event.GetSnapshot().GetPlayers() {
			temp := &riotTypes.Event{}
			copier.Copy(temp, event)
			temp.Snapshot.Players = []*riotTypes.PlayerInGame{player}
			id := player.GetPlayerID().Value
			if id == 0 {
				continue
			}

			res[id] = append(res[id], temp)
		}
	}

	return res
}

func TransformPlayerPositionWithMapData(positions []*riotTypes.Position, mapdata *riotTypes.APIMap) {
	for _, pos := range positions {
		pos.TransformWithMapData(mapdata.XMultiplier, mapdata.XScalarToAdd, mapdata.YMultiplier, mapdata.YScalarToAdd)
	}
}

func GetRoundTimings(events []*riotTypes.Event) []int {
	res := []int{}

	for _, event := range events {
		if event.RoundStarted != nil {
			res = append(res, event.Metadata.SequenceNumber)
		}
	}

	return res
}

func RoundNumberFromSequence(meta *riotTypes.Metadata, roundChanges []int) int {
	for rNum, breakpoint := range roundChanges {
		if meta.SequenceNumber < breakpoint {
			return rNum
		}
	}

	return 0
}

func GroupDamageEvents(events []*riotTypes.Event) map[int][]*riotTypes.DamageEvent {
	res := map[int][]*riotTypes.DamageEvent{}

	for _, event := range events {
		if event.DamageEvent != nil {
			res[event.DamageEvent.CauserID.Value] = append(res[event.DamageEvent.CauserID.Value], event.DamageEvent)
		}
	}

	return res
}

func ReadDamageEvents(events []*riotTypes.DamageEvent, playerMappings map[int]string) {
	for _, event := range events {
		if event.KillEvent {
			log.Infof("%v -> KILLED -> %v -> SHOT %v", playerMappings[event.CauserID.Value], playerMappings[event.VictimID.Value], event.Location)
		} else {
			log.Infof("%v -> SHOT -> %v -> FOR -> %v DMG", playerMappings[event.CauserID.Value], playerMappings[event.VictimID.Value], event.DamageAmount)
		}
	}
}
