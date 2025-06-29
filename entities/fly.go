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
	fly := &Fly{
		Enemy: NewEnemy(startX, startY, []*ebiten.Image{
			common.GetImage(assets.Fly),
			common.GetImage(assets.Fly2),
			common.GetImage(assets.Fly3),
		}),
		targetX: -1,
		targetY: -1,
	}
	fly.animationMode = AnimateContinuously
	return fly
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

	oldX, oldY := fe.gridX, fe.gridY

	newX, newY := oldX, oldY
	// Try to reduce distance to the target
	if targetX != oldX {
		nextX := oldX + int(math.Copysign(1, float64(targetX-oldX)))
		if level.IsWalkable(nextX, oldY) {
			newX = nextX
		}
	} else if targetY != oldY {
		nextY := oldY + int(math.Copysign(1, float64(targetY-oldY)))
		if level.IsWalkable(oldX, nextY) {
			newY = nextY
		}
	}

	return updatePosDirection(fe, oldX, oldY, newX, newY)
}
