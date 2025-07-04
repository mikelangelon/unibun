package entities

import (
	"image"
	"log/slog"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mikelangelon/unibun/assets"
	"github.com/mikelangelon/unibun/common"
)

type Mouse struct {
	*Enemy
	path             []image.Point // The sequence of grid coordinates to follow
	currentPathIndex int

	// Store initial state for reset TODO find better way
	initialPathIndex int
	initialDirection int
	pathDirection    int
	initialAsset     []byte
}

func NewMouse(startX, startY int, path []image.Point) *Mouse {
	if len(path) == 0 {
		slog.Warn("Mouse created with an empty path. It will not move.", "startX", startX, "startY", startY)
	}
	initialX, initialY := startX, startY
	if len(path) > 0 {
		initialX = path[0].X
		initialY = path[0].Y
	}
	return &Mouse{
		Enemy: NewEnemy(initialX, initialY, []*ebiten.Image{
			common.GetImage(assets.Mouse),
			common.GetImage(assets.Mouse2),
			common.GetImage(assets.Mouse3),
		}),
		path:             path,
		currentPathIndex: 0,
		pathDirection:    1,
		initialPathIndex: 0,
		initialDirection: 1,
	}
}

// Reset puts the enemy back to work at the beginning of its path
func (e *Mouse) Reset() {
	e.Enemy.Reset()
	e.currentPathIndex = e.initialPathIndex
	e.pathDirection = e.initialDirection
}

func (e *Mouse) Update(_ Level) bool {
	if len(e.path) <= 1 {
		return true
	}
	oldX, oldY := e.gridX, e.gridY
	newX, newY := e.followPath()
	return updatePosDirection(e.Enemy, oldX, oldY, newX, newY)
}

func (e *Mouse) followPath() (int, int) {
	potentialNextPathIndex := e.currentPathIndex + e.pathDirection

	if potentialNextPathIndex >= len(e.path) {
		// Reached the end, reverse direction and move to the second to last point
		e.pathDirection = -1
		e.currentPathIndex = len(e.path) - 2
	} else if potentialNextPathIndex < 0 {
		// If it was at path[0], it should move to path[1].
		e.pathDirection = 1
		e.currentPathIndex = 1
	} else {
		// Continue in the current direction
		e.currentPathIndex = potentialNextPathIndex
	}
	targetPoint := e.path[e.currentPathIndex]
	return targetPoint.X, targetPoint.Y

}
