package entities

import (
	"bytes"
	"github.com/mikelangelon/unibun/assets"
	"github.com/mikelangelon/unibun/config"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type BurgerPatty struct {
	GridX, GridY int
	Image        *ebiten.Image
}

func NewBurgerPatty(startX, startY int) BurgerPatty {
	playerDecoded, _, err := image.Decode(bytes.NewReader(assets.BurgerPatty))
	if err != nil {
		return BurgerPatty{}
	}
	img := ebiten.NewImageFromImage(playerDecoded)
	return BurgerPatty{
		GridX: startX,
		GridY: startY,
		Image: img,
	}
}

func (bp *BurgerPatty) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(bp.GridX*config.TileSize), float64(bp.GridY*config.TileSize))
	screen.DrawImage(bp.Image, op)
}
