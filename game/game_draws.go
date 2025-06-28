package game

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/mikelangelon/unibun/config"
	"github.com/mikelangelon/unibun/entities"
)

// Constants for Turn Order UI
const (
	turnOrderIconSize      = 24
	turnOrderIconSpacing   = 6
	turnOrderTextMarginX   = 10
	turnOrderTextOffsetY   = 4
	turnOrderIconTopMargin = 18
)

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.currentGameState {
	case StateMenu:
		g.menu.drawMenu(screen)
	case StateLevelSelect:
		g.levelManager.draw(screen)
	case StateTutorial:
		drawTutorial(screen)
	case StateGameComplete:
		g.drawGameComplete(screen)
	case StateGameOver:
		g.drawGameOver(screen)
	case StatePlaying, StateEndless, StatePaused, StateIntro:
		g.drawPlaying(screen)
		if g.currentGameState == StateIntro {
			g.drawIntro(screen)
		}
		if g.currentGameState == StatePaused {
			g.drawPaused(screen)
		}
	}
}

func (g *Game) drawGameComplete(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 20, G: 20, B: 20, A: 255})

	const paddingLeft = 300
	title := "Congratulations!"
	titleX := paddingLeft
	ebitenutil.DebugPrintAt(screen, title, titleX, 150)

	msg1 := "You have completed all the levels!"
	msg1X := paddingLeft
	ebitenutil.DebugPrintAt(screen, msg1, msg1X, 180)

	msg2 := "Thanks for playing UniBun!"
	msg2X := paddingLeft
	ebitenutil.DebugPrintAt(screen, msg2, msg2X, 250)

	prompt := " --> Press Enter to return to the Main Menu"
	promptX := paddingLeft
	ebitenutil.DebugPrintAt(screen, prompt, promptX, 330)

	g.drawEffects(screen)
}

func (g *Game) drawGameOver(screen *ebiten.Image) {
	const paddingLeft = 300
	screen.Fill(color.RGBA{R: 20, G: 20, B: 20, A: 255})
	// TODO too similar as drawGameComplete
	title := fmt.Sprintf("You failed. You passed %d level(s).", g.currentEndlessLevel)
	titleX := paddingLeft
	ebitenutil.DebugPrintAt(screen, title, titleX, 150)

	prompt := " --> Press Enter to return to the Main Menu"
	promptX := paddingLeft
	ebitenutil.DebugPrintAt(screen, prompt, promptX, 230)

	prompt2 := " --> Press Space to retry"
	ebitenutil.DebugPrintAt(screen, prompt2, promptX, 280)
}

func (g *Game) drawPaused(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, 0, 0, float64(config.WindowWidth), float64(config.WindowHeight), color.RGBA{0, 0, 0, 128})

	for i, option := range g.pauseMenuOptions {
		btnColor := color.RGBA{0x40, 0x40, 0x40, 0xFF}
		if i == g.selectedPauseMenuOption {
			btnColor = color.RGBA{0x80, 0x80, 0x80, 0xFF}
		}
		ebitenutil.DrawRect(screen, float64(option.rect.Min.X), float64(option.rect.Min.Y), float64(option.rect.Dx()), float64(option.rect.Dy()), btnColor)

		charWidth := 6
		charHeight := 16
		textX := option.rect.Min.X + (option.rect.Dx()-len(option.text)*charWidth)/2
		textY := option.rect.Min.Y + (option.rect.Dy()-charHeight)/2
		ebitenutil.DebugPrintAt(screen, option.text, textX, textY)
	}
}

func (g *Game) drawPlaying(screen *ebiten.Image) {
	screen.Fill(color.Black)
	if g.gameScreen == nil {
		g.gameScreen = ebiten.NewImage(g.level.ScreenWidth(), g.level.ScreenHeight())
	}
	g.gameScreen.Clear()

	g.gameScreen.Fill(color.RGBA{R: 0x10, G: 0x10, B: 0x10, A: 0xff})
	g.currentLevel().Draw(g.gameScreen, g.patty == nil)
	for _, character := range g.turnManager.turnOrderDisplay {
		if character != nil {
			if g.animationManager.isMergeAnimationPlaying() {
				if p, ok := character.(*entities.Player); ok {
					if p.PlayerType == config.TopBun || p.PlayerType == config.BottomBun {
						continue
					}
				}
			}
			character.Draw(g.gameScreen)
		}
	}
	if g.patty != nil {
		if !g.animationManager.blockingAnimation() {
			g.patty.Draw(g.gameScreen)
		}
	}

	if g.animationManager.isMergeAnimationPlaying() {
		g.drawMergeAnimation(g.gameScreen)
	}
	if g.animationManager.isWinningPlaying() {
		g.drawWinAnimation(g.gameScreen)
	}

	g.drawEffects(g.gameScreen)

	// Draw winning message
	if g.status == Win {
		drawWinning(g.gameScreen)
	}
	// Finally draw screen
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(config.PaddingLeft), float64(config.PaddingTop))
	if g.shake != nil {
		g.shake.shake(op)
	}
	screen.DrawImage(g.gameScreen, op)

	g.drawTurnOrder(screen)
}

func (g *Game) drawMergeAnimation(screen *ebiten.Image) {
	topBun := g.turnManager.getPlayerType(config.TopBun)
	bottomBun := g.turnManager.getPlayerType(config.BottomBun)
	patty := g.patty
	g.animationManager.drawMergeAnimation(screen, patty, topBun, bottomBun)
}

func (g *Game) drawWinAnimation(screen *ebiten.Image) {
	g.animationManager.drawWinningAnimation(screen)
}

func (g *Game) drawIntro(screen *ebiten.Image) {
	if g.introDelayTimer > 0 {
		return
	}

	overlayColor := color.RGBA{0, 0, 0, 80}
	ebitenutil.DrawRect(screen, 0, 0, float64(config.WindowWidth), float64(config.WindowHeight), overlayColor)

	boxWidth := 400
	boxHeight := 150
	boxX := (config.WindowWidth - boxWidth) / 2
	boxY := (config.WindowHeight-boxHeight)/2 + 100
	boxColor := color.RGBA{40, 40, 40, 200}
	ebitenutil.DrawRect(screen, float64(boxX), float64(boxY), float64(boxWidth), float64(boxHeight), boxColor)

	introText := g.currentLevel().IntroText
	lines := strings.Split(introText, "\n")
	lineHeight := 16
	totalTextHeight := len(lines) * lineHeight
	startY := boxY + (boxHeight-totalTextHeight)/2

	for i, line := range lines {
		charWidth := 6
		textX := boxX + (boxWidth-len(line)*charWidth)/2
		textY := startY + i*lineHeight
		ebitenutil.DebugPrintAt(screen, line, textX, textY)
	}
}

func (g *Game) drawEffects(screen *ebiten.Image) {
	g.animationManager.drawEffects(screen)
}

func (g *Game) drawTurnOrder(screen *ebiten.Image) {
	uiAreaStartY := float64(g.currentLevel().ScreenHeight()) + config.PaddingTop

	orderText := "Order:"
	textRenderX := float64(config.PaddingLeft + turnOrderTextMarginX)

	textRenderY := uiAreaStartY + float64(turnOrderIconTopMargin+turnOrderTextOffsetY)
	ebitenutil.DebugPrintAt(screen, orderText, int(textRenderX), int(textRenderY))

	// TODO: check other ways? Maybe use ebiten/text and text.BoundString
	iconStartX := textRenderX + float64(len(orderText)*7) + float64(turnOrderIconSpacing) // Approx 7px per char for default font
	for i, entity := range g.turnManager.turnOrderDisplay {
		iconX := iconStartX + float64(i*(turnOrderIconSize+turnOrderIconSpacing))
		iconY := uiAreaStartY + float64(turnOrderIconTopMargin)

		switch item := entity.(type) {
		case *entities.Player:
			drawIcon(screen, item.Image, iconX, iconY)
		case entities.Enemier:
			drawIcon(screen, item.Icon(), iconX, iconY)
		}
	}
}

func drawIcon(screen *ebiten.Image, icon *ebiten.Image, iconX, iconY float64) {
	op := &ebiten.DrawImageOptions{}
	pixelX := float64(iconX)
	pixelY := float64(iconY)
	op.GeoM.Translate(pixelX, pixelY)
	screen.DrawImage(icon, op)
}

func merge2Images(img1, img2 *ebiten.Image) *ebiten.Image {
	mergedImage := ebiten.NewImage(config.TileSize, config.TileSize)

	opBase := &ebiten.DrawImageOptions{}
	mergedImage.DrawImage(img1, opBase)

	opOverlay := &ebiten.DrawImageOptions{}

	w, h := img2.Bounds().Dx(), img2.Bounds().Dy()

	// stretch second image to be always a 32x12 and put it on top
	// TODO: A bit strange to have the cheese on top of the top bun... but this is simpler
	sx := float64(config.TileSize) / float64(w)
	sy := 12.0 / float64(h)
	opOverlay.GeoM.Scale(sx, sy)
	mergedImage.DrawImage(img2, opOverlay)

	return mergedImage
}
