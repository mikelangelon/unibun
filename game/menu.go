package game

import (
	"bytes"
	"github.com/mikelangelon/unibun/assets"
	"github.com/mikelangelon/unibun/common"
	"image"
	"image/color"
	"log/slog"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/mikelangelon/unibun/config"
)

type menuOption struct {
	text   string
	rect   image.Rectangle
	action func(*Game)
}

type menu struct {
	menuBackground     *ebiten.Image
	menuOptions        []*menuOption
	selectedMenuOption int
	fallingIngredients []*fallingIngredient
	spawnTimer         int
	ingredientImages   []*ebiten.Image
}

func (m *menu) spawnFallingIngredient() {
	if len(m.fallingIngredients) > 15 {
		return
	}
	img := m.ingredientImages[rand.IntN(len(m.ingredientImages))]
	item := &fallingIngredient{
		x:     rand.Float64() * float64(config.WindowWidth),
		y:     -float64(config.TileSize),
		speed: 2 + rand.Float64()*4,
		image: img,
	}
	m.fallingIngredients = append(m.fallingIngredients, item)
}

type fallingIngredient struct {
	x, y  float64
	speed float64
	image *ebiten.Image
}

func newMenu() *menu {
	centerX := 467
	startY := 400
	img, _, err := image.Decode(bytes.NewReader(assets.MenuBackground))
	if err != nil {
		slog.Error("failed to decode menu background", "error", err)
	}
	m := &menu{
		menuBackground: ebiten.NewImageFromImage(img),
		ingredientImages: []*ebiten.Image{
			common.GetImage(assets.TopBun),
			common.GetImage(assets.BottomBun),
			common.GetImage(assets.BurgerPatty),
			common.GetImage(assets.Lettuce),
			common.GetImage(assets.Cheese),
		},
		fallingIngredients: []*fallingIngredient{},
		menuOptions: []*menuOption{
			{
				text: "Play",
				action: func(game *Game) {
					game.currentGameState = StateLevelSelect
				},
			},
			{
				text: "Endless",
				action: func(game *Game) {
					game.startEndlessGame()
				},
			},
			{
				text: "How to Play",
				action: func(game *Game) {
					game.currentGameState = StateTutorial
				},
			},
		},
	}
	for i, op := range m.menuOptions {
		op.rect = image.Rect(
			centerX-config.MenuOptionWidth/2,
			startY+i*(config.MenuOptionHeight+config.MenuOptionSpacing),
			centerX+config.MenuOptionWidth/2,
			startY+i*(config.MenuOptionHeight+config.MenuOptionSpacing)+config.MenuOptionHeight,
		)
	}
	return m
}

func (m *menu) updateMenu(g *Game) error {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		m.selectedMenuOption--
		if m.selectedMenuOption < 0 {
			m.selectedMenuOption = len(m.menuOptions) - 1
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		m.selectedMenuOption++
		if m.selectedMenuOption >= len(m.menuOptions) {
			m.selectedMenuOption = 0
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		m.menuOptions[m.selectedMenuOption].action(g)
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		for i, option := range m.menuOptions {
			if image.Pt(x, y).In(option.rect) {
				m.selectedMenuOption = i
				option.action(g)
				break
			}
		}
	}
	m.updateFallingIngredients()
	return nil
}

func (m *menu) updateFallingIngredients() {
	m.spawnTimer--
	if m.spawnTimer <= 0 {
		m.spawnTimer = 5
		m.spawnFallingIngredient()
	}

	var updated []*fallingIngredient
	for _, item := range m.fallingIngredients {
		item.y += item.speed
		if item.y < float64(config.WindowHeight) {
			updated = append(updated, item)
		}
	}
	m.fallingIngredients = updated
}

// FIXME: find replacement for deprecated ebitenutil.DrawRect
func (m *menu) drawMenu(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.6, 0.6)
	screen.DrawImage(m.menuBackground, op)

	for i, option := range m.menuOptions {
		btnColor := color.RGBA{0x40, 0x40, 0x40, 0xFF}
		if i == m.selectedMenuOption {
			btnColor = color.RGBA{0x80, 0x80, 0x80, 0xFF}
		}
		ebitenutil.DrawRect(screen, float64(option.rect.Min.X), float64(option.rect.Min.Y), float64(option.rect.Dx()), float64(option.rect.Dy()), btnColor)

		charWidth := 6
		charHeight := 16
		textX := option.rect.Min.X + (option.rect.Dx()-len(option.text)*charWidth)/2
		textY := option.rect.Min.Y + (option.rect.Dy()-charHeight)/2
		ebitenutil.DebugPrintAt(screen, option.text, textX, textY)
	}
	for _, item := range m.fallingIngredients {
		itemOp := &ebiten.DrawImageOptions{}
		itemOp.GeoM.Translate(item.x, item.y)
		screen.DrawImage(item.image, itemOp)
	}
}
