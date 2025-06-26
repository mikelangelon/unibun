package entities

import (
	"github.com/mikelangelon/unibun/common"
	"image"
	"log/slog"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mikelangelon/unibun/assets"
	"github.com/mikelangelon/unibun/config"
)

type PathEnemy struct {
	GridX, GridY     int
	image            *ebiten.Image
	path             []image.Point // The sequence of grid coordinates to follow
	currentPathIndex int
	direction        int // 1 for forward, -1 for backward

	// Store initial state for reset TODO find better way
	initialGridX, initialGridY int
	initialPathIndex           int
	initialDirection           int
	initialAsset               []byte
}

func NewPathEnemy(startX, startY int, path []image.Point, enemyAsset []byte) *PathEnemy {
	if len(path) == 0 {
		slog.Warn("PathEnemy created with an empty path. It will not move.", "startX", startX, "startY", startY)
	}
	initialX, initialY := startX, startY
	if len(path) > 0 {
		initialX = path[0].X
		initialY = path[0].Y
	}

	return &PathEnemy{
		GridX:            initialX,
		GridY:            initialY,
		image:            common.GetImage(assets.Pidgeon),
		path:             path,
		initialGridX:     initialX,
		initialGridY:     initialY,
		currentPathIndex: 0,
		direction:        1,
		initialPathIndex: 0,
		initialDirection: 1,
		initialAsset:     enemyAsset,
	}
}

func (pe *PathEnemy) Draw(screen *ebiten.Image) {
	if pe.image == nil {
		return
	}
	op := &ebiten.DrawImageOptions{}
	pixelX := float64(pe.GridX * config.TileSize)
	pixelY := float64(pe.GridY * config.TileSize)
	op.GeoM.Translate(pixelX, pixelY)
	screen.DrawImage(pe.image, op)
}

func (pe *PathEnemy) Update(_ Level) bool {
	if len(pe.path) <= 1 { // If no path or a single point path, don't move.
		return false
	}

	// Calculate the next potential index in the path
	potentialNextPathIndex := pe.currentPathIndex + pe.direction

	if potentialNextPathIndex >= len(pe.path) {
		// Reached the end, reverse direction and move to the second to last point
		pe.direction = -1
		// Ensure there's a point to move back to. If path has only 2 points, this makes it len-2 = 0.
		// If it was at path[len-1], it should move to path[len-2].
		pe.currentPathIndex = len(pe.path) - 2
	} else if potentialNextPathIndex < 0 {
		// Reached the beginning while going backward, reverse direction and move to the second point
		pe.direction = 1
		// If it was at path[0], it should move to path[1].
		pe.currentPathIndex = 1
	} else {
		// Continue in the current direction
		pe.currentPathIndex = potentialNextPathIndex
	}

	targetPoint := pe.path[pe.currentPathIndex]

	if pe.GridX == targetPoint.X && pe.GridY == targetPoint.Y {
		return false // Didn't move
	}

	pe.GridX = targetPoint.X
	pe.GridY = targetPoint.Y
	return true
}

func (pe *PathEnemy) Collision(player *Player) bool {
	return pe.GridX == player.GridX && pe.GridY == player.GridY
}

// Reset puts the enemy back to work at the beginning of its path
func (pe *PathEnemy) Reset() {
	if len(pe.path) > 0 {
		pe.GridX = pe.path[0].X
		pe.GridY = pe.path[0].Y
	}
	pe.currentPathIndex = pe.initialPathIndex
	pe.direction = pe.initialDirection
}

func (e *PathEnemy) Image() *ebiten.Image {
	return e.image
}
