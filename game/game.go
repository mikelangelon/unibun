package game

import (
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"log"
	"log/slog"

	"github.com/hajimehoshi/ebiten/v2"
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
}

type turnManager struct {
	currentTurn      int
	turnOrderDisplay []character
}

type character interface {
	Draw(screen *ebiten.Image)
	Update(level entities.Level) bool
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

func NewGame() *Game {
	g := Game{
		levels: []*level.Level{level.NewLevel0(), level.NewLevel1(), level.NewLevel2()},
		turnManager: turnManager{
			currentTurn: 0,
		},
		status: Playing,
	}
	g.levelToTurn()
	return &g
}

func (g *Game) Update() error {
	if g.shake != nil {
		g.shake.Update()
	}
	actorEntry := g.turnManager.turnOrderDisplay[0]
	switch actor := actorEntry.(type) {
	case *entities.Enemy:
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
			oldX, oldY := actor.GridX, actor.GridY
			playedMoved := actor.Update(g.currentLevel())
			if !playedMoved {
				break
			}
			isBun := actor.PlayerType == config.TopBun || actor.PlayerType == config.BottomBun
			patty := g.currentLevel().BurgerPatty
			if isBun && patty != nil && actor.GridX == patty.GridX && actor.GridY == patty.GridY {
				// A bun is trying to move into the patty's space. Try to push it.
				pattyNextX := patty.GridX + (actor.GridX - oldX)
				pattyNextY := patty.GridY + (actor.GridY - oldY)
				canPattyMove := g.currentLevel().IsWalkable(pattyNextX, pattyNextY)
				if !canPattyMove {
					playedMoved = false
					actor.GridX, actor.GridY = oldX, oldY
					break
				} else {
					// TODO Check if patty's next spot is occupied by other players, enemies...
					patty.GridX = pattyNextX
					patty.GridY = pattyNextY
				}
			}
			if !g.alreadyMerged() {
				g.attemptMergeBurger()
			} else {
				if actor.PlayerType == config.MergedBurgerType {
					for _, v := range g.currentLevel().Winning {
						if actor.GridX == v.X && actor.GridY == v.Y {
							g.status = Win
							g.nextLevelDelayTimer = config.NextLevelDelayDuration
							log.Println("YOU WIN! Merged burger reached the win tile.")
						}
					}

				}
			}
			if g.status == Playing {
				g.advanceTurn()
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
		} else {
			g.increaseLevel()
			g.levelToTurn()
		}
	}
	return nil
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
		case *entities.Enemy:
			drawIcon(screen, item.Image, iconX, iconY)
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
	firstElement := g.turnManager.turnOrderDisplay[0]
	g.turnManager.turnOrderDisplay = append(g.turnManager.turnOrderDisplay[1:], firstElement)
	switch g.turnManager.turnOrderDisplay[0].(type) {
	case *entities.Enemy:
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
	var characters []character
	for _, v := range g.currentLevel().TurnOrderPattern {
		switch actualActor := v.(type) {
		case *entities.Enemy:
			characters = append(characters, actualActor)
		case entities.Player:
			characters = append(characters, &actualActor)
		}
	}
	g.turnManager.turnOrderDisplay = characters
}

func (g *Game) checkCollisionToPlayer(enemy *entities.Enemy) {
	for _, v := range g.turnManager.turnOrderDisplay {
		switch player := v.(type) {
		case *entities.Player:
			if enemy.Collision(player) {
				g.needsRestart = true
			}
		}
	}
}
func (g *Game) levelToTurn() {
	var characters []character
	for _, v := range g.currentLevel().TurnOrderPattern {
		switch actualActor := v.(type) {
		case *entities.Enemy:
			characters = append(characters, actualActor)
		case entities.Player:
			characters = append(characters, &actualActor)
		}
	}
	g.turnManager.turnOrderDisplay = characters
}
