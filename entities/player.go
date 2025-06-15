package entities

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/mikelangelon/unibun/level"
	"image"
	"log/slog"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mikelangelon/unibun/assets"
	"github.com/mikelangelon/unibun/config"
)

type Player struct {
	gridX, gridY int
	PlayerType   config.PlayerType
	Image        *ebiten.Image
}

func NewPlayer(startX, startY int, playerType config.PlayerType) *Player {
	var b = []byte{}
	switch playerType {
	case config.TopBun:
		b = assets.TopBun
	case config.BottomBun:
		b = assets.BottomBun
	case config.Cheese:
		b = assets.Cheese
	case config.BurguerPatty:
		b = assets.BurguerPatty
	}
	playerDecoded, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		slog.Error("unexpected error decoding player image", "error", err)
		return nil
	}
	img := ebiten.NewImageFromImage(playerDecoded)
	return &Player{
		gridX:      startX,
		gridY:      startY,
		Image:      img,
		PlayerType: playerType,
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	pixelX := float64(p.gridX * config.TileSize)
	pixelY := float64(p.gridY * config.TileSize)
	op.GeoM.Translate(pixelX, pixelY)
	screen.DrawImage(p.Image, op)
}

func (p *Player) Update(level *level.Level) bool {
	targetX, targetY := p.gridX, p.gridY
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
		p.gridX = targetX
		p.gridY = targetY
	}
	return true
}
