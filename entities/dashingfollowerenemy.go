package entities

import (
	"fmt"
	"github.com/mikelangelon/unibun/common"
	"log/slog"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/mikelangelon/unibun/assets"
	"github.com/mikelangelon/unibun/config"
)

type DashingFollowerEnemy struct {
	FollowerEnemy
	dashState *DashState

	// turnsToDash defines how many turns before the enemy attempts to dash
	turnsToDash int
	// dashCounter counter for turns until next dash
	dashCounter int
}

func NewDashingFollowerEnemy(startX, startY int, targetType config.PlayerType, turnsToDash int) *DashingFollowerEnemy {
	fe := FollowerEnemy{
		Enemy:            NewEnemy(startX, startY, imageByTarget(targetType, assets.Snake)),
		targetPlayerType: targetType,
		targetX:          -1,
		targetY:          -1,
	}

	return &DashingFollowerEnemy{
		FollowerEnemy: fe,
		dashState:     NewDashState(),
		turnsToDash:   turnsToDash,
		dashCounter:   turnsToDash, // Start with full counter
	}
}

func (dfe *DashingFollowerEnemy) Draw(screen *ebiten.Image) {
	dfe.Enemy.Draw(screen)

	if dfe.dashCounter > 0 {
		s := fmt.Sprintf("%d", dfe.dashCounter)
		x := dfe.gridX*config.TileSize + config.TileSize - 10
		y := dfe.gridY*config.TileSize + 2
		ebitenutil.DebugPrintAt(screen, s, x, y-20)
	}
}

func (dfe *DashingFollowerEnemy) Update(level Level) bool {
	if dfe.dashState.IsActive() {
		newX, newY, _, finished := dfe.dashState.Update(dfe.gridX, dfe.gridY, level, false)
		dfe.gridX = newX
		dfe.gridY = newY
		return finished
	}

	dfe.dashCounter--
	if dfe.dashCounter <= 0 {
		if dfe.targetX != -1 && dfe.targetY != -1 {
			dx, dy := 0, 0
			diffX := dfe.targetX - dfe.gridX
			diffY := dfe.targetY - dfe.gridY

			// Avoid diagonal option
			if math.Abs(float64(diffX)) >= math.Abs(float64(diffY)) {
				if diffX != 0 {
					dx = int(math.Copysign(1, float64(diffX)))
				}
			} else {
				if diffY != 0 {
					dy = int(math.Copysign(1, float64(diffY)))
				}
			}

			// Going to dash
			if dx != 0 || dy != 0 {
				slog.Info("DashingFollowerEnemy trying to dash", "from_x", dfe.gridX, "from_y", dfe.gridY, "dx", dx, "dy", dy)
				if dfe.dashState.Start(dfe.gridX, dfe.gridY, dx, dy, 5, level, false) {
					dfe.dashCounter = dfe.turnsToDash

					newX, newY, _, finished := dfe.dashState.Update(dfe.gridX, dfe.gridY, level, false)
					dfe.gridX = newX
					dfe.gridY = newY
					return finished
				}
			}
		}

		// If code reaches here --> it means the dash couldn't even start due to obstacle)
		// Reset count and continue turn TODO: Try other possibilities
		dfe.dashCounter = dfe.turnsToDash
		return true
	}

	return dfe.FollowerEnemy.Update(level)
}

func (dfe *DashingFollowerEnemy) Reset() {
	dfe.FollowerEnemy.Reset()
	dfe.dashState.Reset()
	dfe.dashCounter = dfe.turnsToDash
}

func imageByTarget(targetType config.PlayerType, b []byte) *ebiten.Image {
	op := &ebiten.DrawImageOptions{}
	// Coloring based on target
	switch targetType {
	case config.TopBun:
		op.ColorScale.Scale(1, 0.5, 0.5, 1)
	case config.BottomBun:
		op.ColorScale.Scale(0.5, 0.5, 1, 1)
	case config.Lettuce:
		op.ColorScale.Scale(0.1, 1, 0.1, 1)
	case config.Cheese:
		op.ColorScale.Scale(1, 1, 0.1, 1)
	}
	coloredImage := ebiten.NewImage(config.TileSize, config.TileSize)
	coloredImage.DrawImage(common.GetImage(b), op)
	return coloredImage
}
