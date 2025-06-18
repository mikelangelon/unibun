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
		Winning: []Position{
			{X: 6, Y: 6},
		},
	}
	lvl.TurnOrderPattern = []interface{}{
		entities.NewPlayer(3, 1, config.TopBun),
		entities.NewPlayer(3, 5, config.BottomBun),
		entities.NewEnemy(10, 10),
	}
	return lvl
}

func NewLevel1() *Level {
	w := Cell{Type: CellTypeWall}
	c := Cell{Type: CellTypeFloor}
	lvl := &Level{
		cells: [][]Cell{
			{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
			{w, c, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w, w, c, c, c, c, c, w},
			{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w, c, c, c, c, c, c, w},
			{w, c, w, c, c, c, c, c, c, c, w, c, c, c, c, c, c, w, c, w, w, w, w, c, w},
			{w, c, w, c, c, c, c, c, c, c, w, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
			{w, c, w, c, c, c, c, c, c, c, w, c, c, c, c, w, w, w, w, w, w, w, w, c, w},
			{w, c, w, c, c, c, c, c, c, c, w, c, c, c, c, c, w, c, c, c, c, c, c, c, w},
			{w, c, c, c, c, c, c, c, c, c, w, c, c, c, c, c, w, c, w, w, c, w, w, w, w},
			{w, w, w, c, c, c, c, c, c, c, w, c, c, c, c, c, w, c, w, c, c, w, w, c, w},
			{w, c, c, c, c, c, c, c, c, c, w, c, c, c, c, c, w, c, w, c, c, w, w, c, w},
			{w, c, w, c, c, c, c, c, c, c, w, c, c, c, c, w, w, c, w, c, c, c, w, c, w},
			{w, c, w, c, c, c, c, c, c, c, w, c, c, c, c, c, c, c, c, c, c, c, w, c, w},
			{w, c, w, w, w, w, w, w, w, c, w, w, w, w, c, w, c, w, w, w, w, c, w, c, w},
			{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
			{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
		},
		BurgerPatty: entities.NewBurgerPatty(8, 8),
		Winning: []Position{
			{X: 11, Y: 10},
		},
	}

	lvl.TurnOrderPattern = []interface{}{
		entities.NewPlayer(3, 1, config.TopBun),
		entities.NewPlayer(10, 5, config.BottomBun),
		entities.NewEnemy(10, 10),
	}
	return lvl
}

func NewLevel2() *Level {
	w := Cell{Type: CellTypeWall}
	c := Cell{Type: CellTypeFloor}
	lvl := &Level{
		cells: [][]Cell{
			{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
			{w, c, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w, w, c, c, c, c, c, w},
			{w, c, c, c, w, w, w, w, c, w, w, w, c, c, c, c, c, w, c, c, c, c, c, c, w},
			{w, c, w, w, w, c, c, c, c, c, c, c, c, c, c, c, c, w, c, w, w, w, w, c, w},
			{w, c, w, c, c, c, c, w, c, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
			{w, c, w, c, c, c, c, c, c, w, c, w, w, c, c, w, w, w, w, w, w, w, w, c, w},
			{w, c, w, c, c, c, c, c, c, w, c, w, w, c, c, c, w, c, c, c, c, c, c, c, w},
			{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w, c, w, w, c, w, w, w, w},
			{w, w, w, c, c, c, c, c, c, w, c, w, w, c, c, c, w, c, w, c, c, w, w, c, w},
			{w, c, c, c, c, c, c, c, c, w, c, w, w, c, c, c, w, c, w, c, c, w, w, c, w},
			{w, c, w, c, c, c, c, c, c, w, c, w, w, c, c, w, w, c, w, c, c, c, w, c, w},
			{w, c, w, c, c, c, c, c, c, c, c, c, c, w, c, c, c, c, c, c, c, c, w, c, w},
			{w, c, w, w, w, w, w, w, w, w, c, w, w, w, c, w, c, w, w, w, w, c, w, c, w},
			{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
			{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
		},
		BurgerPatty: entities.NewBurgerPatty(21, 12),
		Winning: []Position{
			{X: 20, Y: 6},
		},
	}

	lvl.TurnOrderPattern = []interface{}{
		entities.NewPlayer(10, 10, config.TopBun),
		entities.NewPlayer(1, 13, config.BottomBun),
		entities.NewEnemy(1, 1),
	}
	return lvl
}
