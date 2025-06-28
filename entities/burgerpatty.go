package entities

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mikelangelon/unibun/assets"
	"github.com/mikelangelon/unibun/common"
	"github.com/mikelangelon/unibun/config"
)

type BurgerPatty struct {
	GridX, GridY               int
	Image                      *ebiten.Image
	initialGridX, initialGridY int
	pulseOffset                float64
}

func NewBurgerPatty(startX, startY int) BurgerPatty {
	return BurgerPatty{
		GridX:        startX,
		GridY:        startY,
		Image:        common.GetImage(assets.BurgerPatty),
		initialGridX: startX,
		initialGridY: startY,
	}
}

func (bp *BurgerPatty) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	scale := 1.0 + 0.1*math.Sin(bp.pulseOffset/20.0)

	w, h := bp.Image.Bounds().Dx(), bp.Image.Bounds().Dy()
	centerX, centerY := float64(w)/2, float64(h)/2

	op.GeoM.Translate(-centerX, -centerY)
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(centerX, centerY)

	op.GeoM.Translate(float64(bp.GridX*config.TileSize), float64(bp.GridY*config.TileSize))

	screen.DrawImage(bp.Image, op)
	bp.pulseOffset++
}

func (bp *BurgerPatty) Reset() {
	bp.GridX = bp.initialGridX
	bp.GridY = bp.initialGridY
}
