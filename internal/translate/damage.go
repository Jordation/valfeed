package translate

import (
	"fmt"
	"strings"

	"github.com/Jordation/jsonl/internal/types"
	riotTypes "github.com/Jordation/jsonl/provider/types"
)

type DamageTranslator struct {
	PlayerMap   map[int]string
	WeaponMap   map[string]string
	OutputQueue chan<- *types.CombatEvent
}

func NewDamageTranslator(playerMap map[int]string, weaponMap map[string]string, outChan chan<- *types.CombatEvent) *DamageTranslator {
	return &DamageTranslator{
		PlayerMap:   playerMap,
		WeaponMap:   weaponMap,
		OutputQueue: outChan,
	}
}

func (t *DamageTranslator) HandleEvent(event *riotTypes.Event) {
	if event.DamageEvent == nil {
		return
	}
	if event.DamageEvent.KillEvent {
		t.handleKillEvent(event.DamageEvent, event.Metadata.SequenceNumber, event.Metadata.GameID.Value)
	} else {
		t.handleDamageEvent(event.DamageEvent, event.Metadata.SequenceNumber, event.Metadata.GameID.Value)
	}

}

func (t *DamageTranslator) handleKillEvent(event *riotTypes.DamageEvent, seqNum int, ID string) {
	weapon := ""
	if event.Weapon == nil {
		weapon = fmt.Sprintf("UNKNOWN_NO_WEAPON_ON_EVT_%d", seqNum)
	} else {
		weapon = t.WeaponMap[strings.ToLower(event.Weapon.Fallback.GUID)]
	}

	t.OutputQueue <- &types.CombatEvent{
		ID:          ID,
		Type:        types.Killed,
		SequenceNum: seqNum,
		Causer:      t.PlayerMap[event.CauserID.Value],
		Victim:      t.PlayerMap[event.VictimID.Value],
		DmgLoc:      damageTypes()[event.Location],
		DmgOnHit:    event.DamageDealt,
		RawDmg:      event.DamageAmount,
		Wallbang:    event.WallPen,
		Weapon:      weapon,
	}
}

func (t *DamageTranslator) handleDamageEvent(event *riotTypes.DamageEvent, seqNum int, ID string) {
	min := 1.5
	if event.DamageAmount < min || event.CauserID == nil {
		return
	}

	weapon := ""
	if event.Weapon == nil {
		weapon = fmt.Sprintf("UNKNOWN_NO_WEAPON_ON_EVT_%d", seqNum)
	} else {
		weapon = t.WeaponMap[strings.ToLower(event.Weapon.Fallback.GUID)]
	}
	t.OutputQueue <- &types.CombatEvent{
		ID:          ID,
		Type:        types.Shot,
		SequenceNum: seqNum,
		Causer:      t.PlayerMap[event.CauserID.Value],
		Victim:      t.PlayerMap[event.VictimID.Value],
		DmgLoc:      damageTypes()[event.Location],
		DmgOnHit:    event.DamageDealt,
		RawDmg:      event.DamageAmount,
		Wallbang:    event.WallPen,
		Weapon:      weapon,
	}
}

func damageTypes() map[string]string {
	return map[string]string{
		"HEAD":    "headshot",
		"BODY":    "bodyshot",
		"LEG":     "legshot",
		"GENERAL": "aoe damage",
	}
}
