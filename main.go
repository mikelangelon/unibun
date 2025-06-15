package main

import (
	"log/slog"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mikelangelon/unibun/config"
	"github.com/mikelangelon/unibun/game"
)

func main() {
	ebiten.SetWindowSize(config.WindowWidth, config.WindowHeight)
	ebiten.SetWindowTitle("UniBun")
	game := game.NewGame()
	if err := ebiten.RunGame(game); err != nil {
		slog.Error("error running game", "error", err)
	}
}
