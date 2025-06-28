package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mikelangelon/unibun/entities"
)

type animationManager struct {
	mergeAnimation *entities.MergeAnimation
	// TODO: properly extract only common parts instead of using the same
	winningAnimation *entities.MergeAnimation
	effects          []entities.Confetti
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
	a.mergeAnimation.ActivateMerge(patty, topBun, bottomBun)
}

func (a *animationManager) playWinningAnimation(x, y int) {
	if a.isWinningPlaying() {
		return
	}
	a.winningAnimation.Activate(120)
	a.winningAnimation.Confetti = entities.CreateConfetti(x, y)
}

func (a *animationManager) playConfettiEffect(x, y int) {
	a.effects = append(a.effects, entities.CreateConfetti(x, y))
}

func (a *animationManager) playKillEffect(x, y int) {
	a.effects = append(a.effects, entities.CreateBlood(x, y))
}

func (a *animationManager) drawMergeAnimation(screen *ebiten.Image, patty *entities.BurgerPatty, topBun, bottomBun *entities.Player) {
	a.mergeAnimation.DrawMergeAnimation(screen, patty, topBun, bottomBun)
}

func (a *animationManager) drawWinningAnimation(screen *ebiten.Image) {
	a.winningAnimation.Draw(screen)
}

func (a *animationManager) drawEffects(screen *ebiten.Image) {
	for _, effect := range a.effects {
		effect.Draw(screen)
	}
}

func (a *animationManager) Update() {
	if a.isMergeAnimationPlaying() {
		a.mergeAnimation.Update()
	}
	if a.isWinningPlaying() {
		a.winningAnimation.Update()
	}

	for _, effect := range a.effects {
		effect.Update()
	}
}
