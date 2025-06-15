package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mikelangelon/unibun/config"
)

type Game struct {
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black) // You can choose any color for the padding.
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// The logical screen size is now the full window size.
	return config.WindowWidth, config.WindowHeight
}
