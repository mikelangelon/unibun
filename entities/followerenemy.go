package entities

import (
	"github.com/mikelangelon/unibun/common"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mikelangelon/unibun/assets"
	"github.com/mikelangelon/unibun/config"
)

type FollowerEnemy struct {
	*Enemy
	targetPlayerType config.PlayerType
	targetX, targetY int
}

func NewFollowerEnemy(startX, startY int, targetType config.PlayerType) *FollowerEnemy {
	return &FollowerEnemy{
		Enemy:            NewEnemy(startX, startY, duckColorByTarget(targetType, assets.Duck)),
		targetPlayerType: targetType,
		targetX:          -1,
		targetY:          -1,
	}
}

func (fe *FollowerEnemy) SetTarget(x, y int) {
	fe.targetX = x
	fe.targetY = y
}

func (fe *FollowerEnemy) Update(level Level) bool {
	return followTarget(level, fe.Enemy, fe.targetX, fe.targetY)
}

func (fe *FollowerEnemy) Reset() {
	fe.Enemy.Reset()
	fe.targetX = -1
	fe.targetY = -1
}

func (fe *FollowerEnemy) GetTargetPlayerType() config.PlayerType {
	return fe.targetPlayerType
}

func duckColorByTarget(targetType config.PlayerType, b []byte) []*ebiten.Image {
	op := &ebiten.DrawImageOptions{}
	switch targetType {
	case config.TopBun:
		op.ColorScale.Scale(0.9, 0.9, 0.1, 1)
	case config.BottomBun:
		op.ColorScale.Scale(1, 0.6, 0.6, 1)
	}
	coloredImage := ebiten.NewImage(config.TileSize, config.TileSize)
	coloredImage.DrawImage(common.GetImage(b), op)
	return []*ebiten.Image{coloredImage}
}
