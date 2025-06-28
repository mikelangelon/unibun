package assets

import (
	_ "embed"
)

var (
	//go:embed topBun.png
	TopBun []byte

	//go:embed bottomBun.png
	BottomBun []byte

	//go:embed patty.png
	BurgerPatty []byte

	//go:embed cheese.png
	Cheese []byte

	//go:embed lettuce.png
	Lettuce []byte

	//go:embed client.png
	Client []byte

	// Animals
	//go:embed pidgeon.png
	Pidgeon []byte
	//go:embed mouse.png
	Mouse []byte
	//go:embed snake.png
	Snake []byte
	//go:embed duck.png
	Duck []byte
	//go:embed fly.png
	Fly []byte
	//go:embed fly2.png
	Fly2 []byte
	//go:embed fly3.png
	Fly3 []byte

	//go:embed tile1.png
	FloorTile []byte
	//go:embed tile2.png
	FloorTileB []byte
	//go:embed tile3.png
	FloorTileC []byte
	//go:embed tile4.png
	FloorTileD []byte
	//go:embed tile5.png
	FloorTileE []byte
	//go:embed tile6.png
	FloorTileF []byte
	//go:embed menu.png
	MenuBackground []byte

	//go:embed jamTheme.png
	JamTheme []byte
	//go:embed jam2025.png
	Jam2025 []byte

	/**
	Sounds
	*/
	//go:embed eat-323883.mp3
	EatingSound []byte
	//go:embed "Bonus - Character Select.mp3"
	MenuMusic []byte
	//go:embed "1- World Map.mp3"
	MainMusic []byte
)
