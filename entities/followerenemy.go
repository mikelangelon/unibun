package entities

import (
	"math"

	"github.com/mikelangelon/unibun/common"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mikelangelon/unibun/assets"
	"github.com/mikelangelon/unibun/config"
)

type FollowerEnemy struct {
	gridX, gridY               int
	initialGridX, initialGridY int
	image                      *ebiten.Image
	targetPlayerType           config.PlayerType
	targetX, targetY           int
}

func NewFollowerEnemy(startX, startY int, targetType config.PlayerType) *FollowerEnemy {
	return &FollowerEnemy{
		gridX:            startX,
		gridY:            startY,
		initialGridX:     startX,
		initialGridY:     startY,
		image:            common.GetImage(assets.Pidgeon),
		targetPlayerType: targetType,
		targetX:          -1,
		targetY:          -1,
	}
}

func (fe *FollowerEnemy) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	pixelX := float64(fe.gridX * config.TileSize)
	pixelY := float64(fe.gridY * config.TileSize)
	op.GeoM.Translate(pixelX, pixelY)

	// Coloring based on target
	switch fe.targetPlayerType {
	case config.TopBun:
		op.ColorScale.Scale(1, 0.5, 0.5, 1)
	case config.BottomBun:
		op.ColorScale.Scale(0.5, 0.5, 1, 1)
	case config.Lettuce:
		op.ColorScale.Scale(0.5, 1, 0.5, 1)
	case config.Cheese:
		op.ColorScale.Scale(1, 1, 0.5, 1)
	}
	screen.DrawImage(fe.image, op)
}

func (fe *FollowerEnemy) SetTarget(x, y int) {
	fe.targetX = x
	fe.targetY = y
}

func (fe *FollowerEnemy) Update(level Level) bool {
	if fe.targetX == -1 || fe.targetY == -1 {
		// Move random if there is no target
		fe.gridX, fe.gridY = nextRandomMove(level, fe.gridX, fe.gridY)
		return true
	}

	currentX, currentY := fe.gridX, fe.gridY

	// Try to reduce distance to the target
	if fe.targetX != currentX {
		nextX := currentX + int(math.Copysign(1, float64(fe.targetX-currentX)))
		if level.IsWalkable(nextX, currentY) {
			fe.gridX = nextX
		}
	} else if fe.targetY != currentY {
		nextY := currentY + int(math.Copysign(1, float64(fe.targetY-currentY)))
		if level.IsWalkable(currentX, nextY) {
			fe.gridY = nextY
		}
	}

	return true
}

func (fe *FollowerEnemy) Collision(player *Player) bool {
	return fe.gridX == player.GridX && fe.gridY == player.GridY
}

func (fe *FollowerEnemy) Image() *ebiten.Image {
	return fe.image
}

func (fe *FollowerEnemy) Reset() {
	fe.gridX = fe.initialGridX
	fe.gridY = fe.initialGridY
	fe.targetX = -1
	fe.targetY = -1
}

func (fe *FollowerEnemy) GetTargetPlayerType() config.PlayerType {
	return fe.targetPlayerType
}
