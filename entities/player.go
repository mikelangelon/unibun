package entities

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image"
	"log/slog"

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
	targetX, targetY := p.GridX, p.GridY
	playerAttemptedMove := false

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
		targetX--
		playerAttemptedMove = true
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
		targetX++
		playerAttemptedMove = true
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		targetY--
		playerAttemptedMove = true
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		targetY++
		playerAttemptedMove = true
	}
	if !playerAttemptedMove {
		return false
	}
	if level.IsWalkable(targetX, targetY) {
		p.GridX = targetX
		p.GridY = targetY
	}
	return true
}
