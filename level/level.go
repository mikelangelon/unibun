package level

import (
	"math"
	"math/rand/v2"

	"github.com/mikelangelon/unibun/entities"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mikelangelon/unibun/config"
)

type Level struct {
	cells            [][]Cell
	TurnOrderPattern []interface{}
	BurgerPatty      entities.BurgerPatty
	Winning          []Position
	WinningImg       *ebiten.Image
	FloorTileImg     []*ebiten.Image
	IntroText        string

	FloorMap    *ebiten.Image
	pulseOffset float64
}

func (l *Level) getRandomTile() *ebiten.Image {
	defaultTileProb := 20
	r := rand.IntN(defaultTileProb + len(l.FloorTileImg))
	if r < defaultTileProb {
		return l.FloorTileImg[0]
	}
	return l.FloorTileImg[r-defaultTileProb]
}

func (l *Level) Draw(screen *ebiten.Image, isBurgerMerged bool) {
	if l.FloorMap == nil {
		l.FloorMap = ebiten.NewImage(l.ScreenWidth(), l.ScreenHeight())
		for ri, row := range l.cells {
			for ci, cell := range row {
				cellX := float64(ci * config.TileSize)
				cellY := float64(ri * config.TileSize)

				if cell.Type == CellTypeFloor && len(l.FloorTileImg) > 0 {
					tile := l.getRandomTile()
					op := &ebiten.DrawImageOptions{}
					if (ri+ci)%2 == 0 {
						op.ColorScale.Scale(0.8, 0.8, 0.8, 1.0)
					} else {
						op.ColorScale.Scale(0.5, 0.5, 0.5, 1.0)
					}
					op.GeoM.Translate(cellX, cellY)
					l.FloorMap.DrawImage(tile, op)
					continue
				}
			}
		}
	}

	screen.DrawImage(l.FloorMap, &ebiten.DrawImageOptions{})
	for _, v := range l.Winning {
		winRectX := float64(v.X * config.TileSize)
		winRectY := float64(v.Y * config.TileSize)
		op := &ebiten.DrawImageOptions{}

		if isBurgerMerged {
			scale := 1.0 + 0.1*math.Sin(l.pulseOffset/20.0)
			w, h := l.WinningImg.Bounds().Dx(), l.WinningImg.Bounds().Dy()
			centerX, centerY := float64(w)/2, float64(h)/2

			op.GeoM.Translate(-centerX, -centerY)
			op.GeoM.Scale(scale, scale)
			op.GeoM.Translate(centerX, centerY)
		}

		op.GeoM.Translate(winRectX, winRectY)
		screen.DrawImage(l.WinningImg, op)
	}
	if isBurgerMerged {
		l.pulseOffset++
	}
}

func (l *Level) OutOfBounds(gridX, gridY int) bool {
	return gridX < 0 || gridX >= len(l.cells[0]) || gridY < 0 || gridY >= len(l.cells)
}

func (l *Level) GetCell(gridX, gridY int) *Cell {
	if l.OutOfBounds(gridX, gridY) {
		return nil
	}
	return &l.cells[gridY][gridX]
}

func (l *Level) gridCols() int {
	return len(l.cells[0])
}

func (l *Level) gridRows() int {
	return len(l.cells)
}

func (l *Level) ScreenHeight() int {
	return l.gridRows() * config.TileSize
}

func (l *Level) ScreenWidth() int {
	return l.gridCols() * config.TileSize
}

func (l *Level) IsWalkable(gridX, gridY int) bool {
	if l.OutOfBounds(gridX, gridY) {
		return false
	}
	cell := l.GetCell(gridX, gridY)
	if cell == nil {
		return false
	}
	switch cell.Type {
	case CellTypeFloor, CellTypeEmpty:
		return true
	case CellTypeWall:
		return false
	default:
		return false
	}

}
