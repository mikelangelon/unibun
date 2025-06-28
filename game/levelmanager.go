package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/mikelangelon/unibun/config"
	"github.com/mikelangelon/unibun/level"
)

const numRegularLevels = 15

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
			1:  level.NewIntro,
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
			16: level.Unite,
		},
	}
}

func (l *levelManager) Update() error {
	cols := 5
	// The secret game is displayed different than the rest.
	if l.selectedLevelBox < numRegularLevels {
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) && (l.selectedLevelBox+1)%cols != 0 && l.selectedLevelBox < numRegularLevels-1 {
			l.selectedLevelBox++
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) && l.selectedLevelBox%cols != 0 {
			l.selectedLevelBox--
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) && l.selectedLevelBox >= cols {
			l.selectedLevelBox -= cols
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
			if l.selectedLevelBox < numRegularLevels-cols {
				l.selectedLevelBox += cols
			} else if l.AllLevelsCompleted() {
				l.selectedLevelBox = numRegularLevels
			}
		}
	} else { // Secret level is selected
		if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
			l.selectedLevelBox = 12 // Go back to middle of last row
		}
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
	if l.currentLevelIndex < len(l.levelConstructors) {
		l.selectedLevelBox = l.currentLevelIndex
	}
}

func (l *levelManager) AllLevelsCompleted() bool {
	return len(l.completedLevels) >= numRegularLevels
}

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

// draws 15 boxes, in a 3x5 grid.
func (l *levelManager) draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 20, G: 20, B: 20, A: 255})
	ebitenutil.DebugPrintAt(screen, "Select a Level", config.WindowWidth/2-70, 40)

	for i := 0; i < numRegularLevels; i++ {
		col := i % cols
		row := i / cols
		boxX := startX + col*(boxSize+padding)
		boxY := startY + row*(boxSize+padding)

		boxColor := color.RGBA{40, 40, 40, 255}
		levelNum := i + 1
		if l.completedLevels[levelNum] {
			boxColor = color.RGBA{20, 60, 20, 255}
		}
		if i == l.selectedLevelBox {
			boxColor = color.RGBA{90, 90, 90, 255}
		}
		ebitenutil.DrawRect(screen, float64(boxX), float64(boxY), float64(boxSize), float64(boxSize), color.RGBA{60, 60, 60, 255})
		ebitenutil.DrawRect(screen, float64(boxX+5), float64(boxY+5), float64(boxSize)-10, float64(boxSize)-10, boxColor)
		levelNumStr := fmt.Sprintf("%d", levelNum)
		textX := boxX + (boxSize-len(levelNumStr)*6)/2
		textY := boxY + (boxSize-16)/2
		ebitenutil.DebugPrintAt(screen, levelNumStr, textX, textY)
	}

	if l.AllLevelsCompleted() {
		levelNum := 16
		col := 2
		row := 3
		boxX := startX + col*(boxSize+padding)
		boxY := startY + row*(boxSize+padding)

		boxColor := color.RGBA{100, 20, 100, 255} // Purple for secret
		if l.completedLevels[levelNum] {
			boxColor = color.RGBA{20, 60, 20, 255} // Green for completed
		}
		if l.selectedLevelBox == numRegularLevels {
			boxColor = color.RGBA{150, 80, 150, 255} // Brighter purple for selected
		}
		ebitenutil.DrawRect(screen, float64(boxX), float64(boxY), float64(boxSize), float64(boxSize), color.RGBA{60, 60, 60, 255})
		ebitenutil.DrawRect(screen, float64(boxX+5), float64(boxY+5), float64(boxSize)-10, float64(boxSize)-10, boxColor)
		ebitenutil.DebugPrintAt(screen, "??", boxX+(boxSize-12)/2, boxY+(boxSize-16)/2)
	}
}
