package entities

import (
	"math"

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
	if fe.targetX == -1 || fe.targetY == -1 {
		// Move random if there is no target
		fe.gridX, fe.gridY = nextRandomMove(level, fe.Enemy)
		return true
	}

	currentX, currentY := fe.gridX, fe.gridY

	newX, newY := currentX, currentY
	// Try to reduce distance to the target
	if fe.targetX != currentX {
		nextX := currentX + int(math.Copysign(1, float64(fe.targetX-currentX)))
		if level.IsWalkable(nextX, currentY) {
			newX = nextX
		}
	} else if fe.targetY != currentY {
		nextY := currentY + int(math.Copysign(1, float64(fe.targetY-currentY)))
		if level.IsWalkable(currentX, nextY) {
			newY = nextY
		}
	}

	return updatePosDirection(fe.Enemy, currentX, newX, newY)
}

func (fe *FollowerEnemy) Reset() {
	fe.Enemy.Reset()
	fe.targetX = -1
	fe.targetY = -1
}

func (fe *FollowerEnemy) GetTargetPlayerType() config.PlayerType {
	return fe.targetPlayerType
}

func duckColorByTarget(targetType config.PlayerType, b []byte) *ebiten.Image {
	op := &ebiten.DrawImageOptions{}
	switch targetType {
	case config.TopBun:
		op.ColorScale.Scale(0.9, 0.9, 0.1, 1)
	case config.BottomBun:
		op.ColorScale.Scale(1, 0.6, 0.6, 1)
	}
	coloredImage := ebiten.NewImage(config.TileSize, config.TileSize)
	coloredImage.DrawImage(common.GetImage(b), op)
	return coloredImage
}
