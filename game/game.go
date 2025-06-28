package game

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"log"
	"log/slog"
	"os"
	"strings"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mikelangelon/unibun/assets"
	"github.com/mikelangelon/unibun/config"
	"github.com/mikelangelon/unibun/entities"
	"github.com/mikelangelon/unibun/level"
)

type Game struct {
	level               *level.Level
	currentLevelIndex   int
	currentEndlessLevel int
	turnManager         turnManager
	patty               *entities.BurgerPatty

	gameScreen *ebiten.Image
	status     Status

	// TODO Maybe is not needed? Use status instead?
	needsRestart bool
	shake        *Shake
	// delays
	enemyTurnDelayTimer int
	introDelayTimer     int

	// Menu related fields
	currentGameState        GameState
	menuBackground          *ebiten.Image
	menuOptions             []MenuOption
	selectedMenuOption      int
	pauseMenuOptions        []MenuOption
	selectedPauseMenuOption int

	resetTimer int
	// Audio players
	audios            audios
	previousGameState GameState // To change music when GameState changes.
	stateBeforePause  GameState

	animationManager *animationManager

	// To select level
	levelConstructors map[int]func() *level.Level
	selectedLevelBox  int
}

type character interface {
	Draw(screen *ebiten.Image)
	Update(level entities.Level) bool
	Reset()
}

func NewGame() *Game {
	g := Game{}
	// init audios
	a, err := newAudios()
	if err != nil {
		log.Fatal(err)
	}
	g.audios = a
	img, _, err := image.Decode(bytes.NewReader(assets.MenuBackground))
	if err != nil {
		log.Fatal(err)
	}
	g.menuBackground = ebiten.NewImageFromImage(img)

	g.initMenu()
	g.initPauseMenu()
	g.initLevelConstructors()
	g.currentGameState = StateMenu
	g.selectedLevelBox = 0
	g.previousGameState = -1
	g.resetTimer = 0
	g.animationManager = newAnimationManager()
	g.introDelayTimer = 0

	return &g
}
func (g *Game) initLevelConstructors() {
	g.levelConstructors = map[int]func() *level.Level{
		1: level.NewLevel1,
		2: level.NewLevel2,
		3: level.NewLevel3,
		4: level.NewLevel4,
	}
}

func (g *Game) initPauseMenu() {
	g.pauseMenuOptions = []MenuOption{
		{
			Text: "Continue",
			Action: func(game *Game) {
				game.currentGameState = game.stateBeforePause
			},
		},
		{
			Text: "Restart",
			Action: func(game *Game) {
				game.Reset()
				game.currentGameState = game.stateBeforePause
			},
		},
		{
			Text: "Menu",
			Action: func(game *Game) {
				game.currentGameState = StateMenu
			},
		},
	}

	totalMenuHeight := config.MenuOptionHeight*3 + config.MenuOptionSpacing*2
	startY := (config.WindowHeight - totalMenuHeight) / 2
	centerX := config.WindowWidth / 2

	for i := range g.pauseMenuOptions {
		option := &g.pauseMenuOptions[i]
		option.Rect = image.Rect(
			centerX-config.MenuOptionWidth/2,
			startY+i*(config.MenuOptionHeight+config.MenuOptionSpacing),
			centerX+config.MenuOptionWidth/2,
			startY+i*(config.MenuOptionHeight+config.MenuOptionSpacing)+config.MenuOptionHeight,
		)
	}
	g.selectedPauseMenuOption = 0
}
func (g *Game) Update() error {
	if g.currentGameState != g.previousGameState {
		g.handleGameStateChange()
		g.previousGameState = g.currentGameState
	}

	switch g.currentGameState {
	case StateMenu:
		return g.updateMenu()
	case StatePlaying, StateEndless:
		return g.updatePlaying()
	case StateLevelSelect:
		return g.updateLevelSelect()
	case StateIntro:
		return g.updateIntro()
	case StatePaused:
		return g.updatePaused()
	case StateExiting:
		os.Exit(0)
	}
	return nil
}

func (g *Game) handleGameStateChange() {
	if g.currentGameState == StatePaused {
		return
	}
	if g.previousGameState == StatePaused && g.currentGameState != StateMenu {
		return
	}
	if g.previousGameState == StateIntro && g.currentGameState == StatePlaying {
		return
	}
	if g.previousGameState == StateMenu && g.currentGameState == StateLevelSelect {
		return
	}
	g.audios.menuMusicPlayer.Pause()
	g.audios.mainMusicPlayer.Pause()

	switch g.currentGameState {
	case StateMenu, StateLevelSelect:
		g.audios.menuMusicPlayer.Rewind()
		g.audios.menuMusicPlayer.Play()
	case StateIntro, StatePlaying, StateEndless:
		g.introDelayTimer = 10
		g.audios.menuMusicPlayer.Rewind()
		g.audios.menuMusicPlayer.Play()
	}
}

func (g *Game) updateLevelSelect() error {
	cols := 5
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) && (g.selectedLevelBox+1)%cols != 0 && g.selectedLevelBox < 14 {
		g.selectedLevelBox++
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) && g.selectedLevelBox%cols != 0 {
		g.selectedLevelBox--
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) && g.selectedLevelBox >= cols {
		g.selectedLevelBox -= cols
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) && g.selectedLevelBox < 10 {
		g.selectedLevelBox += cols
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.currentLevelIndex = g.selectedLevelBox + 1
		g.startLevel(g.currentLevelIndex)
	}
	return nil
}

func (g *Game) updatePaused() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		g.selectedPauseMenuOption--
		if g.selectedPauseMenuOption < 0 {
			g.selectedPauseMenuOption = len(g.pauseMenuOptions) - 1
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		g.selectedPauseMenuOption++
		if g.selectedPauseMenuOption >= len(g.pauseMenuOptions) {
			g.selectedPauseMenuOption = 0
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.pauseMenuOptions[g.selectedPauseMenuOption].Action(g)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.currentGameState = g.previousGameState
	}
	return nil
}

func (g *Game) updateIntro() error {
	if g.introDelayTimer > 0 {
		g.introDelayTimer--
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.currentGameState = StatePlaying
	}
	return nil
}

func (g *Game) updatePlaying() error {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.stateBeforePause = g.currentGameState
		g.currentGameState = StatePaused
		return nil
	}

	if g.shake != nil {
		g.shake.Update()
	}

	if g.animationManager.isMergeAnimationPlaying() {
		g.updateMergeAnimation()
		return nil
	}
	if g.animationManager.isMergeAnimationPlaying() {
		g.updateMergeAnimation()
		return nil
	}
	if g.animationManager.isWinningPlaying() {
		g.updateWinAnimation()
		return nil
	}

	g.updateEffects()

	if g.needsRestart {
		g.resetTimer--
		if g.resetTimer <= 0 {
			g.Reset()
		}
		return nil
	}

	actorEntry := g.turnManager.turnOrderDisplay[0]
	switch actor := actorEntry.(type) {
	case entities.Enemier:
		g.handleEnemyTurn(actor)
	case *entities.Player:
		if actor == nil {
			break
		}

		// If player is in the middle of a dash, continue it.
		if actor.IsDashing() {
			actor.Update(g.currentLevel())
			g.checkCollisionToPlayerOnPlayerTurn(actor)
			if g.needsRestart {
				return nil
			}
			break
		}

		actor.Update(g.currentLevel())

		// TODO Too many if statements. Try to fix
		dx, dy, isMoving, isDashing := actor.GetMoveInput()
		if !isMoving {
			break
		}

		if isDashing {
			if actor.CanDash {
				// The dash is initiated -->  Update() will keep to move in next turns
				if actor.StartDash(g.currentLevel(), dx, dy) {
					g.advanceTurn()
				}
			}
		} else {
			path := actor.CalculateMovePath(g.currentLevel(), dx, dy)
			if len(path) == 0 {
				break
			}

			moved := false
			for _, point := range path {
				oldX, oldY := actor.GridX, actor.GridY
				actor.GridX, actor.GridY = point.X, point.Y
				moved = true

				pushedPatty := false
				// Start of step-by-step collision checks
				isBun := actor.PlayerType == config.TopBun || actor.PlayerType == config.BottomBun
				if isBun && g.patty != nil && actor.GridX == g.patty.GridX && actor.GridY == g.patty.GridY {
					pattyNextX := g.patty.GridX + (actor.GridX - oldX)
					pattyNextY := g.patty.GridY + (actor.GridY - oldY)
					if !g.currentLevel().IsWalkable(pattyNextX, pattyNextY) || g.isTileOccupiedByCharacter(pattyNextX, pattyNextY) {
						actor.GridX, actor.GridY = oldX, oldY // Revert move
						moved = true
						break // Stop path execution
					} else {
						g.patty.GridX = pattyNextX
						g.patty.GridY = pattyNextY
						pushedPatty = true
					}
				}

				oldCanDash, oldCanWalk := actor.CanDash, actor.CanWalkThroughWalls
				g.bunCollidesBun(actor)
				g.checkBunCheeseMerge()
				g.checkBunLettuceMerge()
				g.checkCollisionToPlayerOnPlayerTurn(actor)

				// Stop path execution on collision/merge or after pushing patty
				if g.needsRestart || g.justMerged(actor, oldCanDash, oldCanWalk) || pushedPatty {
					break
				}
			}

			if moved {
				if !g.alreadyMerged() {
					g.attemptMergeBurger()
				} else {
					g.checkWinCondition(g.turnManager.getPlayerType(config.MergedBurgerType))
				}
				if !g.animationManager.blockingAnimation() {
					g.advanceTurn()
				}
			}
		}
	}
	return nil
}

func (g *Game) updateEffects() {
	g.animationManager.Update()
}

func (g *Game) isTileOccupiedByCharacter(x, y int) bool {
	for _, char := range g.turnManager.turnOrderDisplay {
		switch entity := char.(type) {
		case *entities.Player:
			if entity.GridX == x && entity.GridY == y {
				return true
			}
		case entities.Enemier:
			dummyPlayer := &entities.Player{GridX: x, GridY: y}
			if entity.Collision(dummyPlayer) {
				return true
			}
		}
	}
	return false
}
func (g *Game) updateMergeAnimation() {
	if !g.animationManager.isMergeAnimationPlaying() {
		return
	}
	g.animationManager.Update()

	if !g.animationManager.isMergeAnimationPlaying() {
		g.performMerge()
		g.advanceTurn()
	}
}

func (g *Game) updateWinAnimation() {
	if !g.animationManager.isWinningPlaying() {
		return
	}
	g.animationManager.Update()

	if !g.animationManager.isWinningPlaying() {
		if g.currentGameState == StateEndless {
			g.currentEndlessLevel++
			g.startEndlessGame() // This reloads the endless mode
		} else {
			g.increaseLevel()
			g.startLevel(g.currentLevelIndex)
		}
	}
}

func (g *Game) startLevel(levelNum int) {
	constructor, ok := g.levelConstructors[levelNum]
	if !ok {
		// Plan B for now
		constructor = level.NewLevel0
	}
	g.level = constructor()
	g.status = Playing
	g.levelToTurn()

	if g.currentLevel().IntroText != "" {
		g.currentGameState = StateIntro
	} else {
		g.currentGameState = StatePlaying
	}
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

	overlayColor := color.RGBA{0, 0, 0, 50}
	ebitenutil.DrawRect(screen, 0, 0, float64(config.WindowWidth), float64(config.WindowHeight), overlayColor)

	boxWidth := 400
	boxHeight := 150
	boxX := (config.WindowWidth - boxWidth) / 2
	boxY := (config.WindowHeight - boxHeight) / 2
	boxColor := color.RGBA{0x20, 0x20, 0x20, 50}
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

// draws 15 boxes, in a 3x5 grid.
func (g *Game) drawLevelSelect(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 20, G: 20, B: 20, A: 255})
	ebitenutil.DebugPrintAt(screen, "Select a Level", config.WindowWidth/2-70, 40)

	const (
		cols        = 5
		rows        = 3
		boxSize     = 80
		padding     = 20
		totalWidth  = cols*boxSize + (cols-1)*padding
		totalHeight = rows*boxSize + (rows-1)*padding
		startX      = (config.WindowWidth - totalWidth) / 2
		startY      = (config.WindowHeight - totalHeight) / 2
	)

	for i := 0; i < 15; i++ {
		col := i % cols
		row := i / cols
		boxX := startX + col*(boxSize+padding)
		boxY := startY + row*(boxSize+padding)

		boxColor := color.RGBA{0x40, 0x40, 0x40, 0xFF}
		if i == g.selectedLevelBox {
			boxColor = color.RGBA{0x90, 0x90, 0x90, 0xFF}
		}
		ebitenutil.DrawRect(screen, float64(boxX), float64(boxY), float64(boxSize), float64(boxSize), boxColor)

		levelNumStr := fmt.Sprintf("%d", i+1)
		textX := boxX + (boxSize-len(levelNumStr)*6)/2
		textY := boxY + (boxSize-16)/2
		ebitenutil.DebugPrintAt(screen, levelNumStr, textX, textY)
	}
}

func (g *Game) drawEffects(screen *ebiten.Image) {
	g.animationManager.drawEffects(screen)
}

func (g *Game) handleEnemyTurn(enemy entities.Enemier) {
	if g.enemyTurnDelayTimer > 0 {
		g.enemyTurnDelayTimer--
		return
	}
	if fe, ok := enemy.(*entities.FollowerEnemy); ok {
		g.setFollowerTarget(fe)
	} else if dfe, ok := enemy.(*entities.DashingFollowerEnemy); ok {
		g.setFollowerTarget(&dfe.FollowerEnemy)
	}

	g.checkCollisionToPlayer(enemy)
	turnConsumed := enemy.Update(g.currentLevel())
	if turnConsumed {
		g.checkCollisionToPlayer(enemy)
		g.advanceTurn()
	}
}

func (g *Game) setFollowerTarget(fe *entities.FollowerEnemy) {
	targetPlayer := g.turnManager.getPlayerType(fe.GetTargetPlayerType())
	gridX, gridY := -1, -1
	if targetPlayer != nil {
		gridX, gridY = targetPlayer.GridX, targetPlayer.GridY
	}
	fe.SetTarget(gridX, gridY)
}

func (g *Game) checkWinCondition(actor *entities.Player) {
	if g.animationManager.isWinningPlaying() || actor == nil || actor.PlayerType != config.MergedBurgerType {
		return
	}
	for _, v := range g.currentLevel().Winning {
		if actor.GridX == v.X && actor.GridY == v.Y {
			g.startWinAnimation(actor.GridX, actor.GridY)
			break
		}
	}
}

func (g *Game) startWinAnimation(gridX, gridY int) {
	g.animationManager.playWinningAnimation(gridX, gridY)
	if !g.animationManager.isWinningPlaying() {
		g.status = Win
		log.Println("YOU WIN! Merged burger reached the win tile.")
	}
}

func (g *Game) justMerged(p *entities.Player, oldCanDash, oldCanWalk bool) bool {
	if p.PlayerType != config.TopBun && p.PlayerType != config.BottomBun {
		return false
	}
	return (p.CanDash && !oldCanDash) || (p.CanWalkThroughWalls && !oldCanWalk)
}

func (g *Game) startEndlessGame() {
	slog.Info("Starting new endless game")
	g.currentGameState = StateEndless
	g.level = level.NewEndlessLevel(g.currentEndlessLevel)
	g.startLevel(0)
}

func (g *Game) bunCollidesBun(player *entities.Player) {
	if player.PlayerType != config.TopBun && player.PlayerType != config.BottomBun {
		return
	}

	otherBunType := config.BottomBun
	if player.PlayerType == config.BottomBun {
		otherBunType = config.TopBun
	}
	otherBun := g.turnManager.getPlayerType(otherBunType)
	if otherBun != nil && player.GridX == otherBun.GridX && player.GridY == otherBun.GridY {
		log.Println("Buns collided!")
		g.needsRestart = true
	}
}

// checkBunCheeseMerge checks if bun collides to cheese, or the other way around
func (g *Game) checkBunCheeseMerge() {
	cheeses := g.turnManager.getPlayerTypes(config.Cheese)
	if len(cheeses) == 0 {
		return
	}
	topBun := g.turnManager.getPlayerType(config.TopBun)
	bottomBun := g.turnManager.getPlayerType(config.BottomBun)
	for _, v := range cheeses {
		if topBun != nil && v.CollisionTo(topBun.GridX, topBun.GridY) {
			g.cheesePower(topBun, v)
		}
		if bottomBun != nil && v.CollisionTo(bottomBun.GridX, bottomBun.GridY) {
			g.cheesePower(bottomBun, v)
		}
	}
}

func (g *Game) cheesePower(bun, cheese *entities.Player) {
	log.Println("Bun and Cheese merged! Bun can now dash.")
	bun.CanDash = true
	bun.Image = merge2Images(bun.Image, cheese.Image)
	var newTurnOrder []character
	for _, char := range g.turnManager.turnOrderDisplay {
		if char == cheese {
			continue
		}
		newTurnOrder = append(newTurnOrder, char)
	}
	g.turnManager.turnOrderDisplay = newTurnOrder
}

func (g *Game) checkBunLettuceMerge() {
	lettucePlayers := g.turnManager.getPlayerTypes(config.Lettuce)
	if len(lettucePlayers) == 0 {
		return
	}

	topBun := g.turnManager.getPlayerType(config.TopBun)
	bottomBun := g.turnManager.getPlayerType(config.BottomBun)

	for _, lettuce := range lettucePlayers {
		if topBun != nil && topBun.CollisionTo(lettuce.GridX, lettuce.GridY) {
			g.lettucePower(topBun, lettuce)
			return
		}
		if bottomBun != nil && bottomBun.CollisionTo(lettuce.GridX, lettuce.GridY) {
			g.lettucePower(bottomBun, lettuce)
			return
		}
	}
}

func (g *Game) lettucePower(bun, lettuce *entities.Player) {
	slog.Info("Bun and Lettuce merged! Bun can now walk through walls")
	bun.CanWalkThroughWalls = true // Grant power of walking on through the walls
	bun.Image = merge2Images(bun.Image, lettuce.Image)

	var newTurnOrder []character
	for _, char := range g.turnManager.turnOrderDisplay {
		if char == lettuce {
			continue
		}
		newTurnOrder = append(newTurnOrder, char)
	}
	g.turnManager.turnOrderDisplay = newTurnOrder
}

func (g *Game) increaseLevel() {
	if g.currentLevelIndex >= 15 {
		// TODO Show an ending
	}
	g.currentLevelIndex++
	g.status = Playing
}
func (g *Game) Draw(screen *ebiten.Image) {
	switch g.currentGameState {
	case StateMenu:
		g.drawMenu(screen)
	case StateLevelSelect:
		g.drawLevelSelect(screen)
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

func (g *Game) drawPaused(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, 0, 0, float64(config.WindowWidth), float64(config.WindowHeight), color.RGBA{0, 0, 0, 128})

	for i, option := range g.pauseMenuOptions {
		btnColor := color.RGBA{0x40, 0x40, 0x40, 0xFF}
		if i == g.selectedPauseMenuOption {
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

func (g *Game) drawPlaying(screen *ebiten.Image) {
	screen.Fill(color.Black)
	if g.gameScreen == nil {
		g.gameScreen = ebiten.NewImage(g.level.ScreenWidth(), g.level.ScreenHeight())
	}
	g.gameScreen.Clear()

	g.gameScreen.Fill(color.RGBA{R: 0x10, G: 0x10, B: 0x10, A: 0xff})
	g.currentLevel().Draw(g.gameScreen)
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

// Constants for Turn Order UI
const (
	turnOrderIconSize      = 24
	turnOrderIconSpacing   = 6
	turnOrderTextMarginX   = 10
	turnOrderTextOffsetY   = 4
	turnOrderIconTopMargin = 18
)

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
			drawIcon(screen, item.Image(), iconX, iconY)
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
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return config.WindowWidth, config.WindowHeight
}

func (g *Game) currentLevel() *level.Level {
	return g.level
}

func (g *Game) buildTurnOrderDisplay() {
	if len(g.turnManager.turnOrderDisplay) == 0 {
		return
	}

	if oldActor, ok := g.turnManager.turnOrderDisplay[0].(*entities.Player); ok {
		oldActor.IsActiveTurn = false
	}

	firstElement := g.turnManager.turnOrderDisplay[0]
	g.turnManager.turnOrderDisplay = append(g.turnManager.turnOrderDisplay[1:], firstElement)

	if newActor, ok := g.turnManager.turnOrderDisplay[0].(*entities.Player); ok {
		newActor.IsActiveTurn = true
	}

	if _, ok := g.turnManager.turnOrderDisplay[0].(entities.Enemier); ok {
		g.enemyTurnDelayTimer = config.EnemyTurnDelayDuration
	}
}

// advanceTurn moves to the next actor in the turn pattern.
func (g *Game) advanceTurn() {
	currentLvl := g.currentLevel()
	if len(currentLvl.TurnOrderPattern) == 0 {
		log.Println("Warning: Current level has no turn order pattern defined.")
		return
	}

	g.turnManager.currentTurn = (g.turnManager.currentTurn + 1) % len(currentLvl.TurnOrderPattern)
	g.buildTurnOrderDisplay()
}

func (g *Game) performMerge() {
	slog.Info("Burger components united!")
	mergedImage := ebiten.NewImage(config.TileSize, config.TileSize)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, float64(10)/2.0)
	topBunPlayer := g.turnManager.getPlayerType(config.TopBun)
	bottomBunPlayer := g.turnManager.getPlayerType(config.BottomBun)

	// Draw in visual stack order: bottom, patty, top
	mergedImage.DrawImage(bottomBunPlayer.Image, op)
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, float64(0)/2.0)
	mergedImage.DrawImage(g.patty.Image, op)
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, float64(-10)/2.0)
	mergedImage.DrawImage(topBunPlayer.Image, op)

	mergedPlayer := entities.Player{
		GridX:               g.patty.GridX,
		GridY:               g.patty.GridY,
		PlayerType:          config.MergedBurgerType,
		Speed:               1,
		Image:               mergedImage,
		CanDash:             topBunPlayer.CanDash || bottomBunPlayer.CanDash,
		CanWalkThroughWalls: topBunPlayer.CanWalkThroughWalls || bottomBunPlayer.CanWalkThroughWalls,
	}
	var charactersWithoutMergedOnes []character
	for _, v := range g.turnManager.turnOrderDisplay {
		if v == topBunPlayer || v == bottomBunPlayer {
			continue
		}
		charactersWithoutMergedOnes = append(charactersWithoutMergedOnes, v)
	}
	charactersWithoutMergedOnes = append(charactersWithoutMergedOnes, &mergedPlayer)
	g.turnManager.turnOrderDisplay = charactersWithoutMergedOnes
	g.patty = nil
}

func (g *Game) attemptMergeBurger() {
	if g.animationManager.isMergeAnimationPlaying() || !g.canBeMerged() {
		return
	}
	topBun := g.turnManager.getPlayerType(config.TopBun)
	bottomBun := g.turnManager.getPlayerType(config.BottomBun)
	// Start animation
	g.animationManager.playMergeAnimation(g.patty, topBun, bottomBun)
}

func (g *Game) canBeMerged() bool {
	if g.alreadyMerged() || g.patty == nil {
		// Already merged, or components missing for a merge
		return false
	}

	topBunPlayer := g.turnManager.getPlayerType(config.TopBun)
	bottomBunPlayer := g.turnManager.getPlayerType(config.BottomBun)
	patty := g.currentLevel().BurgerPatty

	merged := false
	// Corrected: TopBun-Patty-BottomBun or BottomBun-Patty-TopBun
	if topBunPlayer.GridY == patty.GridY && patty.GridY == bottomBunPlayer.GridY { // Same row
		// Case 1: TopBun, Patty, BottomBun
		if topBunPlayer.GridX == patty.GridX-1 && bottomBunPlayer.GridX == patty.GridX+1 {
			merged = true
		}
		// Case 2: BottomBun, Patty, TopBun
		if bottomBunPlayer.GridX == patty.GridX-1 && topBunPlayer.GridX == patty.GridX+1 {
			merged = true
		}
	}
	// Check vertical alignment: TopBun-Patty-BottomBun or BottomBun-Patty-TopBun
	if topBunPlayer.GridX == patty.GridX && patty.GridX == bottomBunPlayer.GridX { // Same column
		// Case 1: TopBun above, Patty, BottomBun below
		if topBunPlayer.GridY == patty.GridY-1 && bottomBunPlayer.GridY == patty.GridY+1 {
			merged = true
		}
		// Case 2: BottomBun above, Patty, TopBun below
		// For now... not allowed, it's weird to have the burguer the other way around
	}
	return merged
}

func (g *Game) alreadyMerged() bool {
	return g.turnManager.getPlayerType(config.MergedBurgerType) != nil
}

func (g *Game) Reset() {
	log.Println("Game Over! Restarting...")
	g.needsRestart = false
	g.resetTimer = 0
	g.shake = newShake(shakeDefaultDuration, shakeDefaultMagnitude)
	g.patty = &g.currentLevel().BurgerPatty
	if g.patty != nil {
		g.patty.Reset()
	}

	var characters []character
	for _, v := range g.currentLevel().TurnOrderPattern {
		switch actualActor := v.(type) {
		case *entities.PathEnemy:
			characters = append(characters, actualActor)
		case *entities.Enemy:
			characters = append(characters, actualActor)
		case *entities.FollowerEnemy:
			characters = append(characters, actualActor)
		case *entities.DashingFollowerEnemy:
			characters = append(characters, actualActor)
		case entities.Player:
			characters = append(characters, &actualActor)
		}
	}
	for _, char := range characters {
		char.Reset()
	}
	g.turnManager.turnOrderDisplay = characters

	if len(g.turnManager.turnOrderDisplay) > 0 {
		if p, ok := g.turnManager.turnOrderDisplay[0].(*entities.Player); ok {
			p.IsActiveTurn = true
		}
	}
}

func (g *Game) checkCollisionToPlayer(enemy entities.Enemier) {
	for _, v := range g.turnManager.turnOrderDisplay {
		switch player := v.(type) {
		case *entities.Player:
			if enemy.Collision(player) {
				if !g.needsRestart {
					g.shake = newShake(shakeDefaultDuration, shakeDefaultMagnitude)
					g.resetTimer = shakeDefaultDuration + 10 // delay a bit
				}
				g.audios.eatingSoundPlayer.Rewind()
				g.audios.eatingSoundPlayer.Play()
				g.needsRestart = true
			}
		}
	}
}

func (g *Game) checkCollisionToPlayerOnPlayerTurn(player *entities.Player) {
	for i := len(g.turnManager.turnOrderDisplay) - 1; i >= 0; i-- {
		v := g.turnManager.turnOrderDisplay[i]
		if v == player {
			continue
		}

		switch enemy := v.(type) {
		case entities.Enemier:
			if enemy.Collision(player) {
				isBun := player.PlayerType == config.TopBun || player.PlayerType == config.BottomBun
				if isBun && player.IsDashing() {
					g.turnManager.turnOrderDisplay = append(g.turnManager.turnOrderDisplay[:i], g.turnManager.turnOrderDisplay[i+1:]...)
					g.animationManager.playKillEffect(player.GridX, player.GridY)
				} else if !g.needsRestart {
					g.shake = newShake(shakeDefaultDuration, shakeDefaultMagnitude)
					g.resetTimer = shakeDefaultDuration + 10
					g.audios.eatingSoundPlayer.Rewind()
					g.audios.eatingSoundPlayer.Play()
					g.needsRestart = true
					return
				}
			}
		}
	}
}

func (g *Game) levelToTurn() {
	g.patty = &g.currentLevel().BurgerPatty
	var characters []character
	for _, v := range g.currentLevel().TurnOrderPattern {
		switch actualActor := v.(type) {
		case entities.Enemier:
			characters = append(characters, actualActor)
		case entities.Player:
			characters = append(characters, &actualActor)
		}
	}
	g.turnManager.turnOrderDisplay = characters
	// init active turn
	if len(g.turnManager.turnOrderDisplay) > 0 {
		if p, ok := g.turnManager.turnOrderDisplay[0].(*entities.Player); ok {
			p.IsActiveTurn = true
		}
	}
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
