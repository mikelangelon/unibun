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
	EnemyTurnDelayDuration = 6  // 100ms delay at 60FPS (100 / (1000/60))
	NextLevelDelayDuration = 30 // 500ms delay at 60FPS (500 / (1000/60))

	DashStepDelay = 1 // Increase for slower speed

	MenuOptionWidth   = 120
	MenuOptionHeight  = 30
	MenuOptionSpacing = 10
)
