package level

import "github.com/mikelangelon/unibun/config"

func NewLevel0() *Level {
	w := Cell{Type: CellTypeWall}
	c := Cell{Type: CellTypeFloor}

	lvl := &Level{
		cells: [][]Cell{
			{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
			{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
			{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
			{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
			{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
			{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
			{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
			{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
			{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
			{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
			{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
			{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
			{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
			{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
			{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
		},
	}
	lvl.TurnOrderPattern = []interface{}{
		config.TopBun,
		config.BottomBun,
		config.Cheese,
	}
	return lvl
}
