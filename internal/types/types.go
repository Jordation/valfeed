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
	GameID      string
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

type PlayerPositionEvents struct {
	GameID         string
	SequenceNumber int
	Events         []*PlayerPosition
}

type PlayerPosition struct {
	PlayerID int
	GameID   string
	X        float64
	Y        float64
	Z        float64
}

type MatchState struct {
	Rounds []*RoundState
}

type RoundState struct {
	RoundNum  int
	Winner    int
	Finished  bool
	WinReason string
	Players   []*PlayerState
	Events    []interface{}
}

type PlayerState struct {
	ID     int
	Name   string
	Team   int
	Kills  int
	Died   bool
	Weapon string
}

type TranslatedEvent struct {
	Event any
	ID    string
}
