package types

type Map struct {
	Winner  int
	Players map[int]string
	Scores  map[int]int
	Teams   map[int]*Team
	Rounds  []*Round
}

type Team struct {
	Players   []int
	StartSide string
}

type Round struct {
	Winner     string
	TeamMoney  map[int]int
	ScoreBoard map[int]*Score
	History    map[int]RoundHistory
	Events     []string
	Meta       *RoundMeta
}

type RoundMeta struct {
	RoundStart int
	PlayStart  int
	End        int
}

type RoundHistory struct {
	Player string
	Kills  int
	Events []string
}

type Score struct {
	Player string
	Kills  int
	Deaths int
}
