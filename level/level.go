package level

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/mikelangelon/unibun/config"
)

type Level struct {
	cells [][]Cell
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
}
