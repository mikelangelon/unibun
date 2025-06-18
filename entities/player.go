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

	dashMove *dashMove
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
	if p.dashMove != nil {
		// If currently dashing, process the next step of the dash.
		return p.processDashStep()
	}
	dx, dy := 0, 0
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
		return false // Nothing to update
	}

	isShiftPressed := ebiten.IsKeyPressed(ebiten.KeyShiftLeft)
	if isShiftPressed {
		// Attempt to start a dash.
		if p.startDash(level, dx, dy) {
			// Dash successfully initiated. The turn is not over yet;
			return false
		}
		return false
	} else {
		return p.performSingleMove(level, dx, dy)
	}
}

// move is a single step movement
func (p *Player) performSingleMove(level Level, dx, dy int) bool {
	targetX := p.GridX + dx
	targetY := p.GridY + dy

	if !level.IsWalkable(targetX, targetY) {
		return false
	}
	p.GridX = targetX
	p.GridY = targetY
	return true
}
