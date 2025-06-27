package level

import (
	"github.com/mikelangelon/unibun/assets"
	"github.com/mikelangelon/unibun/common"
	"github.com/mikelangelon/unibun/config"
	"github.com/mikelangelon/unibun/entities"
	"image"
)

func NewLevel0() *Level {
	w := Cell{Type: CellTypeWall}
	c := Cell{Type: CellTypeFloor}
	lvl := newLevel()
	lvl.BurgerPatty = entities.NewBurgerPatty(13, 4)
	lvl.Winning = []Position{
		{X: 13, Y: 9},
	}
	lvl.cells = [][]Cell{
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, w, w, c, c, c, c, c, c, c, c, w, w, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, w, c, c, w, w, w, w, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, w, w, w, w, c, c, w, w, w, w, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
	}
	lvl.IntroText = "Your goal is to unite the buns with the patty,\nand deliver it to the client.\n\nPress Enter to continue"
	examplePath := []image.Point{
		{X: 10, Y: 13}, {X: 11, Y: 13}, {X: 12, Y: 13}, {X: 13, Y: 13}, {X: 14, Y: 13},
	}
	lvl.TurnOrderPattern = []interface{}{
		entities.NewPlayer(3, 1, config.TopBun),
		entities.NewPlayer(21, 1, config.BottomBun),
		//entities.NewPlayer(2, 1, config.Cheese),
		//entities.NewPlayer(2, 2, config.Lettuce),
		entities.NewDashingFollowerEnemy(1, 13, config.TopBun, 3), // New dashing follower enemy targeting Lettuce, dashes every 5 turns
		// entities.NewEnemy(10, 10),
		//entities.NewFollowerEnemy(8, 9, config.Cheese), // New follower enemy targeting Cheese
		entities.NewPathEnemy(examplePath[0].X, examplePath[0].Y, examplePath, assets.Pidgeon),
	}
	lvl.WinningImg = common.GetImage(assets.Client)
	return lvl
}

func UseOtherObject() *Level {
	w := Cell{Type: CellTypeWall}
	c := Cell{Type: CellTypeFloor}
	lvl := newLevel()
	lvl.BurgerPatty = entities.NewBurgerPatty(4, 3)
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
		{w, c, c, c, c, c, c, c, c, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
	}
	path := []image.Point{
		{X: 9, Y: 13}, {X: 10, Y: 13}, {X: 11, Y: 13}, {X: 12, Y: 13}, {X: 13, Y: 13}, {X: 14, Y: 13},
	}
	path2 := []image.Point{
		{X: 10, Y: 11}, {X: 11, Y: 11}, {X: 12, Y: 11}, {X: 13, Y: 11}, {X: 14, Y: 11},
	}
	lvl.TurnOrderPattern = []interface{}{
		entities.NewPlayer(5, 1, config.TopBun),
		entities.NewPlayer(5, 6, config.BottomBun),
		entities.NewPlayer(7, 5, config.Cheese),
		entities.NewPathEnemy(path[0].X, path[0].Y, path, assets.Pidgeon),
		entities.NewPathEnemy(path2[0].X, path2[0].Y, path2, assets.Pidgeon),
	}
	lvl.WinningImg = common.GetImage(assets.Client)
	return lvl
}

// Bring cheese to TopBun, avoid first snake.
func ManyObstacles() *Level {
	w := Cell{Type: CellTypeWall}
	c := Cell{Type: CellTypeFloor}
	lvl := newLevel()
	lvl.BurgerPatty = entities.NewBurgerPatty(9, 12)
	lvl.Winning = []Position{
		{X: 10, Y: 8},
	}
	lvl.cells = [][]Cell{
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, w, w, w, w, w, w, w, c, w, w, w, w, w, c, w, w, c, w, w, w, w, w, w, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, w, w, c, w, w, w, w, w, w, w, w, w, w, w, w, w, c, w, w, w, w, w, w, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, w, w, w, w, w, w, w, c, w, w, w, w, w, c, w, w, c, w, w, w, w, w, w, w},
		{w, c, c, c, c, c, c, c, c, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, w, w, c, w, w, w, w, w, w, w, w, w, w, w, w, w, c, w, w, w, w, w, w, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
	}
	path := []image.Point{
		{X: 9, Y: 13}, {X: 10, Y: 13}, {X: 11, Y: 13}, {X: 12, Y: 13}, {X: 13, Y: 13}, {X: 14, Y: 13},
	}
	path2 := []image.Point{
		{X: 10, Y: 11}, {X: 11, Y: 11}, {X: 12, Y: 11}, {X: 13, Y: 11}, {X: 14, Y: 11},
	}
	lvl.TurnOrderPattern = []interface{}{
		entities.NewPlayer(5, 1, config.TopBun),
		entities.NewPlayer(5, 6, config.BottomBun),
		entities.NewPlayer(7, 5, config.Cheese),
		entities.NewPathEnemy(path[0].X, path[0].Y, path, assets.Pidgeon),
		entities.NewPathEnemy(path2[0].X, path2[0].Y, path2, assets.Pidgeon),
		entities.NewDashingFollowerEnemy(4, 1, config.TopBun, 3),
		entities.NewDashingFollowerEnemy(8, 8, config.TopBun, 3),
	}
	lvl.WinningImg = common.GetImage(assets.Client)
	return lvl
}

func NewLevel1b() *Level {
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
		{w, c, c, w, c, c, c, c, c, c, c, c, c, c, c, c, c, w, c, c, c, c, c, c, w},
		{w, c, c, w, c, c, c, c, c, c, w, c, c, c, c, c, c, w, c, w, w, w, w, c, w},
		{w, c, c, c, c, c, c, c, c, c, w, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, w, c, c, c, c, w, w, w, w, w, w, w, w, c, w},
		{w, c, c, w, c, c, c, c, c, c, w, c, c, c, c, c, w, c, c, c, c, c, c, c, w},
		{w, c, c, w, c, c, c, c, c, c, w, c, c, c, c, c, w, c, w, w, c, w, w, w, w},
		{w, w, c, c, c, c, c, c, c, c, w, c, c, c, c, c, w, c, w, c, c, w, w, c, w},
		{w, c, c, c, c, c, c, c, c, c, w, c, c, c, c, c, w, c, w, c, c, w, w, c, w},
		{w, c, c, w, c, c, c, c, c, c, w, c, c, c, c, w, w, c, w, c, c, c, w, c, w},
		{w, c, c, w, c, c, c, c, c, c, w, c, c, c, c, c, c, c, c, c, c, c, w, c, w},
		{w, c, c, w, w, w, w, w, w, c, w, w, w, w, c, w, c, w, w, w, w, c, w, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
	}
	path1 := []image.Point{
		{X: 5, Y: 3}, {X: 5, Y: 4}, {X: 5, Y: 5}, {X: 5, Y: 6}, {X: 5, Y: 7},
	}
	path2 := []image.Point{
		{X: 11, Y: 7}, {X: 12, Y: 7}, {X: 13, Y: 7}, {X: 14, Y: 7},
	}
	lvl.TurnOrderPattern = []interface{}{
		entities.NewPlayer(3, 1, config.TopBun),
		entities.NewPlayer(10, 5, config.BottomBun),
		entities.NewEnemy(10, 10),
		entities.NewPathEnemy(path1[0].X, path1[0].Y, path1, assets.Pidgeon),
		entities.NewPathEnemy(path2[0].X, path2[0].Y, path2, assets.Pidgeon),
	}
	return lvl
}

func NewLevel1c() *Level {
	w := Cell{Type: CellTypeWall}
	c := Cell{Type: CellTypeFloor}

	lvl := newLevel()
	lvl.BurgerPatty = entities.NewBurgerPatty(18, 9)
	lvl.Winning = []Position{
		{X: 11, Y: 10},
	}
	lvl.cells = [][]Cell{
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, c, w, w, w, w, w, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w, w, w, w, w, w, w, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w, w, w, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w, w, w, c, c, c, c, w},
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
	}
	path1 := []image.Point{
		{X: 5, Y: 3}, {X: 5, Y: 4}, {X: 5, Y: 5}, {X: 5, Y: 6}, {X: 5, Y: 7},
	}
	path2 := []image.Point{
		{X: 11, Y: 7}, {X: 12, Y: 7}, {X: 13, Y: 7}, {X: 14, Y: 7},
	}
	lvl.TurnOrderPattern = []interface{}{
		entities.NewPlayer(3, 1, config.TopBun),
		entities.NewPlayer(10, 6, config.BottomBun),
		entities.NewPlayer(15, 8, config.Lettuce),
		entities.NewEnemy(10, 10),
		entities.NewPathEnemy(path1[0].X, path1[0].Y, path1, assets.Pidgeon),
		entities.NewPathEnemy(path2[0].X, path2[0].Y, path2, assets.Pidgeon),
	}
	return lvl
}

func NewLevelLettucePresentation() *Level {
	w := Cell{Type: CellTypeWall}
	c := Cell{Type: CellTypeFloor}

	lvl := newLevel()
	lvl.BurgerPatty = entities.NewBurgerPatty(11, 6)
	lvl.Winning = []Position{
		{X: 11, Y: 10},
	}
	lvl.cells = [][]Cell{
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w, w, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w, w, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, w, w, w, w, w, c, c, c, c, c, c, c, c, w, w, w},
		{w, c, c, c, c, c, c, c, c, w, w, w, w, w, c, c, c, c, c, c, c, c, w, w, w},
		{w, c, c, c, c, c, c, c, c, w, w, c, w, w, c, c, c, c, c, c, c, c, w, w, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
	}
	path1 := []image.Point{
		{X: 5, Y: 3}, {X: 5, Y: 4}, {X: 5, Y: 5}, {X: 5, Y: 6}, {X: 5, Y: 7},
	}
	path2 := []image.Point{
		{X: 11, Y: 7}, {X: 12, Y: 7}, {X: 13, Y: 7}, {X: 14, Y: 7},
	}
	lvl.TurnOrderPattern = []interface{}{
		entities.NewPlayer(11, 1, config.TopBun),
		entities.NewPlayer(5, 8, config.BottomBun),
		entities.NewPlayer(11, 3, config.Lettuce),
		entities.NewEnemy(10, 10),
		entities.NewPathEnemy(path1[0].X, path1[0].Y, path1, assets.Pidgeon),
		entities.NewPathEnemy(path2[0].X, path2[0].Y, path2, assets.Pidgeon),
	}
	return lvl
}

// Take the cheese & lettuce and run against all!
func NewLevelLettuceMaze() *Level {
	w := Cell{Type: CellTypeWall}
	c := Cell{Type: CellTypeFloor}

	lvl := newLevel()
	lvl.BurgerPatty = entities.NewBurgerPatty(13, 9)
	lvl.Winning = []Position{
		{X: 10, Y: 11},
	}
	lvl.cells = [][]Cell{
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
		{w, c, c, c, c, c, c, c, c, w, w, w, w, w, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, w, w, w, w, w, w, c, c, c, c, c, c, c, w, w, w, w, w, w, w, w, c, w},
		{w, c, c, w, w, w, w, w, w, w, w, w, w, w, w, w, w, c, c, c, c, w, c, c, w},
		{w, w, c, w, w, w, w, w, w, w, w, w, w, w, w, w, w, c, w, w, c, w, c, w, w},
		{w, w, c, w, w, w, w, w, w, w, w, w, w, w, w, w, w, c, w, w, c, c, c, w, w},
		{w, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w, w, w, w, c, w, w},
		{w, w, c, w, w, w, w, w, w, c, w, w, w, w, w, w, w, c, w, w, w, w, c, w, w},
		{w, w, c, w, w, w, w, w, w, c, w, w, w, w, w, w, w, c, c, c, c, w, c, w, w},
		{w, w, c, w, w, w, w, w, w, c, c, c, c, c, w, w, w, w, w, w, c, w, c, w, w},
		{w, c, c, w, w, w, w, w, w, w, w, w, w, c, w, w, w, w, w, w, c, c, c, c, w},
		{w, c, w, w, w, w, w, w, w, w, c, w, w, c, w, w, w, w, w, w, w, w, w, c, w},
		{w, c, w, w, w, w, w, c, c, c, c, w, w, c, c, c, c, c, c, w, w, w, w, c, w},
		{w, c, c, c, c, c, c, c, w, w, c, c, c, c, w, w, w, w, c, c, c, c, c, c, w},
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
	}
	path1 := []image.Point{
		{X: 7, Y: 12}, {X: 8, Y: 12}, {X: 9, Y: 12}, {X: 10, Y: 12},
	}
	path2 := []image.Point{
		{X: 10, Y: 13}, {X: 11, Y: 13}, {X: 12, Y: 13}, {X: 13, Y: 13},
	}
	path3 := []image.Point{
		{X: 8, Y: 6}, {X: 9, Y: 6}, {X: 10, Y: 6}, {X: 11, Y: 6}, {X: 12, Y: 6},
	}
	path4 := []image.Point{
		{X: 22, Y: 7}, {X: 22, Y: 6}, {X: 22, Y: 5}, {X: 22, Y: 6},
	}
	lvl.TurnOrderPattern = []interface{}{
		entities.NewPlayer(11, 2, config.TopBun),
		entities.NewPlayer(13, 2, config.Lettuce),
		entities.NewPlayer(14, 2, config.Cheese),
		entities.NewPlayer(2, 8, config.BottomBun),
		entities.NewPathEnemy(path1[0].X, path1[0].Y, path1, assets.Pidgeon),
		entities.NewPathEnemy(path2[0].X, path2[0].Y, path2, assets.Pidgeon),
		entities.NewPathEnemy(path3[0].X, path3[0].Y, path3, assets.Pidgeon),
		entities.NewPathEnemy(path4[0].X, path4[0].Y, path4, assets.Pidgeon),
	}
	return lvl
}

// A bit more challenging, bottom has to go via the down path (left-down)
func NewLevelLettuceMazeHard() *Level {
	w := Cell{Type: CellTypeWall}
	c := Cell{Type: CellTypeFloor}

	lvl := newLevel()
	lvl.BurgerPatty = entities.NewBurgerPatty(13, 9)
	lvl.Winning = []Position{
		{X: 10, Y: 11},
	}
	lvl.cells = [][]Cell{
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
		{w, w, w, w, w, w, w, w, c, c, c, c, c, c, c, w, w, w, w, w, w, w, w, w, w},
		{w, w, c, w, w, w, w, w, w, w, w, w, w, w, w, w, w, c, c, c, c, w, w, w, w},
		{w, w, c, w, w, w, w, w, w, w, w, w, w, w, w, w, w, c, w, w, c, w, c, w, w},
		{w, w, c, w, w, w, w, w, w, w, w, w, w, w, w, w, w, c, w, w, c, c, c, w, w},
		{w, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w, w, w, w, c, w, w},
		{w, w, c, w, w, c, w, w, w, c, w, w, w, c, w, w, w, c, w, w, w, w, c, w, w},
		{w, w, c, w, w, c, w, w, w, c, w, w, w, c, w, w, w, c, c, c, c, w, c, w, w},
		{w, w, c, w, w, c, w, c, c, c, c, c, c, c, c, c, c, c, w, w, c, w, c, w, w},
		{w, c, c, c, c, c, c, c, w, w, w, w, w, c, w, w, w, w, w, w, c, c, c, c, w},
		{w, c, w, w, w, w, w, c, c, c, c, w, w, c, w, w, w, w, w, w, w, w, w, c, w},
		{w, c, w, w, w, w, w, c, w, w, c, w, w, c, c, c, c, c, c, w, w, w, w, c, w},
		{w, c, c, c, c, c, c, c, w, w, c, c, c, c, w, w, w, w, c, c, c, c, c, c, w},
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
	}
	path1 := []image.Point{
		{X: 7, Y: 11}, {X: 8, Y: 11}, {X: 9, Y: 11}, {X: 10, Y: 11},
	}
	path2 := []image.Point{
		{X: 10, Y: 13}, {X: 11, Y: 13}, {X: 12, Y: 13}, {X: 13, Y: 13},
	}
	path3 := []image.Point{
		{X: 8, Y: 6}, {X: 9, Y: 6}, {X: 10, Y: 6}, {X: 11, Y: 6}, {X: 12, Y: 6},
	}
	path4 := []image.Point{
		{X: 22, Y: 7}, {X: 22, Y: 6}, {X: 22, Y: 5}, {X: 22, Y: 6},
	}
	path5 := []image.Point{
		{X: 7, Y: 9}, {X: 8, Y: 9}, {X: 9, Y: 9}, {X: 10, Y: 9}, {X: 11, Y: 9},
	}
	lvl.TurnOrderPattern = []interface{}{
		entities.NewPlayer(11, 2, config.TopBun),
		entities.NewPlayer(13, 2, config.Lettuce),
		entities.NewPlayer(2, 8, config.BottomBun),
		entities.NewPathEnemy(path1[0].X, path1[0].Y, path1, assets.Pidgeon),
		entities.NewPathEnemy(path2[0].X, path2[0].Y, path2, assets.Pidgeon),
		entities.NewPathEnemy(path3[0].X, path3[0].Y, path3, assets.Pidgeon),
		entities.NewPathEnemy(path4[0].X, path4[0].Y, path4, assets.Pidgeon),
		entities.NewPathEnemy(path5[0].X, path5[0].Y, path5, assets.Pidgeon),
	}
	return lvl
}

// Medium difficulty?
func NewLevelCheeseMaze() *Level {
	w := Cell{Type: CellTypeWall}
	c := Cell{Type: CellTypeFloor}

	lvl := newLevel()
	lvl.BurgerPatty = entities.NewBurgerPatty(13, 9)
	lvl.Winning = []Position{
		{X: 12, Y: 11},
	}
	lvl.cells = [][]Cell{
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
		{w, c, c, c, c, c, c, c, c, w, w, w, w, w, w, c, c, c, c, c, c, c, c, w, w},
		{w, c, w, w, w, w, w, w, c, c, c, c, c, c, c, c, w, w, w, w, w, w, c, w, w},
		{w, c, c, w, w, w, w, w, w, w, w, w, w, w, w, w, w, c, c, c, c, w, c, w, w},
		{w, w, c, w, w, w, w, w, w, w, w, w, w, w, w, w, w, c, w, w, c, w, c, w, w},
		{w, w, c, w, w, w, w, w, w, w, w, w, w, w, w, w, w, c, w, w, c, c, c, w, w},
		{w, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w, w, c, w, c, w, w},
		{w, w, c, w, w, c, w, w, w, c, w, w, w, c, w, w, w, c, c, w, c, w, c, w, w},
		{w, w, c, w, w, c, w, c, c, c, c, c, w, c, w, w, w, w, c, c, c, w, c, w, w},
		{w, w, c, w, w, c, w, c, w, w, w, c, c, c, c, c, c, c, c, w, c, w, c, w, w},
		{w, c, c, c, c, c, c, c, w, w, w, w, w, c, w, w, w, w, c, w, c, c, c, c, w},
		{w, c, w, w, c, w, w, c, w, w, w, w, c, c, w, w, w, w, c, w, w, w, w, c, w},
		{w, c, w, w, c, w, w, c, c, c, c, w, w, c, c, c, c, c, c, w, w, w, w, c, w},
		{w, c, c, c, c, c, c, c, w, w, c, c, c, c, w, w, w, w, c, c, c, c, c, c, w},
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
	}
	path1 := []image.Point{
		{X: 1, Y: 10}, {X: 2, Y: 10}, {X: 3, Y: 10}, {X: 4, Y: 10}, {X: 5, Y: 10}, {X: 6, Y: 10},
	}
	path2 := []image.Point{
		{X: 10, Y: 13}, {X: 11, Y: 13}, {X: 12, Y: 13}, {X: 13, Y: 13},
	}
	path3 := []image.Point{
		{X: 8, Y: 6}, {X: 9, Y: 6}, {X: 10, Y: 6}, {X: 11, Y: 6}, {X: 12, Y: 6},
	}
	path4 := []image.Point{
		{X: 22, Y: 7}, {X: 22, Y: 6}, {X: 22, Y: 5}, {X: 22, Y: 6},
	}
	path5 := []image.Point{
		{X: 7, Y: 8}, {X: 8, Y: 8}, {X: 9, Y: 8}, {X: 10, Y: 8}, {X: 11, Y: 8},
	}
	path6 := []image.Point{
		{X: 7, Y: 8}, {X: 7, Y: 9}, {X: 7, Y: 10}, {X: 7, Y: 11}, {X: 7, Y: 12}, {X: 7, Y: 13},
	}
	lvl.TurnOrderPattern = []interface{}{
		entities.NewPlayer(11, 2, config.TopBun),
		entities.NewPlayer(13, 2, config.Cheese),
		entities.NewPlayer(2, 8, config.BottomBun),
		entities.NewPathEnemy(path1[0].X, path1[0].Y, path1, assets.Pidgeon),
		entities.NewPathEnemy(path2[0].X, path2[0].Y, path2, assets.Pidgeon),
		entities.NewPathEnemy(path3[0].X, path3[0].Y, path3, assets.Pidgeon),
		entities.NewPathEnemy(path4[0].X, path4[0].Y, path4, assets.Pidgeon),
		entities.NewPathEnemy(path5[0].X, path5[0].Y, path5, assets.Pidgeon),
		entities.NewPathEnemy(path6[0].X, path6[0].Y, path6, assets.Pidgeon),
	}
	return lvl
}

func SnakesLevel() *Level {
	w := Cell{Type: CellTypeWall}
	c := Cell{Type: CellTypeFloor}
	lvl := newLevel()
	lvl.BurgerPatty = entities.NewBurgerPatty(11, 7)
	lvl.Winning = []Position{
		{X: 12, Y: 9},
	}
	lvl.cells = [][]Cell{
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
		{w, c, c, c, w, w, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, w, c, w, w, w, c, w, w, w, w, c, w, w, w, w, c, w, w, w, c, w, c, w},
		{w, c, w, c, w, w, w, c, w, c, c, c, c, w, w, w, w, c, w, w, w, c, w, c, w},
		{w, c, w, c, w, w, w, c, c, c, w, w, w, w, w, w, w, c, w, w, w, c, w, c, w},
		{w, c, w, c, w, w, w, c, w, w, c, c, c, c, c, c, c, c, w, w, w, c, w, c, w},
		{w, c, w, c, w, w, w, c, w, w, c, c, c, c, w, w, w, c, w, w, w, c, w, c, w},
		{w, c, w, c, w, w, w, c, w, w, c, c, c, c, w, w, w, c, w, w, w, c, w, c, w},
		{w, c, w, c, w, w, w, c, w, w, c, c, c, c, w, w, w, c, w, w, w, c, w, c, w},
		{w, c, w, c, w, w, w, c, w, w, c, c, c, c, w, w, w, c, w, w, w, c, w, c, w},
		{w, c, w, c, w, w, w, c, w, w, w, c, c, c, w, w, w, c, w, w, w, c, w, c, w},
		{w, c, w, c, w, w, w, c, c, c, c, c, w, w, w, w, w, c, w, w, w, c, w, c, w},
		{w, c, w, c, w, w, w, c, w, w, w, w, w, w, w, w, w, c, w, w, w, c, w, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w, w, w, c, c, c, w},
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
	}
	// examplePath := []image.Point{
	// 	{X: 10, Y: 4}, {X: 11, Y: 4}, {X: 12, Y: 4}, {X: 13, Y: 4}, {X: 14, Y: 4},
	// }
	lvl.TurnOrderPattern = []interface{}{
		entities.NewPlayer(1, 1, config.BottomBun),
		entities.NewPlayer(23, 13, config.TopBun),
		entities.NewDashingFollowerEnemy(1, 13, config.BottomBun, 3), // New dashing follower enemy targeting Lettuce, dashes every 5 turns
		// entities.NewEnemy(10, 10),
		//entities.NewFollowerEnemy(8, 9, config.Cheese), // New follower enemy targeting Cheese
		//entities.NewPathEnemy(examplePath[0].X, examplePath[0].Y, examplePath, assets.Pidgeon),
		entities.NewDashingFollowerEnemy(23, 1, config.BottomBun, 3), // New dashing follower enemy targeting Lettuce, dashes every 5 turns
		entities.NewDashingFollowerEnemy(12, 9, config.BottomBun, 3), // New dashing follower enemy targeting Lettuce, dashes every 5 turns
		entities.NewDashingFollowerEnemy(12, 6, config.BottomBun, 3), // New dashing follower enemy targeting Lettuce, dashes every 5 turns
		// entities.NewEnemy(10, 10),
		//entities.NewFollowerEnemy(8, 9, config.Cheese), // New follower enemy targeting Cheese
		//entities.NewPathEnemy(examplePath[0].X, examplePath[0].Y, examplePath, assets.Pidgeon),
	}
	lvl.WinningImg = common.GetImage(assets.Client)
	return lvl
}

// Merge both, dash first enemy north-est, and go down with both
func FourSnakes() *Level {
	w := Cell{Type: CellTypeWall}
	c := Cell{Type: CellTypeFloor}
	lvl := newLevel()
	lvl.BurgerPatty = entities.NewBurgerPatty(14, 10)
	lvl.Winning = []Position{
		{X: 14, Y: 5},
	}
	lvl.cells = [][]Cell{
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w, c, c, w},
		{w, c, w, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w, w, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, w, c, c, w, w, w, w, w, c, c, w, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, w, w, w, c, c, c, w, c, c, c, c, c, c, c, c, c, c, c, w, w, w, c, w},
		{w, c, c, c, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w, c, c, c, w},
		{w, c, c, c, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
	}
	lvl.TurnOrderPattern = []interface{}{
		entities.NewPlayer(4, 2, config.TopBun),
		entities.NewPlayer(20, 2, config.BottomBun),
		entities.NewPlayer(5, 2, config.Lettuce),
		entities.NewPlayer(19, 2, config.Cheese),
		entities.NewDashingFollowerEnemy(1, 1, config.TopBun, 3), // New dashing follower enemy targeting Lettuce, dashes every 5 turns
		// entities.NewEnemy(10, 10),
		//entities.NewFollowerEnemy(8, 9, config.Cheese), // New follower enemy targeting Cheese
		//entities.NewPathEnemy(examplePath[0].X, examplePath[0].Y, examplePath, assets.Pidgeon),
		entities.NewDashingFollowerEnemy(23, 1, config.BottomBun, 3), // New dashing follower enemy targeting Lettuce, dashes every 5 turns
		entities.NewDashingFollowerEnemy(23, 13, config.Lettuce, 3),  // New dashing follower enemy targeting Lettuce, dashes every 5 turns
		entities.NewDashingFollowerEnemy(1, 13, config.Cheese, 3),    // New dashing follower enemy targeting Lettuce, dashes every 5 turns

	}
	lvl.WinningImg = common.GetImage(assets.Client)
	return lvl
}

func FourSnakesReturn() *Level {
	w := Cell{Type: CellTypeWall}
	c := Cell{Type: CellTypeFloor}
	lvl := newLevel()
	lvl.BurgerPatty = entities.NewBurgerPatty(14, 10)
	lvl.Winning = []Position{
		{X: 14, Y: 5},
	}
	lvl.cells = [][]Cell{
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w, c, c, w},
		{w, w, w, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w, w, w, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, w, c, c, w, w, w, w, w, c, c, w, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, w, w, w, w, c, c, c, w, c, c, c, c, c, c, c, c, c, c, c, w, w, w, w, w},
		{w, c, c, c, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w, c, c, c, w},
		{w, c, c, c, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
	}
	lvl.TurnOrderPattern = []interface{}{
		entities.NewPlayer(1, 1, config.TopBun),
		entities.NewPlayer(23, 1, config.BottomBun),
		entities.NewPlayer(23, 13, config.Lettuce),
		entities.NewPlayer(1, 13, config.Cheese),
		entities.NewDashingFollowerEnemy(4, 2, config.TopBun, 3), // New dashing follower enemy targeting Lettuce, dashes every 5 turns
		// entities.NewEnemy(10, 10),
		//entities.NewFollowerEnemy(8, 9, config.Cheese), // New follower enemy targeting Cheese
		//entities.NewPathEnemy(examplePath[0].X, examplePath[0].Y, examplePath, assets.Pidgeon),
		entities.NewDashingFollowerEnemy(20, 2, config.BottomBun, 3), // New dashing follower enemy targeting Lettuce, dashes every 5 turns
		entities.NewDashingFollowerEnemy(19, 10, config.Lettuce, 3),  // New dashing follower enemy targeting Lettuce, dashes every 5 turns
		entities.NewDashingFollowerEnemy(5, 10, config.Cheese, 3),    // New dashing follower enemy targeting Lettuce, dashes every 5 turns

	}
	lvl.WinningImg = common.GetImage(assets.Client)
	return lvl
}
func NewEmptyLevel() *Level {
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

// Trick: Avoid the lettuce!
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
