package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mikelangelon/unibun/config"
)

// TODO Rename this monstruosity name
type Enemier interface {
	Draw(screen *ebiten.Image)
	Update(level Level) bool
	Collision(player *Player) bool
	Image() *ebiten.Image
	Reset()
	Position() (int, int)
}

type Enemy struct {
	gridX, gridY               int
	initialGridX, initialGridY int
	image                      *ebiten.Image
	facingDirection            int
}

func NewEnemy(startX, startY int, img *ebiten.Image) *Enemy {
	return &Enemy{
		gridX:           startX,
		gridY:           startY,
		initialGridX:    startX,
		initialGridY:    startY,
		image:           img,
		facingDirection: 1,
	}
}
func (e *Enemy) Collision(player *Player) bool {
	return e.gridX == player.GridX && e.gridY == player.GridY
}

func (e *Enemy) Update(level Level) bool {
	oldX := e.gridX
	newX, newY := nextRandomMove(level, e)
	return updatePosDirection(e, oldX, newX, newY)
}

func (e *Enemy) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	if e.facingDirection == -1 {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(float64(config.TileSize), 0)
	}
	pixelX := float64(e.gridX * config.TileSize)
	pixelY := float64(e.gridY * config.TileSize)
	op.GeoM.Translate(pixelX, pixelY)
	screen.DrawImage(e.image, op)
}

func (e *Enemy) Image() *ebiten.Image {
	return e.image
}

func (e *Enemy) Reset() {
	e.gridX = e.initialGridX
	e.gridY = e.initialGridY
	e.facingDirection = 1
}

func (e *Enemy) Position() (int, int) {
	return e.gridX, e.gridY
}

func updatePosDirection(e *Enemy, oldX, newX, newY int) bool {
	if newX > oldX {
		e.facingDirection = 1
	} else if newX < oldX {
		e.facingDirection = -1
	}
	e.gridX, e.gridY = newX, newY
	return true
}
