package game

import (
	"bytes"
	"image"
	"image/color"
	"log"
	"log/slog"
	"os"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mikelangelon/unibun/assets"
	"github.com/mikelangelon/unibun/config"
	"github.com/mikelangelon/unibun/entities"
	"github.com/mikelangelon/unibun/level"
)

type Game struct {
	levels            []*level.Level
	currentLevelIndex int
	turnManager       turnManager

	gameScreen *ebiten.Image
	status     Status

	// TODO Maybe is not needed? Use status instead?
	needsRestart bool
	shake        *Shake
	// delays
	enemyTurnDelayTimer int
	nextLevelDelayTimer int

	// Menu related fields
	currentGameState   GameState
	menuBackground     *ebiten.Image
	menuOptions        []MenuOption
	selectedMenuOption int

	// Audio players
	audios            audios
	previousGameState GameState // To change music when GameState changes.
}

type turnManager struct {
	currentTurn      int
	turnOrderDisplay []character
}

type character interface {
	Draw(screen *ebiten.Image)
	Update(level entities.Level) bool
	Reset()
}

func (t turnManager) getPlayerType(playerType config.PlayerType) *entities.Player {
	for _, v := range t.turnOrderDisplay {
		switch item := v.(type) {
		case *entities.Player:
			if item.PlayerType == playerType {
				return item
			}
		}
	}
	return nil
}

func (t turnManager) getPlayerTypes(playerType config.PlayerType) []*entities.Player {
	var types []*entities.Player
	for _, v := range t.turnOrderDisplay {
		switch item := v.(type) {
		case *entities.Player:
			if item.PlayerType == playerType {
				types = append(types, item)
			}
		}
	}
	return types
}

func NewGame() *Game {
	g := Game{}
	// init audios
	a, err := newAudios()
	if err != nil {
		log.Fatal(err)
	}
	g.audios = a
	// init active turn
	if len(g.turnManager.turnOrderDisplay) > 0 {
		if p, ok := g.turnManager.turnOrderDisplay[0].(*entities.Player); ok {
			p.IsActiveTurn = true
		}
	}
	img, _, err := image.Decode(bytes.NewReader(assets.MenuBackground))
	if err != nil {
		log.Fatal(err)
	}
	g.menuBackground = ebiten.NewImageFromImage(img)

	g.initLevels()
	g.initMenu()
	g.currentGameState = StateMenu
	g.previousGameState = -1
	return &g
}
func (g *Game) Update() error {
	if g.currentGameState != g.previousGameState {
		g.handleGameStateChange()
		g.previousGameState = g.currentGameState
	}

	switch g.currentGameState {
	case StateMenu:
		return g.updateMenu()
	case StatePlaying, StateRandom:
		return g.updatePlaying()
	case StateExiting:
		os.Exit(0)
	}
	return nil
}

func (g *Game) handleGameStateChange() {
	g.audios.menuMusicPlayer.Pause()
	g.audios.mainMusicPlayer.Pause()

	switch g.currentGameState {
	case StateMenu:
		g.audios.menuMusicPlayer.Rewind()
		g.audios.menuMusicPlayer.Play()
	case StatePlaying, StateRandom:
		g.audios.mainMusicPlayer.Rewind()
		g.audios.mainMusicPlayer.Play()
	}
}
func (g *Game) updatePlaying() error {
	if g.shake != nil {
		g.shake.Update()
	}
	actorEntry := g.turnManager.turnOrderDisplay[0]
	switch actor := actorEntry.(type) {
	case entities.Enemier:
		if g.enemyTurnDelayTimer > 0 {
			g.enemyTurnDelayTimer--
			return nil
		}
		g.checkCollisionToPlayer(actor)
		enemyMoved := actor.Update(g.currentLevel())
		if !enemyMoved {
			break
		}
		g.checkCollisionToPlayer(actor)
		g.advanceTurn()
	case *entities.Player:
		if actor != nil {
			if !actor.Update(g.currentLevel()) {
				break
			}
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

					// Start of step-by-step collision checks
					isBun := actor.PlayerType == config.TopBun || actor.PlayerType == config.BottomBun
					patty := g.currentLevel().BurgerPatty
					if isBun && patty != nil && actor.GridX == patty.GridX && actor.GridY == patty.GridY {
						pattyNextX := patty.GridX + (actor.GridX - oldX)
						pattyNextY := patty.GridY + (actor.GridY - oldY)
						if !g.currentLevel().IsWalkable(pattyNextX, pattyNextY) {
							actor.GridX, actor.GridY = oldX, oldY // Revert move
							moved = false
							break
						} else {
							patty.GridX = pattyNextX
							patty.GridY = pattyNextY
						}
					}

					oldCanDash, oldCanWalk := actor.CanDash, actor.CanWalkThroughWalls
					g.bunCollidesBun(actor)
					g.checkBunCheeseMerge()
					g.checkBunLettuceMerge()
					g.checkCollisionToPlayerOnPlayerTurn(actor)

					// Stop path execution on collision/merge
					if g.needsRestart || g.justMerged(actor, oldCanDash, oldCanWalk) {
						break
					}
				}

				if moved {
					if !g.alreadyMerged() {
						g.attemptMergeBurger()
					} else {
						g.checkWinCondition(g.turnManager.getPlayerType(config.MergedBurgerType))
					}
					g.advanceTurn()
				}
			}
		}
	}
	if g.needsRestart {
		g.Reset()
		return nil
	}
	if g.status == Win {
		if g.nextLevelDelayTimer > 0 {
			g.nextLevelDelayTimer--
			return nil
		}
		if g.currentGameState == StateRandom {
			g.startRandomGame()
		} else {
			g.increaseLevel()
			g.levelToTurn()
		}
	}
	return nil
}

func (g *Game) checkWinCondition(actor *entities.Player) {
	if actor == nil || actor.PlayerType != config.MergedBurgerType {
		return
	}
	for _, v := range g.currentLevel().Winning {
		if actor.GridX == v.X && actor.GridY == v.Y {
			g.status = Win
			g.nextLevelDelayTimer = config.NextLevelDelayDuration
			log.Println("YOU WIN! Merged burger reached the win tile.")
		}
	}
}

func (g *Game) justMerged(p *entities.Player, oldCanDash, oldCanWalk bool) bool {
	if p.PlayerType != config.TopBun && p.PlayerType != config.BottomBun {
		return false
	}
	return (p.CanDash && !oldCanDash) || (p.CanWalkThroughWalls && !oldCanWalk)
}

func (g *Game) initLevels() {
	g.levels = []*level.Level{level.NewLevel0(), level.NewLevel1(), level.NewLevel2(), level.NewLevel3(), level.NewLevel4()}
	g.currentLevelIndex = 0
	g.status = Playing
	g.levelToTurn()
}

func (g *Game) startRandomGame() {
	slog.Info("Starting new random game")
	g.currentGameState = StateRandom
	g.levels = []*level.Level{level.NewRandomLevel()}
	g.currentLevelIndex = 0
	g.status = Playing
	g.levelToTurn()
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
	if len(g.levels) <= g.currentLevelIndex {
		// no more levels :'(
		return
	}
	g.currentLevelIndex++
	g.status = Playing
}
func (g *Game) Draw(screen *ebiten.Image) {
	switch g.currentGameState {
	case StateMenu:
		g.drawMenu(screen)
	case StatePlaying, StateRandom:
		g.drawPlaying(screen)
	}
}
func (g *Game) drawPlaying(screen *ebiten.Image) {
	screen.Fill(color.Black)
	if g.gameScreen == nil {
		g.gameScreen = ebiten.NewImage(g.levels[0].ScreenWidth(), g.levels[0].ScreenHeight())
	}
	g.gameScreen.Clear()

	g.gameScreen.Fill(color.RGBA{R: 0x10, G: 0x10, B: 0x10, A: 0xff})
	g.currentLevel().Draw(g.gameScreen)
	for _, character := range g.turnManager.turnOrderDisplay {
		if character != nil {
			character.Draw(g.gameScreen)
		}
	}
	if g.currentLevel().BurgerPatty != nil {
		g.currentLevel().BurgerPatty.Draw(g.gameScreen)
	}

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
	return g.levels[g.currentLevelIndex]
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

	switch g.turnManager.turnOrderDisplay[0].(type) {
	case *entities.PathEnemy, *entities.Enemy:
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

func (g *Game) attemptMergeBurger() {
	if !g.canBeMerged() {
		return
	}
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
	mergedImage.DrawImage(g.currentLevel().BurgerPatty.Image, op)
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, float64(-10)/2.0)
	mergedImage.DrawImage(topBunPlayer.Image, op)

	mergedPlayer := entities.Player{
		GridX:      g.currentLevel().BurgerPatty.GridX,
		GridY:      g.currentLevel().BurgerPatty.GridY,
		PlayerType: config.MergedBurgerType,
		Image:      mergedImage,
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
	g.currentLevel().BurgerPatty = nil
}

func (g *Game) canBeMerged() bool {
	if g.alreadyMerged() || g.currentLevel().BurgerPatty == nil {
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
	g.shake = newShake(shakeDefaultDuration, shakeDefaultMagnitude)
	for _, char := range g.turnManager.turnOrderDisplay {
		char.Reset()
	}

	var characters []character
	for _, v := range g.currentLevel().TurnOrderPattern {
		switch actualActor := v.(type) {
		case *entities.PathEnemy:
			characters = append(characters, actualActor)
		case *entities.Enemy:
			characters = append(characters, actualActor)
		case entities.Player:
			characters = append(characters, &actualActor)
		}
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
				g.audios.eatingSoundPlayer.Rewind()
				g.audios.eatingSoundPlayer.Play()
				g.needsRestart = true
			}
		}
	}
}

func (g *Game) checkCollisionToPlayerOnPlayerTurn(player *entities.Player) {
	for _, v := range g.turnManager.turnOrderDisplay {
		switch enemy := v.(type) {
		case entities.Enemier:
			if enemy.Collision(player) {
				g.audios.eatingSoundPlayer.Rewind()
				g.audios.eatingSoundPlayer.Play()
				g.needsRestart = true
			}
		}
	}
}

func (g *Game) levelToTurn() {
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
}

func merge2Images(img1, img2 *ebiten.Image) *ebiten.Image {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, float64(0)/2.0)
	mergedImage := ebiten.NewImage(config.TileSize, config.TileSize)
	mergedImage.DrawImage(img1, op)
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, float64(-10)/2.0)
	mergedImage.DrawImage(img2, op)
	return mergedImage
}
