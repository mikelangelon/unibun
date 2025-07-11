package entities

import (
	"github.com/mikelangelon/unibun/common"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mikelangelon/unibun/assets"
	"github.com/mikelangelon/unibun/config"
)

type Duck struct {
	*Enemy
	targetPlayerType config.PlayerType
	targetX, targetY int
}

func NewDuck(startX, startY int, targetType config.PlayerType) *Duck {
	return &Duck{
		Enemy: NewEnemy(startX, startY,
			[]*ebiten.Image{
				duckColorByTarget(targetType, assets.Duck),
				duckColorByTarget(targetType, assets.Duck2),
				duckColorByTarget(targetType, assets.Duck3),
			},
		),
		targetPlayerType: targetType,
		targetX:          -1,
		targetY:          -1,
	}
}

func (fe *Duck) SetTarget(x, y int) {
	fe.targetX = x
	fe.targetY = y
}

func (fe *Duck) Update(level Level) bool {
	return followTarget(level, fe.Enemy, fe.targetX, fe.targetY)
}

func (fe *Duck) Reset() {
	fe.Enemy.Reset()
	fe.targetX = -1
	fe.targetY = -1
}

func (fe *Duck) GetTargetPlayerType() config.PlayerType {
	return fe.targetPlayerType
}

func duckColorByTarget(targetType config.PlayerType, b []byte) *ebiten.Image {
	op := &ebiten.DrawImageOptions{}
	switch targetType {
	case config.TopBun:
		op.ColorScale.Scale(0.9, 0.9, 0.1, 1)
	case config.BottomBun:
		op.ColorScale.Scale(0.8, 1.1, 1.2, 1)
	}
	coloredImage := ebiten.NewImage(config.TileSize, config.TileSize)
	op.GeoM.Scale(1, 0.9)
	coloredImage.DrawImage(common.GetImage(b), op)
	return coloredImage
}
