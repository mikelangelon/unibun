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
	BurguerPatty []byte

	//go:embed cheese.png
	Cheese []byte
)
