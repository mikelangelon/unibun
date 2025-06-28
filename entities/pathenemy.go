package entities

import (
	"github.com/mikelangelon/unibun/common"
	"image"
	"log/slog"

	"github.com/mikelangelon/unibun/assets"
)

type PathEnemy struct {
	*Enemy
	path             []image.Point // The sequence of grid coordinates to follow
	currentPathIndex int

	// Store initial state for reset TODO find better way
	initialPathIndex int
	initialDirection int
	initialAsset     []byte
}

func NewPathEnemy(startX, startY int, path []image.Point) *PathEnemy {
	if len(path) == 0 {
		slog.Warn("PathEnemy created with an empty path. It will not move.", "startX", startX, "startY", startY)
	}
	initialX, initialY := startX, startY
	if len(path) > 0 {
		initialX = path[0].X
		initialY = path[0].Y
	}
	return &PathEnemy{
		Enemy:            NewEnemy(initialX, initialY, common.GetImage(assets.Mouse)),
		path:             path,
		currentPathIndex: 0,
		initialPathIndex: 0,
	}
}

// Reset puts the enemy back to work at the beginning of its path
func (e *PathEnemy) Reset() {
	e.Enemy.Reset()
	e.currentPathIndex = e.initialPathIndex
}

func (e *PathEnemy) Update(_ Level) bool {
	oldX := e.gridX
	newX, newY := e.followPath()
	return updatePosDirection(e.Enemy, oldX, newX, newY)
}

func (e *PathEnemy) followPath() (int, int) {
	potentialNextPathIndex := e.currentPathIndex + e.facingDirection

	if potentialNextPathIndex >= len(e.path) {
		// Reached the end, reverse direction and move to the second to last point
		e.currentPathIndex = len(e.path) - 2
	} else if potentialNextPathIndex < 0 {
		// If it was at path[0], it should move to path[1].
		e.currentPathIndex = 1
	} else {
		// Continue in the current direction
		e.currentPathIndex = potentialNextPathIndex
	}
	targetPoint := e.path[e.currentPathIndex]
	return targetPoint.X, targetPoint.Y

}
