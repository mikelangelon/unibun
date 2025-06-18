package entities

import (
	"bytes"
	"image"
	"log/slog"

	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mikelangelon/unibun/assets"
	"github.com/mikelangelon/unibun/config"
)

type Player struct {
	GridX, GridY int
	PlayerType   config.PlayerType
	Image        *ebiten.Image
}

type Level interface {
	IsWalkable(gridX, gridY int) bool
}

func NewPlayer(startX, startY int, playerType config.PlayerType) Player {
	var b = []byte{}
	switch playerType {
	case config.TopBun:
		b = assets.TopBun
	case config.BottomBun:
		b = assets.BottomBun
	case config.Cheese:
		b = assets.Cheese
	case config.BurguerPatty:
		b = assets.BurgerPatty
	}
	playerDecoded, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		slog.Error("unexpected error decoding player image", "error", err)
		return Player{}
	}
	img := ebiten.NewImageFromImage(playerDecoded)
	// calculate offsets for centering
	offsettedImg := ebiten.NewImage(32, 32)
	offsetY := float64(32-18) / 2.0
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, offsetY)
	offsettedImg.DrawImage(img, op)
	return Player{
		GridX:      startX,
		GridY:      startY,
		Image:      offsettedImg,
		PlayerType: playerType,
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	pixelX := float64(p.GridX * config.TileSize)
	pixelY := float64(p.GridY * config.TileSize)
	op.GeoM.Translate(pixelX, pixelY)
	screen.DrawImage(p.Image, op)
}

func (p *Player) Update(level Level) bool {
	dx, dy := 0, 0

	// Check for directional input
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
		dx = -1
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
		dx = 1
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		dy = -1
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		dy = 1
	}
	if dx == 0 && dy == 0 {
		return false
	}

	if ebiten.IsKeyPressed(ebiten.KeyShiftLeft) {
		return p.dash(level, dx, dy)
	}
	return p.move(level, dx, dy)
}

// move is a single step movement
func (p *Player) move(level Level, dx, dy int) bool {
	targetX := p.GridX + dx
	targetY := p.GridY + dy

	if !level.IsWalkable(targetX, targetY) {
		return false
	}
	p.GridX = targetX
	p.GridY = targetY
	return true
}

// dash is moving until next obstacle
func (p *Player) dash(level Level, dx, dy int) bool {
	currentX, currentY := p.GridX, p.GridY
	movedInDash := false

	for {
		nextX, nextY := currentX+dx, currentY+dy
		if level.IsWalkable(nextX, nextY) {
			currentX = nextX
			currentY = nextY
			movedInDash = true
		} else {
			// Hit an obstacle or wall
			break
		}
	}

	if movedInDash {
		p.GridX = currentX
		p.GridY = currentY
		return true
	}
	return false
}
