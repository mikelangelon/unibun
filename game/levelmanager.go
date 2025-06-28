package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/mikelangelon/unibun/level"
)

type levelManager struct {
	levelConstructors map[int]func() *level.Level
	currentLevelIndex int
	selectedLevelBox  int
	completedLevels   map[int]bool
	startLevelFunc    func(levelNum int)
}

func newLevelManager(startLevelFunc func(levelNum int)) *levelManager {
	return &levelManager{
		startLevelFunc:   startLevelFunc,
		completedLevels:  make(map[int]bool),
		selectedLevelBox: 0,
		levelConstructors: map[int]func() *level.Level{
			1:  level.NewEmptyLevel, //level.NewIntro,
			2:  level.LettucePresentation,
			3:  level.CheesePresentation,
			4:  level.FirstRealLevel,
			5:  level.AvoidTheLettuce,
			6:  level.NewFlies,
			7:  level.FourSnakes,
			8:  level.PushThePatty,
			9:  level.PuzzleBuns,
			10: level.ManyObstacles,
			11: level.FourSnakesReturn,
			12: level.NewLevelLettuceMaze,
			13: level.AnotherLettuce,
			14: level.NewLevelLettuceMazeHard,
			15: level.SnakesLevel,
		},
	}
}

func (l *levelManager) Update() error {
	cols := 5
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) && (l.selectedLevelBox+1)%cols != 0 && l.selectedLevelBox < 14 {
		l.selectedLevelBox++
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) && l.selectedLevelBox%cols != 0 {
		l.selectedLevelBox--
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) && l.selectedLevelBox >= cols {
		l.selectedLevelBox -= cols
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) && l.selectedLevelBox < 10 {
		l.selectedLevelBox += cols
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		l.currentLevelIndex = l.selectedLevelBox + 1
		l.startLevelFunc(l.currentLevelIndex)
	}
	return nil
}

func (l *levelManager) passNextLevel() {
	if l.currentLevelIndex > 0 {
		l.completedLevels[l.currentLevelIndex] = true
	}
	if l.currentLevelIndex < len(l.levelConstructors)-1 {
		l.selectedLevelBox = l.currentLevelIndex
	}
}
