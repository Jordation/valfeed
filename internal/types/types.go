package types

var DamageEventMap = map[DamageEventType]string{
	Killed: "killed",
	Shot:   "shot",
}

type DamageEventType int

const (
	Killed DamageEventType = iota
	Shot
)

type CombatEvent struct {
	ID          string
	Type        DamageEventType
	Causer      string
	Victim      string
	DmgLoc      string
	DmgOnHit    float64
	RawDmg      float64
	Wallbang    bool
	DetailStr   string
	SequenceNum int
	Weapon      string
}

type RoundEvent struct {
	ID          string
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
	Rounds []*RoundState
}

type RoundState struct {
	Winner    int
	Finished  bool
	WinReason string
	Players   []*PlayerState
	Events    []interface{}
}

type PlayerState struct {
	ID     int
	Kills  int
	Died   bool
	Weapon string
}
