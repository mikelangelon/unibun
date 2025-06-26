package entities

import (
	"github.com/mikelangelon/unibun/common"
	"math"

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
	// TODO Use a different image
	fe := FollowerEnemy{
		gridX:            startX,
		gridY:            startY,
		initialGridX:     startX,
		initialGridY:     startY,
		image:            common.GetImage(assets.Pidgeon),
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
				if dfe.dashState.Start(dfe.gridX, dfe.gridY, dx, dy, 5, level, false) {
					dfe.dashCounter = dfe.turnsToDash

					newX, newY, _, finished := dfe.dashState.Update(dfe.gridX, dfe.gridY, level, false)
					dfe.gridX = newX
					dfe.gridY = newY
					return finished
				}
			}
		}

		dfe.dashCounter = dfe.turnsToDash
	}

	return dfe.FollowerEnemy.Update(level)
}

func (dfe *DashingFollowerEnemy) Reset() {
	dfe.FollowerEnemy.Reset()
	dfe.dashState.Reset()
	dfe.dashCounter = dfe.turnsToDash
}
