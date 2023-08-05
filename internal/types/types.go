package types

type Map struct {
	Teams   map[int]*Team
	Players map[int]string
	Winner  int
	Scores  map[int]int
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
