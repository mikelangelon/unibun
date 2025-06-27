package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mikelangelon/unibun/entities"
)

type animationManager struct {
	mergeAnimation *entities.MergeAnimation
	// TODO: properly extract only common parts instead of using the same
	winningAnimation *entities.MergeAnimation
}

func newAnimationManager() *animationManager {
	return &animationManager{
		mergeAnimation:   &entities.MergeAnimation{},
		winningAnimation: &entities.MergeAnimation{},
	}
}

func (a *animationManager) blockingAnimation() bool {
	return a.mergeAnimation.IsActive
}

func (a *animationManager) isMergeAnimationPlaying() bool {
	return a.mergeAnimation.IsActive
}

func (a *animationManager) isWinningPlaying() bool {
	return a.winningAnimation.IsActive
}

func (a *animationManager) playMergeAnimation(patty *entities.BurgerPatty, topBun, bottomBun *entities.Player) {
	a.mergeAnimation.Activate(patty, topBun, bottomBun)
}

func (a *animationManager) playWinningAnimation(x, y int) {
	if a.isWinningPlaying() {
		return
	}
	a.winningAnimation.IsActive = true
	a.winningAnimation.Duration = 120
	a.winningAnimation.Timer = a.winningAnimation.Duration
	a.winningAnimation.Confetti = entities.CreateConfetti(x, y)
}

func (a *animationManager) drawMergeAnimation(screen *ebiten.Image, patty *entities.BurgerPatty, topBun, bottomBun *entities.Player) {
	a.mergeAnimation.DrawMergeAnimation(screen, patty, topBun, bottomBun)
}

func (a *animationManager) drawWinningAnimation(screen *ebiten.Image) {
	a.winningAnimation.Draw(screen)
}
func (a *animationManager) Update() {
	if a.isMergeAnimationPlaying() {
		a.mergeAnimation.Update()
		if a.mergeAnimation.Timer <= 0 {
			a.mergeAnimation.IsActive = false
		}
	}
	if a.isWinningPlaying() {
		a.winningAnimation.Update()
		if a.winningAnimation.Timer <= 0 {
			a.winningAnimation.IsActive = false
		}
	}
}
