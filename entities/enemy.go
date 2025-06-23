package entities

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mikelangelon/unibun/assets"
	"github.com/mikelangelon/unibun/config"
	"image"
	"log/slog"
	"math/rand"
)

// TODO Rename this monstruosity name
type Enemier interface {
	Draw(screen *ebiten.Image)
	Update(level Level) bool
	Collision(player *Player) bool
	Image() *ebiten.Image
}
type Enemy struct {
	gridX, gridY int
	image        *ebiten.Image
}

func NewEnemy(startX, startY int) *Enemy {
	playerDecoded, _, err := image.Decode(bytes.NewReader(assets.Pidgeon))
	if err != nil {
		slog.Error("unexpected error decoding enemy image", "error", err)
		return nil
	}
	img := ebiten.NewImageFromImage(playerDecoded)
	return &Enemy{
		gridX: startX,
		gridY: startY,
		image: img,
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
	moved := false

	possibleMoves := []struct{ dx, dy int }{
		{0, -1}, {0, 1}, {-1, 0}, {1, 0},
	}
	rand.Shuffle(len(possibleMoves), func(i, j int) {
		possibleMoves[i], possibleMoves[j] = possibleMoves[j], possibleMoves[i]
	})

	for _, move := range possibleMoves {
		targetX, targetY := e.gridX+move.dx, e.gridY+move.dy
		if !level.IsWalkable(targetX, targetY) {
			continue
		}
		e.gridX = targetX
		e.gridY = targetY
		moved = true
		break
	}
	return moved
}

func (e *Enemy) Collision(player *Player) bool {
	return e.gridX == player.GridX && e.gridY == player.GridY
}

func (e *Enemy) Image() *ebiten.Image {
	return e.image
}
