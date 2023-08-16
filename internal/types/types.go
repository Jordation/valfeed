package types

var DamageEventMap = map[DamageEventType]string{
	Died:        "died",
	Killed:      "killed",
	TookDamage:  "took damage",
	DealtDamage: "dealt damage",
}

type DamageEventType int

const (
	Died DamageEventType = iota
	Killed
	TookDamage
	DealtDamage
)

type DamageEvent struct {
	Type        DamageEventType
	Causer      string
	Victim      string
	DmgLoc      string
	DmgDone     float64
	DetailStr   string
	SequenceNum int
	Weapon      string
}

type RoundEvent struct {
	RoundNumber int
	SeqInfo     *RoundData
	Winner      int
	WinReason   string
}

type RoundData struct {
	Start int
	Play  int
	End   int
}
type RoundMeta struct {
	RoundStart int
	PlayStart  int
	End        int
}

type MatchState struct {
}
