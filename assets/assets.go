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

	//go:embed pidgeon.png
	Pidgeon []byte

	//go:embed mouse.png
	Mouse []byte

	//go:embed snake.png
	Snake []byte

	//go:embed duck.png
	Duck []byte

	//go:embed cat.png
	Cat []byte

	//go:embed tile2.png
	FloorTile []byte

	//go:embed tile2b.png
	FloorTileB []byte

	//go:embed tile2c.png
	FloorTileC []byte

	//go:embed menu2.png
	MenuBackground []byte

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
