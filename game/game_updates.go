package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/mikelangelon/unibun/config"
	"github.com/mikelangelon/unibun/entities"
	"os"
)

func (g *Game) Update() error {
	if g.currentGameState != g.previousGameState {
		g.handleGameStateChange()
		g.previousGameState = g.currentGameState
	}

	switch g.currentGameState {
	case StateMenu:
		return g.menu.updateMenu(g)
	case StatePlaying, StateEndless:
		return g.updatePlaying()
	case StateLevelSelect:
		return g.levelManager.Update()
	case StateIntro:
		return g.updateIntro()
	case StateTutorial:
		return g.updateTutorial()
	case StatePaused:
		return g.updatePaused()
	case StateGameComplete:
		return g.updateGameComplete()
	case StateGameOver:
		return g.updateGameOver()
	case StateExiting:
		os.Exit(0)
	}
	return nil
}

func (g *Game) updateTutorial() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.currentGameState = StateMenu
	}
	return nil
}

func (g *Game) updateGameComplete() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.currentGameState = StateMenu
	}
	return nil
}

func (g *Game) updateGameOver() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.currentGameState = StateMenu
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.startEndlessGame()
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
		g.pauseMenuOptions[g.selectedPauseMenuOption].action(g)
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
			g.bunCollidesBun(actor)
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
				pushedPatty, stopPath := false, false
				moved, pushedPatty, stopPath = g.checkBunPatty(actor, oldX, oldY)
				if stopPath {
					break
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
			g.levelManager.passNextLevel() // Mark level as complete
			// If we just beat the secret level (16), show the final win screen.
			// Otherwise, go back to the select screen to reveal the secret level.
			if g.levelManager.currentLevelIndex == 16 {
				g.currentGameState = StateGameComplete
			} else {
				g.currentGameState = StateLevelSelect
			}
		}
	}
}
