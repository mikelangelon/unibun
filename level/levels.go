package level

import (
	"github.com/mikelangelon/unibun/config"
	"github.com/mikelangelon/unibun/entities"
)

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
		BurgerPatty: entities.NewBurgerPatty(3, 3),
	}
	lvl.TurnOrderPattern = []interface{}{
		entities.NewPlayer(3, 1, config.TopBun),
		entities.NewPlayer(3, 5, config.BottomBun),
		entities.NewPlayer(4, 4, config.Cheese),
	}
	return lvl
}
