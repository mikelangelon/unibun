package game

type GameState int

const (
	StateMenu GameState = iota
	StatePlaying
	StateRandom
	StateExiting
)
