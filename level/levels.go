package level

import (
	"github.com/mikelangelon/unibun/assets"
	"github.com/mikelangelon/unibun/common"
	"github.com/mikelangelon/unibun/config"
	"github.com/mikelangelon/unibun/entities"
)

func NewLevel0() *Level {
	w := Cell{Type: CellTypeWall}
	c := Cell{Type: CellTypeFloor}
	lvl := newLevel()
	lvl.BurgerPatty = entities.NewBurgerPatty(3, 3)
	lvl.Winning = []Position{
		{X: 6, Y: 6},
	}
	lvl.cells = [][]Cell{
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
	}
	// examplePath := []image.Point{
	// 	{X: 10, Y: 4}, {X: 11, Y: 4}, {X: 12, Y: 4}, {X: 13, Y: 4}, {X: 14, Y: 4},
	// }
	lvl.TurnOrderPattern = []interface{}{
		entities.NewPlayer(3, 1, config.TopBun),
		entities.NewPlayer(3, 5, config.BottomBun),
		//entities.NewPlayer(2, 1, config.Cheese),
		//entities.NewPlayer(2, 2, config.Lettuce),
		entities.NewDashingFollowerEnemy(1, 13, config.TopBun, 3), // New dashing follower enemy targeting Lettuce, dashes every 5 turns
		// entities.NewEnemy(10, 10),
		//entities.NewFollowerEnemy(8, 9, config.Cheese), // New follower enemy targeting Cheese
		//entities.NewPathEnemy(examplePath[0].X, examplePath[0].Y, examplePath, assets.Pidgeon),
	}
	lvl.WinningImg = common.GetImage(assets.Client)
	return lvl
}

func NewLevel1() *Level {
	w := Cell{Type: CellTypeWall}
	c := Cell{Type: CellTypeFloor}

	lvl := newLevel()
	lvl.BurgerPatty = entities.NewBurgerPatty(8, 8)
	lvl.Winning = []Position{
		{X: 11, Y: 10},
	}
	lvl.cells = [][]Cell{
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
	lvl := newLevel()
	lvl.BurgerPatty = entities.NewBurgerPatty(21, 12)
	lvl.Winning = []Position{
		{X: 20, Y: 6},
	}
	lvl.cells = [][]Cell{
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
	}
	lvl.TurnOrderPattern = []interface{}{
		entities.NewPlayer(10, 10, config.TopBun),
		entities.NewPlayer(1, 13, config.BottomBun),
		entities.NewPlayer(10, 12, config.Lettuce),
		entities.NewEnemy(1, 1),
	}
	return lvl
}

func NewLevel3() *Level {
	w := Cell{Type: CellTypeWall}
	c := Cell{Type: CellTypeFloor}
	lvl := newLevel()
	lvl.BurgerPatty = entities.NewBurgerPatty(12, 7)
	lvl.Winning = []Position{
		{X: 23, Y: 1},
	}
	lvl.cells = [][]Cell{
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w}, // 0
		{w, c, c, c, c, c, w, c, w, w, w, w, w, w, w, w, w, c, w, c, w, c, c, c, w}, // 1
		{w, c, w, c, c, c, w, c, w, c, c, w, c, c, w, c, w, c, w, c, c, c, w, c, w}, // 2
		{w, c, w, c, w, c, w, c, w, c, c, w, c, c, w, c, w, c, w, c, w, c, w, c, w}, // 3
		{w, c, w, c, c, c, w, c, w, c, c, w, c, c, w, c, w, c, c, c, w, c, c, c, w}, // 4
		{w, c, c, c, c, c, w, c, w, w, w, w, w, w, w, c, w, c, w, c, w, c, w, w, w}, // 5
		{w, c, w, c, w, c, w, c, c, c, c, c, c, c, c, c, w, c, w, c, w, c, c, c, w}, // 6
		{w, c, w, c, c, c, w, c, c, c, w, c, c, c, c, c, w, c, c, c, w, c, c, c, w}, // 7 BurgerPatty(12,7)
		{w, c, w, c, w, c, c, c, c, c, c, c, c, c, c, c, c, c, w, c, w, c, w, c, w}, // 8
		{w, c, c, c, c, c, w, c, w, w, w, w, w, w, w, c, w, c, w, c, w, c, c, c, w}, // 9
		{w, c, w, c, c, c, w, c, w, c, c, w, c, c, w, c, w, c, c, c, w, c, w, c, w}, // 10
		{w, c, w, c, w, c, w, c, w, c, c, w, c, c, w, c, w, c, w, c, w, c, c, c, w}, // 11
		{w, c, w, c, c, c, w, c, w, c, c, w, c, c, w, c, w, c, w, c, c, c, w, c, w}, // 12
		{w, c, c, c, c, c, w, c, w, w, w, w, w, w, w, w, w, c, w, c, w, c, c, c, w}, // 13
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w}, // 14
	}

	lvl.TurnOrderPattern = []interface{}{
		entities.NewPlayer(1, 1, config.TopBun),     // Top-left start
		entities.NewPlayer(1, 13, config.BottomBun), // Bottom-left start
		entities.NewEnemy(6, 2),                     // Enemy in the upper section
		entities.NewEnemy(18, 12),                   // Enemy in the lower-right section
		entities.NewEnemy(12, 4),                    // Enemy guarding near a path to burger
	}
	return lvl
}

func newLevel() *Level {
	return &Level{
		TurnOrderPattern: nil,
		BurgerPatty:      entities.NewBurgerPatty(12, 7),
		Winning: []Position{
			{X: 1, Y: 1},
		},
		WinningImg: common.GetImage(assets.Client),
	}
}

func NewLevel4() *Level {
	w := Cell{Type: CellTypeWall}
	c := Cell{Type: CellTypeFloor}
	lvl := newLevel()
	lvl.BurgerPatty = entities.NewBurgerPatty(12, 7)
	lvl.Winning = []Position{
		{X: 1, Y: 1},
	}
	lvl.cells = [][]Cell{
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
		{w, c, c, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, w, w, c, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, c, w},
		{w, c, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w, c, w},
		{w, c, w, c, w, w, w, w, w, w, w, c, w, w, w, w, w, w, w, c, w, c, w, c, w},
		{w, c, w, c, w, c, c, c, c, c, c, c, c, c, c, c, c, c, w, c, w, c, w, c, w},
		{w, c, w, c, w, c, w, w, w, w, w, w, w, w, w, c, w, c, w, c, w, c, w, c, w},
		{w, c, c, c, w, c, c, c, c, c, c, c, c, c, c, c, w, c, c, c, w, c, c, c, w},
		{w, c, w, c, w, c, w, w, w, w, w, w, w, w, w, c, w, c, w, c, w, c, w, c, w},
		{w, c, w, c, w, c, c, c, c, c, c, c, c, c, c, c, c, c, w, c, w, c, w, c, w},
		{w, c, w, c, w, w, w, w, w, w, w, c, w, w, w, w, w, w, w, c, w, c, w, c, w},
		{w, c, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w, c, w},
		{w, c, w, w, c, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, c, w},
		{w, c, c, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
	}

	lvl.TurnOrderPattern = []interface{}{
		entities.NewPlayer(23, 1, config.TopBun),     // Top-right start
		entities.NewPlayer(23, 13, config.BottomBun), // Bottom-right start
		entities.NewEnemy(5, 5),                      // Guarding a path in the top-left area
		entities.NewEnemy(10, 7),                     // Near the burger patty
		entities.NewEnemy(17, 9),                     // Patrolling a corridor in the mid-right area
	}
	return lvl
}
