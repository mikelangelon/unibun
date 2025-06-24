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

	//go:embed pidgeon.png
	Pidgeon []byte
)
