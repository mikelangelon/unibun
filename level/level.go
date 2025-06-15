package level

import (
	"github.com/mikelangelon/unibun/entities"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/mikelangelon/unibun/config"
)

type Level struct {
	cells            [][]Cell
	TurnOrderPattern []interface{}
	BurgerPatty      *entities.BurgerPatty
	Winning          []Position
}

func (l *Level) Draw(screen *ebiten.Image) {
	for ri, row := range l.cells {
		for ci, cell := range row {
			cellX := float64(ci * config.TileSize)
			cellY := float64(ri * config.TileSize)
			var cellColor color.Color
			switch cell.Type {
			case CellTypeFloor:
				cellColor = color.RGBA{R: 0x60, G: 0x60, B: 0x60, A: 0xff}
			case CellTypeWall:
				cellColor = color.RGBA{R: 0x30, G: 0x30, B: 0x80, A: 0xff}
			case CellTypeEmpty:
				cellColor = color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff}
			default:
				cellColor = color.White
			}
			ebitenutil.DrawRect(screen, cellX, cellY, float64(config.TileSize), float64(config.TileSize), cellColor)
		}
	}

	gridColor := color.RGBA{R: 0x40, G: 0x40, B: 0x40, A: 0xff}
	strokeWidth := float32(1)
	for i := 0; i <= l.gridCols(); i++ {
		x := float32(i * config.TileSize)
		vector.StrokeLine(screen, x, 0, x, float32(l.ScreenHeight()), strokeWidth, gridColor, false)
	}
	for i := 0; i <= l.gridRows(); i++ {
		y := float32(i * config.TileSize)
		vector.StrokeLine(screen, 0, y, float32(l.ScreenWidth()), y, strokeWidth, gridColor, false)
	}

	for _, v := range l.Winning {
		winRectX := float64(v.X * config.TileSize)
		winRectY := float64(v.Y * config.TileSize)
		winTileMarkerColor := color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xcc}
		ebitenutil.DrawRect(screen, winRectX, winRectY, float64(config.TileSize), float64(config.TileSize), winTileMarkerColor)
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
