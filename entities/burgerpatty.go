package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mikelangelon/unibun/assets"
	"github.com/mikelangelon/unibun/common"
	"github.com/mikelangelon/unibun/config"
)

type BurgerPatty struct {
	GridX, GridY int
	Image        *ebiten.Image
}

func NewBurgerPatty(startX, startY int) BurgerPatty {
	return BurgerPatty{
		GridX: startX,
		GridY: startY,
		Image: common.GetImage(assets.BurgerPatty),
	}
}

func (bp *BurgerPatty) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(bp.GridX*config.TileSize), float64(bp.GridY*config.TileSize))
	screen.DrawImage(bp.Image, op)
}
