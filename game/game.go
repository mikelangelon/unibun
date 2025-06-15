package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mikelangelon/unibun/config"
	"github.com/mikelangelon/unibun/level"
)

type Game struct {
	levels []*level.Level
}

func NewGame() *Game {
	return &Game{
		levels: []*level.Level{level.NewLevel0()},
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	g.levels[0].Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return config.WindowWidth, config.WindowHeight
}
