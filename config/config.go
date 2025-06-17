package config

const (
	TileSize = 32

	// Padding constants
	PaddingTop    = 32
	PaddingRight  = 32
	PaddingBottom = 96
	PaddingLeft   = 32

	screenWidth  = 800
	screenHeight = 480

	WindowWidth  = screenWidth + PaddingLeft + PaddingRight
	WindowHeight = screenHeight + PaddingTop + PaddingBottom

	// Other game stuff
	EnemyTurnDelayDuration = 12 // 200ms delay at 60FPS (200 / (1000/60))

)
