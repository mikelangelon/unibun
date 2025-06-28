package entities

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mikelangelon/unibun/assets"
	"github.com/mikelangelon/unibun/common"
)

type Fly struct {
	*Enemy
	targetX, targetY int
}

func NewFly(startX, startY int) *Fly {
	return &Fly{
		Enemy: NewEnemy(startX, startY, []*ebiten.Image{
			common.GetImage(assets.Fly),
			common.GetImage(assets.Fly2),
			common.GetImage(assets.Fly3),
		}),
		targetX: -1,
		targetY: -1,
	}
}

func (fe *Fly) SetTarget(x, y int) {
	fe.targetX = x
	fe.targetY = y
}

func (fe *Fly) Update(level Level) bool {
	return followTarget(level, fe.Enemy, fe.targetX, fe.targetY)
}

func (fe *Fly) Reset() {
	fe.Enemy.Reset()
	fe.targetX = -1
	fe.targetY = -1
}

func followTarget(level Level, fe *Enemy, targetX, targetY int) bool {
	if targetX == -1 || targetY == -1 {
		// Move random if there is no target
		fe.gridX, fe.gridY = nextRandomMove(level, fe)
		return true
	}

	currentX, currentY := fe.gridX, fe.gridY

	newX, newY := currentX, currentY
	// Try to reduce distance to the target
	if targetX != currentX {
		nextX := currentX + int(math.Copysign(1, float64(targetX-currentX)))
		if level.IsWalkable(nextX, currentY) {
			newX = nextX
		}
	} else if targetY != currentY {
		nextY := currentY + int(math.Copysign(1, float64(targetY-currentY)))
		if level.IsWalkable(currentX, nextY) {
			newY = nextY
		}
	}

	return updatePosDirection(fe, currentX, newX, newY)
}
