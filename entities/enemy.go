package entities

import (
	"github.com/mikelangelon/unibun/common"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mikelangelon/unibun/assets"
	"github.com/mikelangelon/unibun/config"
)

// TODO Rename this monstruosity name
type Enemier interface {
	Draw(screen *ebiten.Image)
	Update(level Level) bool
	Collision(player *Player) bool
	Image() *ebiten.Image
	Reset()
}
type Enemy struct {
	gridX, gridY               int
	initialGridX, initialGridY int
	image                      *ebiten.Image
}

func NewEnemy(startX, startY int) *Enemy {
	return &Enemy{
		gridX:        startX,
		gridY:        startY,
		initialGridX: startX,
		initialGridY: startY,
		image:        common.GetImage(assets.Pidgeon),
	}
}

func (e *Enemy) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	pixelX := float64(e.gridX * config.TileSize)
	pixelY := float64(e.gridY * config.TileSize)
	op.GeoM.Translate(pixelX, pixelY)
	screen.DrawImage(e.image, op)
}

func (e *Enemy) Update(level Level) bool {
	e.gridX, e.gridY = nextRandomMove(level, e.gridX, e.gridY)
	return true
}

func (e *Enemy) Collision(player *Player) bool {
	return e.gridX == player.GridX && e.gridY == player.GridY
}

func (e *Enemy) Image() *ebiten.Image {
	return e.image
}

func (e *Enemy) Reset() {
	e.gridX = e.initialGridX
	e.gridY = e.initialGridY
}

// Common functions
func nextRandomMove(level Level, x, y int) (int, int) {
	possibleMoves := []struct{ dx, dy int }{
		{0, -1}, {0, 1}, {-1, 0}, {1, 0},
	}
	rand.Shuffle(len(possibleMoves), func(i, j int) {
		possibleMoves[i], possibleMoves[j] = possibleMoves[j], possibleMoves[i]
	})
	for _, move := range possibleMoves {
		targetX, targetY := x+move.dx, y+move.dy
		if !level.IsWalkable(targetX, targetY) {
			continue
		}
		return targetX, targetY
	}
	return x, y
}
