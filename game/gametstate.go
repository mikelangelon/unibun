package game

type GameState int

const (
	StateMenu GameState = iota
	StatePlaying
	StateEndless
	StateExiting
	StatePaused
	StateIntro
)
