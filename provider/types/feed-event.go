package types

const (
	ROUND_START = "ROUND_STARTING"
	PLAY_START  = "IN_ROUND"
	ROUND_END   = "ROUND_ENDING"
)

type Event struct {
	Configuration        *GameConfig           `json:"configuration,omitempty"`
	DamageEvent          *DamageEvent          `json:"damageEvent,omitempty"`
	InventoryTransaction *InventoryTransaction `json:"inventoryTransaction,omitempty"`
	GamePhase            *GamePhase            `json:"gamePhase,omitempty"`
	RoundStarted         *RoundStarted         `json:"roundStarted,omitempty"`
	RoundDecided         *RoundDecided         `json:"roundDecided,omitempty"`
	Snapshot             *Snapshot             `json:"snapshot,omitempty"`
	Metadata             *Metadata             `json:"metadata,omitempty"`
}

type GamePhase struct {
	Phase       string `json:"phase,omitempty"`
	RoundNumber int    `json:"roundNumber,omitempty"`
}

type Result struct {
	RoundNumber     int              `json:"roundNumber,omitempty"`
	SpikeModeResult *SpikeModeResult `json:"spikeModeResult,omitempty"`
	WinningTeam     *WinningTeam     `json:"winningTeam,omitempty"`
}

type RoundDecided struct {
	Result Result `json:"result,omitempty"`
}

type InventoryTransaction struct {
	TransactionType string `json:"transactionType,omitempty"`

	Ability *Ability  `json:"ability,omitempty"`
	Player  *PlayerID `json:"player,omitempty"`
	Weapon  *Weapon   `json:"weapon,omitempty"`
}
type Metadata struct {
	GameVersion    string `json:"gameVersion,omitempty"`
	Playback       int    `json:"playback,omitempty"`
	SequenceNumber int    `json:"sequenceNumber,omitempty"`
	Stage          int    `json:"stage,omitempty"`
	WallTime       string `json:"wallTime,omitempty"`

	EventTime  *EventTime  `json:"eventTime,omitempty"`
	GameID     *GameID     `json:"gameId,omitempty"`
	ServerInfo *ServerInfo `json:"serverInfo,omitempty"`
}
type Snapshot struct {
	Players []*PlayerInGame `json:"players,omitempty"`
}
type RoundStarted struct {
	RoundNumber int        `json:"roundNumber,omitempty"`
	SpikeMode   *SpikeMode `json:"spikeMode,omitempty"`
}
type GameConfig struct {
	SelectedMap *SelectedMap  `json:"selectedMap,omitempty"`
	SpikeMode   *SpikeMode    `json:"spikeMode,omitempty"`
	Players     []*PlayerData `json:"players,omitempty"`
	Teams       []*Team       `json:"teams,omitempty"`
}
type DamageEvent struct {
	DamageAmount float64 `json:"damageAmount,omitempty"`
	DamageDealt  float64 `json:"damageDealt,omitempty"`
	KillEvent    bool    `json:"killEvent,omitempty"`
	Location     string  `json:"location,omitempty"`
	WallPen      bool    `json:"wallPen,omitempty"`

	Weapon   *Weapon   `json:"weapon,omitempty"`
	CauserID *PlayerID `json:"causerId,omitempty"`
	VictimID *PlayerID `json:"victimId,omitempty"`
}
type SelectedAgent struct {
	Fallback *Fallback `json:"fallback,omitempty"`
	Type     string    `json:"type,omitempty"`
}
type PlayerData struct {
	DisplayName string `json:"displayName,omitempty"`
	TagLine     string `json:"tagLine,omitempty"`
	Type        string `json:"type,omitempty"`

	AccountID     *AccountID     `json:"accountId,omitempty"`
	PlayerID      *PlayerID      `json:"playerId,omitempty"`
	SelectedAgent *SelectedAgent `json:"selectedAgent,omitempty"`
}
type SelectedMap struct {
	ID       string    `json:"id,omitempty"`
	Fallback *Fallback `json:"fallback,omitempty"`
}
type TeamID struct {
	Value int `json:"value,omitempty"`
}
type Team struct {
	Name          string      `json:"name,omitempty"`
	PlayersInTeam []*PlayerID `json:"playersInTeam,omitempty"`
	TeamID        *TeamID     `json:"teamId,omitempty"`
}

type Fallback struct {
	DisplayName string `json:"displayName,omitempty"`
	GUID        string `json:"guid,omitempty"`

	InventorySlot *InventorySlot `json:"inventorySlot,omitempty"`
}

type Ability struct {
	Type     string    `json:"type,omitempty"`
	Fallback *Fallback `json:"fallback,omitempty"`
}
type Weapon struct {
	Type     string    `json:"type,omitempty"`
	Fallback *Fallback `json:"fallback,omitempty"`
}

type Abilities struct {
	BaseCharges      int `json:"baseCharges,omitempty"`
	MaxCharges       int `json:"maxCharges,omitempty"`
	TemporaryCharges int `json:"temporaryCharges,omitempty"`

	Ability *Ability `json:"ability,omitempty"`
}

type Skin struct {
	ChromaFallback *ChromaFallback `json:"chromaFallback,omitempty"`
	LevelFallback  *LevelFallback  `json:"levelFallback,omitempty"`
	SkinFallback   *SkinFallback   `json:"skinFallback,omitempty"`
}

type EquippedItem struct {
	DisplayName string `json:"displayName,omitempty"`
	GUID        string `json:"guid,omitempty"`

	Skin *Skin `json:"skin,omitempty"`
	Slot *Slot `json:"slot,omitempty"`
}

type AliveState struct {
	Armor  float64 `json:"armor,omitempty"`
	Health float64 `json:"health,omitempty"`

	EquippedItem *EquippedItem `json:"equippedItem,omitempty"`
	Position     *Position     `json:"position,omitempty"`
	Velocity     *Velocity     `json:"velocity,omitempty"`
	ViewVector   *ViewVector   `json:"viewVector,omitempty"`
}

type Inventory struct {
	DisplayName string `json:"displayName,omitempty"`
	GUID        string `json:"guid,omitempty"`

	Ammunition *Ammunition `json:"ammunition,omitempty"`
	Skin       *Skin       `json:"skin,omitempty"`
	Slot       *Slot       `json:"slot,omitempty"`
}

type Timeseries struct {
	Position   *Position   `json:"position,omitempty"`
	Timestamp  *Timestamp  `json:"timestamp,omitempty"`
	Velocity   *Velocity   `json:"velocity,omitempty"`
	ViewVector *ViewVector `json:"viewVector,omitempty"`
}

type PlayerInGame struct {
	Assists int `json:"assists,omitempty"`
	Deaths  int `json:"deaths,omitempty"`
	Kills   int `json:"kills,omitempty"`
	Money   int `json:"money,omitempty"`

	Abilities  []*APIAbility `json:"abilities,omitempty"`
	Inventory  []*Inventory  `json:"inventory,omitempty"`
	Timeseries []*Timeseries `json:"timeseries,omitempty"`
	AliveState *AliveState   `json:"aliveState,omitempty"`
	PlayerID   *PlayerID     `json:"playerId,omitempty"`
	Scores     *Scores       `json:"scores,omitempty"`
}

type Scores struct {
	CombatScore *CombatScore `json:"combatScore,omitempty"`
}
type SpikeModeResult struct {
	Cause string `json:"cause,omitempty"`

	AttackingTeam *AttackingTeam `json:"attackingTeam,omitempty"`
	DefendingTeam *DefendingTeam `json:"defendingTeam,omitempty"`
}

type CompletedRound struct {
	RoundNumber int `json:"roundNumber,omitempty"`

	SpikeModeResult *SpikeModeResult `json:"spikeModeResult,omitempty"`
	WinningTeam     *WinningTeam     `json:"winningTeam,omitempty"`
}

type SpikeMode struct {
	CurrentRound int `json:"currentRound,omitempty"`
	RoundsToWin  int `json:"roundsToWin,omitempty"`

	CompletedRounds []*CompletedRound `json:"completedRounds,omitempty"`
	AttackingTeam   *AttackingTeam    `json:"attackingTeam,omitempty"`
	DefendingTeam   *DefendingTeam    `json:"defendingTeam,omitempty"`
}

type WinningTeam struct {
	Value int `json:"value,omitempty"`
}
type Slot struct {
	Slot string `json:"slot,omitempty"`
}
type Ammunition struct {
	InMagazine int `json:"inMagazine,omitempty"`
	InReserve  int `json:"inReserve,omitempty"`
}
type PlayerID struct {
	Value int `json:"value,omitempty"`
}
type CombatScore struct {
	RoundScore int `json:"roundScore,omitempty"`
	TotalScore int `json:"totalScore,omitempty"`
}
type Timestamp struct {
	IncludedPauses string `json:"includedPauses,omitempty"`
	OmittingPauses string `json:"omittingPauses,omitempty"`
}
type AttackingTeam struct {
	Value int `json:"value,omitempty"`
}
type DefendingTeam struct {
	Value int `json:"value,omitempty"`
}
type AccountID struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}
type EventTime struct {
	IncludedPauses string `json:"includedPauses,omitempty"`
	OmittingPauses string `json:"omittingPauses,omitempty"`
}
type GameID struct {
	Value string `json:"value,omitempty"`
}
type ServerInfo struct {
	ProcessID   string `json:"processId,omitempty"`
	Rfc190Scope string `json:"rfc190Scope,omitempty"`
}
type ChromaFallback struct {
	DisplayName string `json:"displayName,omitempty"`
	GUID        string `json:"guid,omitempty"`
}
type LevelFallback struct {
	DisplayName string `json:"displayName,omitempty"`
	GUID        string `json:"guid,omitempty"`
}
type SkinFallback struct {
	DisplayName string `json:"displayName,omitempty"`
	GUID        string `json:"guid,omitempty"`
}
type InventorySlot struct {
	Slot string `json:"slot,omitempty"`
}
type Position struct {
	X float64 `json:"x,omitempty"`
	Y float64 `json:"y,omitempty"`
	Z float64 `json:"z,omitempty"`
}
type Velocity struct {
	X float64 `json:"x,omitempty"`
	Y float64 `json:"y,omitempty"`
	Z float64 `json:"z,omitempty"`
}
type ViewVector struct {
	X float64 `json:"x,omitempty"`
	Y float64 `json:"y,omitempty"`
	Z float64 `json:"z,omitempty"`
}

func (d *Event) GetSnapshot() *Snapshot {
	if d.Snapshot != nil {
		return d.Snapshot
	}
	return &Snapshot{}
}

func (s *Snapshot) GetPlayers() []*PlayerInGame {
	if s.Players != nil {
		return s.Players
	}
	return nil
}

func (p *PlayerInGame) GetTimeseries() []*Timeseries {
	if p.Timeseries != nil {
		return p.Timeseries
	}
	return nil
}

func (p *PlayerInGame) GetPlayerID() *PlayerID {
	if p.PlayerID != nil {
		return p.PlayerID
	}
	return &PlayerID{}
}

func (gc *GameConfig) GetMappings() (map[int]string, map[int]int, map[int]string) {
	playerMap := make(map[int]string, 10)
	playerToTeamMap := make(map[int]int, 10)
	sideStartMap := make(map[int]string, 2)

	for _, player := range gc.Players {
		playerMap[player.PlayerID.Value] = player.DisplayName
	}

	for _, team := range gc.Teams {
		for _, player := range team.PlayersInTeam {
			playerToTeamMap[player.Value] = team.TeamID.Value
		}
	}

	sideStartMap[gc.SpikeMode.AttackingTeam.Value] = "atk"
	sideStartMap[gc.SpikeMode.DefendingTeam.Value] = "def"
	return playerMap, playerToTeamMap, sideStartMap
}
