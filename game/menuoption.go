package game

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/mikelangelon/unibun/config"
)

type MenuOption struct {
	Text   string
	Rect   image.Rectangle
	Action func(*Game)
}

func (g *Game) initMenu() {
	centerX := config.WindowWidth / 2
	totalMenuHeight := config.MenuOptionHeight*3 + config.MenuOptionSpacing*2
	startY := (config.WindowHeight-totalMenuHeight)/2 + 100

	g.menuOptions = []MenuOption{
		{
			Text: "Play",
			Action: func(game *Game) {
				game.currentGameState = StateLevelSelect
			},
		},
		{
			Text: "Endless",
			Action: func(game *Game) {
				game.startEndlessGame()
			},
		},
		{
			Text: "How to Play",
			Action: func(game *Game) {
				game.currentGameState = StateTutorial
			},
		},
	}
	for i := range g.menuOptions {
		option := &g.menuOptions[i]
		option.Rect = image.Rect(
			centerX-config.MenuOptionWidth/2,
			startY+i*(config.MenuOptionHeight+config.MenuOptionSpacing),
			centerX+config.MenuOptionWidth/2,
			startY+i*(config.MenuOptionHeight+config.MenuOptionSpacing)+config.MenuOptionHeight,
		)
	}

	g.selectedMenuOption = 0
}

func (g *Game) updateMenu() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		g.selectedMenuOption--
		if g.selectedMenuOption < 0 {
			g.selectedMenuOption = len(g.menuOptions) - 1
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		g.selectedMenuOption++
		if g.selectedMenuOption >= len(g.menuOptions) {
			g.selectedMenuOption = 0
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.menuOptions[g.selectedMenuOption].Action(g)
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		for i, option := range g.menuOptions {
			if image.Pt(x, y).In(option.Rect) {
				g.selectedMenuOption = i
				option.Action(g)
				break
			}
		}
	}
	return nil
}

func (g *Game) drawMenu(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.6, 0.6)
	screen.DrawImage(g.menuBackground, op)

	for i, option := range g.menuOptions {
		btnColor := color.RGBA{0x40, 0x40, 0x40, 0xFF}
		if i == g.selectedMenuOption {
			btnColor = color.RGBA{0x80, 0x80, 0x80, 0xFF}
		}
		ebitenutil.DrawRect(screen, float64(option.Rect.Min.X), float64(option.Rect.Min.Y), float64(option.Rect.Dx()), float64(option.Rect.Dy()), btnColor)

		charWidth := 6
		charHeight := 16
		textX := option.Rect.Min.X + (option.Rect.Dx()-len(option.Text)*charWidth)/2
		textY := option.Rect.Min.Y + (option.Rect.Dy()-charHeight)/2
		ebitenutil.DebugPrintAt(screen, option.Text, textX, textY)
	}
}
