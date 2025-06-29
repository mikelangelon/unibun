package level

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/mikelangelon/unibun/assets"
	"github.com/mikelangelon/unibun/common"
	"github.com/mikelangelon/unibun/config"
	"github.com/mikelangelon/unibun/entities"
)

func NewIntro() *Level {
	w := Cell{Type: CellTypeWall}
	c := Cell{Type: CellTypeFloor}
	lvl := newLevel()
	lvl.BurgerPatty = entities.NewBurgerPatty(13, 4)
	lvl.Winning = []Position{
		{X: 20, Y: 9},
	}
	lvl.cells = [][]Cell{
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, w, w, c, c, c, c, c, c, c, c, w, w, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, w, c, c, w, w, w, w, c, c, c, c, c, c, c, c, w, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, w, w, w, w, c, c, w, w, w, w, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
	}
	lvl.IntroText = "Your goal is to unite the Buns with the Patty,\nand deliver it to the client.\n" +
		"Buns move 2 positions.\n\n\n\nPress Enter to continue"

	path1 := []image.Point{
		{X: 15, Y: 7}, {X: 16, Y: 7}, {X: 17, Y: 7}, {X: 18, Y: 7}, {X: 19, Y: 7}, {X: 20, Y: 7}, {X: 21, Y: 7},
	}
	path2 := []image.Point{
		{X: 14, Y: 1}, {X: 14, Y: 2}, {X: 14, Y: 3}, {X: 14, Y: 4}, {X: 14, Y: 5}, {X: 14, Y: 6},
	}
	lvl.TurnOrderPattern = []interface{}{
		entities.NewPlayer(3, 1, config.TopBun),
		entities.NewPlayer(21, 1, config.BottomBun),
		entities.NewMouse(path1[0].X, path1[0].Y, path1),
		entities.NewMouse(path2[0].X, path2[0].Y, path2),
		entities.NewPigeon(4, 4),
		entities.NewPigeon(10, 6),
	}
	lvl.WinningImg = common.GetImage(assets.Client)
	return lvl
}

func LettucePresentation() *Level {
	w := Cell{Type: CellTypeWall}
	c := Cell{Type: CellTypeFloor}

	lvl := newLevel()
	lvl.BurgerPatty = entities.NewBurgerPatty(11, 6)
	lvl.Winning = []Position{
		{X: 2, Y: 10},
	}
	lvl.IntroText = "If a Bun unites with the Lettuce, \nit would be able to cross walls\n\n\n\n\nPress Enter to continue"
	lvl.cells = [][]Cell{
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w, w, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w, w, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, w, w, w, w, w, c, c, c, c, c, c, c, c, w, w, w},
		{w, c, c, c, c, c, c, c, c, w, w, w, w, w, c, c, c, c, c, c, c, c, w, w, w},
		{w, w, w, w, w, w, w, c, c, w, w, c, w, w, c, c, c, c, c, c, c, c, w, w, w},
		{w, c, c, c, c, c, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, w, w, w, w, c, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, w, c, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, w, c, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
	}
	path1 := []image.Point{
		{X: 7, Y: 1}, {X: 7, Y: 2}, {X: 7, Y: 3}, {X: 7, Y: 4}, {X: 7, Y: 5}, {X: 7, Y: 6}, {X: 7, Y: 7}, {X: 7, Y: 8}, {X: 7, Y: 9},
	}
	path2 := []image.Point{
		{X: 11, Y: 7}, {X: 12, Y: 7}, {X: 13, Y: 7}, {X: 14, Y: 7},
	}
	path3 := []image.Point{
		{X: 5, Y: 7}, {X: 5, Y: 8}, {X: 5, Y: 9}, {X: 5, Y: 10},
	}
	path4 := []image.Point{
		{X: 1, Y: 7}, {X: 2, Y: 7}, {X: 3, Y: 7}, {X: 4, Y: 7}, {X: 5, Y: 7},
	}
	lvl.TurnOrderPattern = []interface{}{
		entities.NewPlayer(13, 1, config.TopBun),
		entities.NewPlayer(7, 9, config.BottomBun),
		entities.NewPlayer(11, 3, config.Lettuce),
		entities.NewPigeon(10, 10),
		entities.NewMouse(path1[0].X, path1[0].Y, path1),
		entities.NewMouse(path2[0].X, path2[0].Y, path2),
		entities.NewMouse(path3[0].X, path3[0].Y, path3),
		entities.NewMouse(path4[0].X, path4[0].Y, path4),
	}
	return lvl
}

func CheesePresentation() *Level {
	w := Cell{Type: CellTypeWall}
	c := Cell{Type: CellTypeFloor}

	lvl := newLevel()
	lvl.BurgerPatty = entities.NewBurgerPatty(2, 13)
	lvl.Winning = []Position{
		{X: 4, Y: 13},
	}
	lvl.IntroText = "If a Bun unites with Cheese \nPress Z + direction to charge a dash next turn\n While dashing, it will kill enemies.\n\n\n\nPress Enter to continue"
	lvl.cells = [][]Cell{
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
		{w, c, c, c, c, w, w, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, w, w, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, w, w, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, w, w, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, w, w, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, w, w, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, w, w, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
	}
	path1 := []image.Point{
		{X: 1, Y: 6}, {X: 2, Y: 6}, {X: 3, Y: 6}, {X: 4, Y: 6},
	}
	path2 := []image.Point{
		{X: 3, Y: 7}, {X: 4, Y: 7}, {X: 3, Y: 7}, {X: 2, Y: 7}, {X: 1, Y: 7}, {X: 2, Y: 7},
	}
	path3 := []image.Point{
		{X: 2, Y: 8}, {X: 1, Y: 8}, {X: 2, Y: 8}, {X: 3, Y: 8}, {X: 4, Y: 8}, {X: 3, Y: 8},
	}
	path4 := []image.Point{
		{X: 4, Y: 9}, {X: 3, Y: 9}, {X: 2, Y: 9}, {X: 1, Y: 9},
	}
	path5 := []image.Point{
		{X: 2, Y: 10}, {X: 3, Y: 10}, {X: 4, Y: 10}, {X: 3, Y: 10}, {X: 2, Y: 10}, {X: 1, Y: 10},
	}
	lvl.TurnOrderPattern = []interface{}{
		entities.NewPlayer(2, 1, config.TopBun),
		entities.NewPlayer(5, 3, config.BottomBun),
		entities.NewPlayer(2, 2, config.Cheese),
		entities.NewMouse(path1[0].X, path1[0].Y, path1),
		entities.NewMouse(path2[0].X, path2[0].Y, path2),
		entities.NewMouse(path3[0].X, path3[0].Y, path3),
		entities.NewMouse(path4[0].X, path4[0].Y, path4),
		entities.NewMouse(path5[0].X, path5[0].Y, path5),
	}
	return lvl
}

func NewFlies() *Level {
	w := Cell{Type: CellTypeWall}
	c := Cell{Type: CellTypeFloor}
	lvl := newLevel()
	lvl.tilesPatterns = tilesPatterns{
		colorScaleOdd:  tileB,
		colorScaleEven: tileE,
	}
	lvl.BurgerPatty = entities.NewBurgerPatty(13, 4)
	lvl.Winning = []Position{
		{X: 13, Y: 3},
	}
	lvl.cells = [][]Cell{
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
		{w, c, c, c, c, c, w, c, c, c, c, c, c, c, c, c, c, w, c, c, c, c, c, c, w},
		{w, c, c, w, w, w, w, c, c, c, c, c, c, c, c, c, c, w, c, c, c, c, c, c, w},
		{w, c, c, w, c, c, c, c, c, c, c, c, c, c, c, c, c, w, c, c, c, c, c, c, w},
		{w, w, c, w, c, c, c, w, w, c, c, c, c, c, c, c, c, w, w, c, c, c, c, c, w},
		{w, w, c, w, c, c, c, c, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, w, c, c, c, c, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, w, c, c, c, c, c, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, w, w, c, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, w, w, w, w, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
	}
	lvl.TurnOrderPattern = []interface{}{
		entities.NewPlayer(3, 1, config.TopBun),
		entities.NewPlayer(21, 1, config.BottomBun),
		entities.NewPlayer(20, 2, config.Cheese),
		entities.NewFly(1, 13),
		entities.NewFly(1, 11),
		entities.NewFly(23, 11),
		entities.NewFly(22, 13),
		entities.NewFly(14, 13),
		entities.NewFly(11, 11),
	}
	lvl.WinningImg = common.GetImage(assets.Client)
	return lvl
}

func PushThePatty() *Level {
	w := Cell{Type: CellTypeWall}
	c := Cell{Type: CellTypeFloor}
	lvl := newLevel()
	lvl.BurgerPatty = entities.NewBurgerPatty(4, 4)
	lvl.Winning = []Position{
		{X: 9, Y: 10},
	}
	lvl.cells = [][]Cell{
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
		{w, c, c, c, c, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, w, c, c, c, c, c, c, w, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, w, c, w, w, w, c, c, w, w, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, w, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, w, w, w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, w, c, c, w, w, c, c, c, c, c, c, w, w, w, w, w},
		{w, c, c, c, c, c, c, w, c, c, c, c, c, w, w, w, w, w, w, w, c, c, c, c, w},
		{w, c, c, c, c, w, c, c, c, w, c, c, w, c, w, c, c, c, c, w, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, w, c, c, c, c, w, c, c, c, c, w},
		{w, c, c, c, c, c, c, w, c, c, c, c, c, c, w, c, c, c, w, w, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, w, c, c, c, c, c, c, c, c, c, w},
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
	}
	lvl.TurnOrderPattern = []interface{}{
		entities.NewPlayer(2, 2, config.TopBun),
		entities.NewPlayer(23, 13, config.BottomBun),
		entities.NewFly(1, 1),
		entities.NewFly(1, 5),
		entities.NewFly(1, 7),
		entities.NewFly(1, 9),
		entities.NewDashingFollowerEnemy(23, 1, config.BottomBun, 3),
		entities.NewDashingFollowerEnemy(22, 2, config.TopBun, 2),
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
	lvl.tilesPatterns = tilesPatterns{
		colorScaleOdd:  tileH,
		colorScaleEven: tileA,
	}
	lvl.TurnOrderPattern = []interface{}{
		entities.NewPlayer(5, 1, config.TopBun),
		entities.NewPlayer(5, 6, config.BottomBun),
		entities.NewPlayer(7, 5, config.Cheese),
		entities.NewMouse(path[0].X, path[0].Y, path),
		entities.NewMouse(path2[0].X, path2[0].Y, path2),
		entities.NewDashingFollowerEnemy(4, 1, config.TopBun, 3),
		entities.NewDashingFollowerEnemy(8, 8, config.TopBun, 3),
	}
	lvl.WinningImg = common.GetImage(assets.Client)
	return lvl
}

func AnotherLettuce() *Level {
	w := Cell{Type: CellTypeWall}
	c := Cell{Type: CellTypeFloor}

	lvl := newLevel()
	lvl.tilesPatterns = tilesPatterns{
		colorScaleOdd:  tileH,
		colorScaleEven: tileB,
	}
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
		entities.NewPigeon(10, 10),
		entities.NewMouse(path1[0].X, path1[0].Y, path1),
		entities.NewMouse(path2[0].X, path2[0].Y, path2),
	}
	return lvl
}

// Take the cheese & lettuce and run against all!
func LettuceCheeseMaze() *Level {
	w := Cell{Type: CellTypeWall}
	c := Cell{Type: CellTypeFloor}

	lvl := newLevel()
	lvl.tilesPatterns = tilesPatterns{
		colorScaleOdd:  tileE,
		colorScaleEven: tileC,
	}
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
		entities.NewMouse(path1[0].X, path1[0].Y, path1),
		entities.NewMouse(path2[0].X, path2[0].Y, path2),
		entities.NewMouse(path3[0].X, path3[0].Y, path3),
		entities.NewMouse(path4[0].X, path4[0].Y, path4),
	}
	return lvl
}

// A bit more challenging, bottom has to go via the down path (left-down)
func NewLevelLettuceMazeHard() *Level {
	w := Cell{Type: CellTypeWall}
	c := Cell{Type: CellTypeFloor}

	lvl := newLevel()
	lvl.tilesPatterns = tilesPatterns{
		colorScaleOdd:  tileA,
		colorScaleEven: tileD,
	}
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
		{w, c, w, w, w, w, w, c, w, w, c, w, w, c, w, w, c, c, c, w, w, w, w, c, w},
		{w, c, c, c, c, c, c, c, w, w, c, c, c, c, c, c, c, w, c, c, c, c, c, c, w},
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
	}
	path1 := []image.Point{
		{X: 7, Y: 11}, {X: 8, Y: 11}, {X: 9, Y: 11}, {X: 10, Y: 11},
	}
	path2 := []image.Point{
		{X: 10, Y: 13}, {X: 11, Y: 13}, {X: 12, Y: 13}, {X: 13, Y: 13}, {X: 14, Y: 13}, {X: 15, Y: 13},
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
		entities.NewMouse(path1[0].X, path1[0].Y, path1),
		entities.NewMouse(path2[0].X, path2[0].Y, path2),
		entities.NewMouse(path3[0].X, path3[0].Y, path3),
		entities.NewMouse(path4[0].X, path4[0].Y, path4),
		entities.NewMouse(path5[0].X, path5[0].Y, path5),
	}
	return lvl
}

// Medium difficulty?
func NotUsed() *Level {
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
		entities.NewMouse(path1[0].X, path1[0].Y, path1),
		entities.NewMouse(path2[0].X, path2[0].Y, path2),
		entities.NewMouse(path3[0].X, path3[0].Y, path3),
		entities.NewMouse(path4[0].X, path4[0].Y, path4),
		entities.NewMouse(path5[0].X, path5[0].Y, path5),
		entities.NewMouse(path6[0].X, path6[0].Y, path6),
	}
	return lvl
}

func SnakesLevel() *Level {
	w := Cell{Type: CellTypeWall}
	c := Cell{Type: CellTypeFloor}
	lvl := newLevel()
	lvl.tilesPatterns = tilesPatterns{
		colorScaleOdd:  tileE,
		colorScaleEven: tileH,
	}
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

	lvl.TurnOrderPattern = []interface{}{
		entities.NewPlayer(1, 1, config.BottomBun),
		entities.NewPlayer(23, 13, config.TopBun),
		entities.NewDashingFollowerEnemy(1, 13, config.BottomBun, 3),
		entities.NewDashingFollowerEnemy(23, 1, config.BottomBun, 3),
		entities.NewDashingFollowerEnemy(12, 9, config.BottomBun, 3),
		entities.NewDashingFollowerEnemy(12, 6, config.BottomBun, 3),
	}
	lvl.WinningImg = common.GetImage(assets.Client)
	return lvl
}

// Merge both, dash first enemy north-est, and go down with both
func FourSnakes() *Level {
	w := Cell{Type: CellTypeWall}
	c := Cell{Type: CellTypeFloor}
	lvl := newLevel()
	lvl.tilesPatterns = tilesPatterns{
		colorScaleOdd:  tileC,
		colorScaleEven: tileA,
	}
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
		entities.NewDashingFollowerEnemy(1, 1, config.TopBun, 3),
		entities.NewDashingFollowerEnemy(23, 1, config.BottomBun, 3),
		entities.NewDashingFollowerEnemy(23, 13, config.Lettuce, 3),
		entities.NewDashingFollowerEnemy(1, 13, config.Cheese, 3),
	}
	lvl.WinningImg = common.GetImage(assets.Client)
	return lvl
}

func FourSnakesReturn() *Level {
	w := Cell{Type: CellTypeWall}
	c := Cell{Type: CellTypeFloor}
	lvl := newLevel()
	lvl.tilesPatterns = tilesPatterns{
		colorScaleOdd:  tileA,
		colorScaleEven: tileGreen,
	}
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
		entities.NewDashingFollowerEnemy(4, 2, config.TopBun, 3),
		entities.NewDashingFollowerEnemy(20, 2, config.BottomBun, 3),
		entities.NewDashingFollowerEnemy(19, 10, config.Lettuce, 3),
		entities.NewDashingFollowerEnemy(5, 10, config.Cheese, 3),
	}
	lvl.WinningImg = common.GetImage(assets.Client)
	return lvl
}

func FirstRealLevel() *Level {
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
	lvl.tilesPatterns = tilesPatterns{
		colorScaleOdd:  tileF,
		colorScaleEven: tileA,
	}
	lvl.TurnOrderPattern = []interface{}{
		entities.NewPlayer(3, 1, config.TopBun),
		entities.NewPlayer(10, 5, config.BottomBun),
		entities.NewPigeon(11, 6),
		entities.NewDuck(1, 9, config.TopBun),
		entities.NewDuck(16, 1, config.BottomBun),
		entities.NewDashingFollowerEnemy(18, 4, config.BottomBun, 4),
	}
	return lvl
}

func AvoidTheLettuce() *Level {
	w := Cell{Type: CellTypeWall}
	c := Cell{Type: CellTypeFloor}
	lvl := newLevel()
	lvl.tilesPatterns = tilesPatterns{
		colorScaleOdd:  tileA,
		colorScaleEven: tileGreen2,
	}
	lvl.BurgerPatty = entities.NewBurgerPatty(21, 12)
	lvl.Winning = []Position{
		{X: 6, Y: 7},
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
		entities.NewDuck(3, 10, config.TopBun),
		entities.NewDuck(5, 6, config.BottomBun),
	}
	return lvl
}

func PuzzleBuns() *Level {
	w := Cell{Type: CellTypeWall}
	c := Cell{Type: CellTypeFloor}
	lvl := newLevel()
	lvl.tilesPatterns = tilesPatterns{
		colorScaleOdd:  tileZ,
		colorScaleEven: tileG,
	}
	lvl.BurgerPatty = entities.NewBurgerPatty(12, 7)
	lvl.Winning = []Position{
		{X: 23, Y: 1},
	}
	lvl.cells = [][]Cell{
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
		{w, c, c, c, c, c, w, c, w, w, w, w, w, w, w, c, w, c, w, c, w, c, c, c, w},
		{w, c, w, c, c, c, w, c, w, c, c, w, c, c, w, c, w, c, w, c, c, c, w, c, w},
		{w, c, w, c, w, c, w, c, w, c, c, w, c, c, w, c, w, c, w, c, w, c, w, c, w},
		{w, c, w, c, c, c, w, c, w, c, c, w, c, c, w, c, w, c, c, c, w, c, c, c, w},
		{w, c, c, c, c, c, w, c, w, w, w, w, w, w, w, c, w, c, w, c, w, c, w, w, w},
		{w, c, w, c, w, c, w, c, c, c, c, c, c, c, c, c, w, c, w, c, w, c, c, c, w},
		{w, c, w, c, c, c, c, c, c, c, w, c, c, c, c, c, w, c, c, c, w, c, c, c, w},
		{w, c, w, c, w, c, w, c, c, c, c, c, c, c, c, c, c, c, w, c, w, c, w, c, w},
		{w, c, c, c, c, c, c, c, w, w, w, w, w, w, w, c, w, c, w, c, w, c, c, c, w},
		{w, c, w, c, c, c, w, c, w, c, c, w, c, c, w, c, w, c, c, c, w, c, w, c, w},
		{w, c, w, c, w, c, w, c, w, c, c, w, c, c, w, c, w, c, w, c, w, c, c, c, w},
		{w, c, w, c, c, c, w, c, w, c, c, w, c, c, w, c, w, c, w, c, c, c, w, c, w},
		{w, c, c, c, c, c, w, c, w, w, w, w, w, w, w, c, w, c, w, c, w, c, c, c, w},
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
	}

	lvl.TurnOrderPattern = []interface{}{
		entities.NewPlayer(1, 1, config.TopBun),
		entities.NewPlayer(1, 13, config.BottomBun),
		entities.NewPigeon(6, 2),
		entities.NewPigeon(18, 12),
		entities.NewPigeon(20, 5),
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
		FloorTileImg: []*ebiten.Image{
			common.GetImage(assets.FloorTile),
			common.GetImage(assets.FloorTileB),
			common.GetImage(assets.FloorTileC),
			common.GetImage(assets.FloorTileD),
			common.GetImage(assets.FloorTileE),
			common.GetImage(assets.FloorTileF),
		},
		pulseOffset: 0,
		tilesPatterns: tilesPatterns{
			colorScaleOdd:  tileA,
			colorScaleEven: tileE,
		},
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
		entities.NewPlayer(23, 1, config.TopBun),
		entities.NewPlayer(23, 13, config.BottomBun),
		entities.NewPigeon(5, 5),
		entities.NewPigeon(10, 7),
		entities.NewPigeon(17, 9),
	}
	return lvl
}

// Not used, for testing
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
	lvl.TurnOrderPattern = []interface{}{
		entities.NewPlayer(3, 1, config.TopBun),
		entities.NewPlayer(3, 5, config.BottomBun),
	}
	lvl.WinningImg = common.GetImage(assets.Client)
	return lvl
}

func Unite() *Level {
	w := Cell{Type: CellTypeWall}
	c := Cell{Type: CellTypeFloor}
	lvl := newLevel()
	lvl.tilesPatterns = tilesPatterns{
		colorScaleOdd:  tileUniteA,
		colorScaleEven: tileUniteB,
	}
	lvl.BurgerPatty = entities.NewBurgerPatty(20, 7)
	lvl.Winning = []Position{
		{X: 11, Y: 2},
	}
	lvl.cells = [][]Cell{
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, c, w},
		{w, c, w, c, w, c, w, c, c, w, c, w, c, w, w, w, c, w, w, w, w, c, c, c, w},
		{w, c, w, c, w, c, w, c, c, w, c, w, c, c, w, c, c, w, c, c, c, c, c, c, w},
		{w, c, w, c, w, c, w, c, w, w, c, w, c, c, w, c, c, w, w, w, w, c, c, c, w},
		{w, c, w, c, w, c, w, w, c, w, c, w, c, c, w, c, c, w, c, c, c, c, c, c, w},
		{w, c, w, w, w, c, w, c, c, w, c, w, c, c, w, c, c, w, w, w, w, c, c, c, w},
		{w, c, c, c, c, c, c, c, c, c, c, c, c, c, w, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, w, c, c, c, w, c, c, c, c, w, c, c, c, c, c, c, c, c, c, w},
		{w, w, w, w, w, w, c, w, c, c, c, c, c, c, w, c, c, c, c, c, c, c, c, c, w},
		{w, c, c, c, c, c, c, w, c, c, c, c, c, c, w, c, c, c, c, c, c, c, c, c, w},
		{w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w, w},
	}
	path1 := []image.Point{
		{X: 14, Y: 2}, {X: 14, Y: 3}, {X: 14, Y: 4}, {X: 14, Y: 5},
	}
	path2 := []image.Point{
		{X: 15, Y: 7}, {X: 16, Y: 7},
	}
	lvl.TurnOrderPattern = []interface{}{
		entities.NewPlayer(1, 1, config.TopBun),
		entities.NewPlayer(22, 7, config.BottomBun),
		entities.NewPlayer(1, 13, config.Lettuce),
		entities.NewPlayer(18, 8, config.Cheese),
		entities.NewPigeon(23, 2),
		entities.NewPigeon(18, 12),
		entities.NewPigeon(8, 8),
		entities.NewFly(23, 13),
		entities.NewFly(19, 13),
		entities.NewFly(23, 1),
		entities.NewFly(19, 1),
		entities.NewDashingFollowerEnemy(10, 1, config.TopBun, 2),
		entities.NewDashingFollowerEnemy(15, 13, config.BottomBun, 2),
		entities.NewDuck(3, 10, config.TopBun),
		entities.NewDuck(5, 6, config.BottomBun),
		entities.NewMouse(path1[0].X, path1[0].Y, path1),
		entities.NewMouse(path2[0].X, path2[0].Y, path2),
	}
	return lvl
}
