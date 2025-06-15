package game

import (
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mikelangelon/unibun/config"
	"github.com/mikelangelon/unibun/entities"
	"github.com/mikelangelon/unibun/level"
)

type Game struct {
	levels      []*level.Level
	players     []*entities.Player
	turnManager turnManager

	gameScreen *ebiten.Image
}

type turnManager struct {
	currentTurn      int
	turnOrderDisplay []interface{}
}

func NewGame() *Game {
	g := Game{
		levels: []*level.Level{level.NewLevel0()},
		players: []*entities.Player{
			entities.NewPlayer(1, 1, config.TopBun),
			entities.NewPlayer(2, 2, config.BottomBun),
			entities.NewPlayer(3, 3, config.BurguerPatty),
			entities.NewPlayer(4, 4, config.Cheese),
		},
		turnManager: turnManager{
			currentTurn: 0,
		},
	}
	g.buildTurnOrderDisplay()
	return &g
}

func (g *Game) Update() error {
	g.players[0].Update(g.levels[0])
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	// Ensure offscreen buffer is initialized
	// g.offscreen will hold the actual game content (screenWidth x screenHeight)
	if g.gameScreen == nil {
		g.gameScreen = ebiten.NewImage(g.levels[0].ScreenWidth(), g.levels[0].ScreenHeight())
	}
	g.gameScreen.Clear() // Clear the offscreen buffer

	g.gameScreen.Fill(color.RGBA{R: 0x10, G: 0x10, B: 0x10, A: 0xff})
	g.levels[0].Draw(g.gameScreen)
	for _, player := range g.players {
		if player != nil {
			player.Draw(g.gameScreen)
		}
	}
	// Finally draw screen
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(config.PaddingLeft), float64(config.PaddingTop))
	screen.DrawImage(g.gameScreen, op)

	g.drawTurnOrder(screen)
}

// Constants for Turn Order UI
const (
	turnOrderIconSize      = 24
	turnOrderIconSpacing   = 6
	turnOrderTextMarginX   = 10
	turnOrderTextOffsetY   = 4
	turnOrderIconTopMargin = 18
)

func (g *Game) drawTurnOrder(screen *ebiten.Image) {
	uiAreaStartY := float64(g.currentLevel().ScreenHeight() + config.PaddingTop)

	orderText := "Order:"
	textRenderX := float64(config.PaddingLeft + turnOrderTextMarginX)

	// Align text top with icon top + small offset for visual balance
	textRenderY := uiAreaStartY + float64(turnOrderIconTopMargin+turnOrderTextOffsetY)
	ebitenutil.DebugPrintAt(screen, orderText, int(textRenderX), int(textRenderY))
	// Estimate width of "Order: " text to position icons.
	// TODO: check other ways? Maybe use ebiten/text and text.BoundString
	iconStartX := textRenderX + float64(len(orderText)*7) + float64(turnOrderIconSpacing) // Approx 7px per char for default font
	for i, entity := range g.turnManager.turnOrderDisplay {
		iconX := iconStartX + float64(i*(turnOrderIconSize+turnOrderIconSpacing))
		iconY := uiAreaStartY + float64(turnOrderIconTopMargin)

		switch item := entity.(type) {
		case *entities.Player:
			op := &ebiten.DrawImageOptions{}
			pixelX := float64(iconX)
			pixelY := float64(iconY)
			op.GeoM.Translate(pixelX, pixelY)
			screen.DrawImage(item.Image, op)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return config.WindowWidth, config.WindowHeight
}

func (g *Game) currentLevel() *level.Level {
	return g.levels[0]
}

func (g *Game) buildTurnOrderDisplay() {
	g.turnManager.turnOrderDisplay = []interface{}{}
	pattern := g.currentLevel().TurnOrderPattern
	if len(pattern) == 0 {
		return
	}

	for i := 0; i < len(pattern); i++ {
		idx := (g.turnManager.currentTurn + i) % len(pattern)
		actor := pattern[idx]

		switch actualActor := actor.(type) {
		case config.PlayerType:
			switch actualActor {
			case config.TopBun:
				g.turnManager.turnOrderDisplay = append(g.turnManager.turnOrderDisplay, g.players[0])
			case config.BottomBun:
				g.turnManager.turnOrderDisplay = append(g.turnManager.turnOrderDisplay, g.players[1])
			case config.Cheese:
				g.turnManager.turnOrderDisplay = append(g.turnManager.turnOrderDisplay, g.players[2])
			}
		}
	}
}
