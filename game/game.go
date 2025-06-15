package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mikelangelon/unibun/config"
	"github.com/mikelangelon/unibun/entities"
	"github.com/mikelangelon/unibun/level"
)

type Game struct {
	levels  []*level.Level
	players []*entities.Player
}

func NewGame() *Game {
	return &Game{
		levels: []*level.Level{level.NewLevel0()},
		players: []*entities.Player{
			entities.NewPlayer(1, 1, config.TopBun),
			entities.NewPlayer(2, 2, config.BottomBun),
			entities.NewPlayer(3, 3, config.BurguerPatty),
			entities.NewPlayer(4, 4, config.Cheese),
		},
	}
}

func (g *Game) Update() error {
	g.players[0].Update(g.levels[0])
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	g.levels[0].Draw(screen)
	for _, player := range g.players {
		if player != nil {
			player.Draw(screen)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return config.WindowWidth, config.WindowHeight
}
