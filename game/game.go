package game

import (
	"image"
	"log"
	"log/slog"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mikelangelon/unibun/config"
	"github.com/mikelangelon/unibun/entities"
	"github.com/mikelangelon/unibun/level"
)

type Game struct {
	level               *level.Level
	currentEndlessLevel int
	turnManager         turnManager
	patty               *entities.BurgerPatty

	gameScreen *ebiten.Image
	status     Status

	// TODO Maybe is not needed? Use status instead?
	needsRestart bool
	shake        *shake
	// delays
	enemyTurnDelayTimer int
	introDelayTimer     int

	// Menu related fields
	currentGameState        GameState
	menu                    *menu
	pauseMenuOptions        []menuOption
	selectedPauseMenuOption int

	resetTimer int
	// Audio players
	audios            audios
	previousGameState GameState // To change music when GameState changes.
	stateBeforePause  GameState

	animationManager *animationManager

	// To select level
	levelManager              *levelManager
	gameCompleteConfettiTimer int
}

type character interface {
	Draw(screen *ebiten.Image)
	Update(level entities.Level) bool
	Reset()
}

func NewGame() *Game {
	g := Game{}
	// init audios
	g.audios, _ = newAudios()
	g.menu = newMenu()
	g.initPauseMenu()
	g.levelManager = newLevelManager(g.startLevel)
	g.currentGameState = StateMenu

	g.previousGameState = -1
	g.resetTimer = 0
	g.animationManager = newAnimationManager()
	g.introDelayTimer = 0
	g.gameCompleteConfettiTimer = 0

	return &g
}

func (g *Game) initPauseMenu() {
	g.pauseMenuOptions = []menuOption{
		{
			text: "Continue",
			action: func(game *Game) {
				game.currentGameState = game.stateBeforePause
			},
		},
		{
			text: "Restart",
			action: func(game *Game) {
				game.Reset()
				game.currentGameState = game.stateBeforePause
			},
		},
		{
			text: "Menu",
			action: func(game *Game) {
				game.currentGameState = StateMenu
			},
		},
	}

	totalMenuHeight := config.MenuOptionHeight*3 + config.MenuOptionSpacing*2
	startY := (config.WindowHeight - totalMenuHeight) / 2
	centerX := config.WindowWidth / 2

	for i := range g.pauseMenuOptions {
		option := &g.pauseMenuOptions[i]
		option.rect = image.Rect(
			centerX-config.MenuOptionWidth/2,
			startY+i*(config.MenuOptionHeight+config.MenuOptionSpacing),
			centerX+config.MenuOptionWidth/2,
			startY+i*(config.MenuOptionHeight+config.MenuOptionSpacing)+config.MenuOptionHeight,
		)
	}
	g.selectedPauseMenuOption = 0
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
	case StateMenu, StateLevelSelect, StateTutorial:
		g.audios.menuMusicPlayer.Rewind()
		g.audios.menuMusicPlayer.Play()
	case StateIntro, StatePlaying, StateEndless:
		g.introDelayTimer = 10
		g.audios.mainMusicPlayer.Rewind()
		g.audios.mainMusicPlayer.Play()
	}
}

// TODO: Fix ugly return of 3 params
// checkBunPatty returns 3 params: moved, pusedPatty, stopPath
func (g *Game) checkBunPatty(actor *entities.Player, oldX, oldY int) (bool, bool, bool) {
	moved := true

	pushedPatty := false
	// Start of step-by-step collision checks
	isBun := actor.PlayerType == config.TopBun || actor.PlayerType == config.BottomBun
	if isBun && g.patty != nil && actor.GridX == g.patty.GridX && actor.GridY == g.patty.GridY {
		pattyNextX := g.patty.GridX + (actor.GridX - oldX)
		pattyNextY := g.patty.GridY + (actor.GridY - oldY)
		if !g.currentLevel().IsWalkable(pattyNextX, pattyNextY) || g.isTileOccupiedByCharacter(pattyNextX, pattyNextY) {
			actor.GridX, actor.GridY = oldX, oldY // Revert move
			moved = true
			return moved, false, true
		} else {
			g.patty.GridX = pattyNextX
			g.patty.GridY = pattyNextY
			pushedPatty = true
		}
	}
	return moved, pushedPatty, false
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

func (g *Game) startLevel(levelNum int) {
	constructor, ok := g.levelManager.levelConstructors[levelNum]
	if !ok {
		// All filled, this should not happen
		constructor = level.NewEmptyLevel
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

func (g *Game) handleEnemyTurn(enemy entities.Enemier) {
	if g.enemyTurnDelayTimer > 0 {
		g.enemyTurnDelayTimer--
		return
	}
	if fe, ok := enemy.(*entities.Duck); ok {
		g.setFollowerTarget(fe)
	} else if dfe, ok := enemy.(*entities.Snake); ok {
		g.setFollowerTarget(&dfe.Duck)
	} else if dfe, ok := enemy.(*entities.Fly); ok {
		g.setFlyTarget(dfe)
	}

	g.checkCollisionToPlayer(enemy)
	turnConsumed := enemy.Update(g.currentLevel())
	if turnConsumed {
		g.checkCollisionToPlayer(enemy)
		g.checkCollisionToPatty(enemy)
		g.advanceTurn()
	}
}

func (g *Game) setFollowerTarget(fe *entities.Duck) {
	targetPlayer := g.turnManager.getPlayerType(fe.GetTargetPlayerType())
	gridX, gridY := -1, -1
	if targetPlayer != nil {
		gridX, gridY = targetPlayer.GridX, targetPlayer.GridY
	}
	fe.SetTarget(gridX, gridY)
}

func (g *Game) setFlyTarget(f *entities.Fly) {
	if g.patty != nil {
		f.SetTarget(g.patty.GridX, g.patty.GridY)
	} else {
		targetPlayer := g.turnManager.getPlayerType(config.MergedBurgerType)
		gridX, gridY := -1, -1
		if targetPlayer != nil {
			gridX, gridY = targetPlayer.GridX, targetPlayer.GridY
		}
		f.SetTarget(gridX, gridY)
	}

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
	g.currentEndlessLevel = 0
	g.currentGameState = StateEndless
	g.level = level.NewEndlessLevel(g.currentEndlessLevel)
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
		if !g.needsRestart {
			g.shake = newShake(g.shakeDuration(), shakeDefaultMagnitude)
			g.resetTimer = g.shakeDuration() + 10 // delay a bit
		}
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

// cheesePower unites bun to lettuce and gives it the power of THE killer dash
func (g *Game) cheesePower(bun, cheese *entities.Player) {
	bun.CanDash = true
	bun.Image = merge2Images(bun.Image, cheese.Image, 15.0)
	var newTurnOrder []character
	for _, char := range g.turnManager.turnOrderDisplay {
		if char == cheese {
			continue
		}
		newTurnOrder = append(newTurnOrder, char)
	}
	g.turnManager.turnOrderDisplay = newTurnOrder
}

// lettucePower unites bun to lettuce and gives it the power of crossing walls
func (g *Game) lettucePower(bun, lettuce *entities.Player) {
	bun.CanWalkThroughWalls = true
	bun.Image = merge2Images(bun.Image, lettuce.Image, 0.0)

	var newTurnOrder []character
	for _, char := range g.turnManager.turnOrderDisplay {
		if char == lettuce {
			continue
		}
		newTurnOrder = append(newTurnOrder, char)
	}
	g.turnManager.turnOrderDisplay = newTurnOrder
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

// **** MERGING STUFF ****
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
		DashState:           entities.NewDashState(),
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
	g.needsRestart = false
	g.resetTimer = 0
	g.shake = newShake(shakeDefaultDuration, shakeDefaultMagnitude)
	g.patty = &g.currentLevel().BurgerPatty
	if g.patty != nil {
		g.patty.Reset()
	}

	if g.currentGameState == StateEndless {
		g.currentGameState = StateGameOver
		return
	}

	var characters []character
	for _, v := range g.currentLevel().TurnOrderPattern {
		switch actualActor := v.(type) {
		case entities.Enemier:
			characters = append(characters, actualActor)
		case *entities.Mouse:
			characters = append(characters, actualActor)
		case *entities.Pigeon:
			characters = append(characters, actualActor)
		case *entities.Duck:
			characters = append(characters, actualActor)
		case *entities.Snake:
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

// *** COLLISIONS
func (g *Game) checkCollisionToPlayer(enemy entities.Enemier) {
	for _, v := range g.turnManager.turnOrderDisplay {
		switch player := v.(type) {
		case *entities.Player:
			if enemy.Collision(player) {
				if !g.needsRestart {
					g.shake = newShake(g.shakeDuration(), shakeDefaultMagnitude)
					g.resetTimer = g.shakeDuration() + 10 // delay a bit
				}
				g.audios.eatingSoundPlayer.Rewind()
				g.audios.eatingSoundPlayer.Play()
				g.needsRestart = true
			}
		}
	}
}

func (g *Game) checkCollisionToPatty(enemy entities.Enemier) {
	if g.patty == nil {
		return
	}
	if _, ok := enemy.(*entities.Fly); ok {
		if enemy.Collision(&entities.Player{GridX: g.patty.GridX, GridY: g.patty.GridY}) {
			if !g.needsRestart {
				g.shake = newShake(g.shakeDuration(), shakeDefaultMagnitude)
				g.resetTimer = g.shakeDuration() + 10 // delay a bit
			}
			g.audios.eatingSoundPlayer.Rewind()
			g.audios.eatingSoundPlayer.Play()
			g.needsRestart = true
		}
	}

	for _, v := range g.turnManager.turnOrderDisplay {
		switch player := v.(type) {
		case *entities.Player:
			if enemy.Collision(player) {
				if !g.needsRestart {
					g.shake = newShake(g.shakeDuration(), shakeDefaultMagnitude)
					g.resetTimer = g.shakeDuration() + 10 // delay a bit
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
		// TODO Extend with collision player to player
		switch enemy := v.(type) {
		case entities.Enemier:
			if enemy.Collision(player) {
				isBun := player.PlayerType == config.TopBun || player.PlayerType == config.BottomBun
				if isBun && player.IsDashing() {
					g.turnManager.turnOrderDisplay = append(g.turnManager.turnOrderDisplay[:i], g.turnManager.turnOrderDisplay[i+1:]...)
					g.animationManager.playKillEffect(player.GridX, player.GridY)
				} else if !g.needsRestart {
					g.shake = newShake(g.shakeDuration(), shakeDefaultMagnitude)
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

func (g *Game) shakeDuration() int {
	if g.currentGameState == StateEndless {
		return shakeDefaultDuration
	}
	return shakeDefaultDuration * 2
}
