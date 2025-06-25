package level

import (
	"math/rand/v2"

	"github.com/mikelangelon/unibun/config"
	"github.com/mikelangelon/unibun/entities"
)

type wallShape struct {
	width, height int
}

// NewRandomLevel generates a new level with random elements, as walls (wallShape).
func NewRandomLevel() *Level {
	lvl := NewLevel0()

	// Clear existing entities
	lvl.TurnOrderPattern = []interface{}{}
	lvl.BurgerPatty = nil
	lvl.Winning = []Position{}

	// Put some walls randomly
	placeRandomWalls(lvl)

	// Get available points to put entities.
	spawnPoints := getSpawnPoints(lvl)

	// Place winning tile in one of the spawnPoints
	if len(spawnPoints) > 0 {
		idx := rand.IntN(len(spawnPoints))
		winPos := spawnPoints[idx]
		lvl.Winning = append(lvl.Winning, Position{X: winPos.X, Y: winPos.Y})
		spawnPoints = append(spawnPoints[:idx], spawnPoints[idx+1:]...)
	}
	// Place Patty
	if len(spawnPoints) > 0 {
		idx := rand.IntN(len(spawnPoints))
		pattyPos := spawnPoints[idx]
		lvl.BurgerPatty = entities.NewBurgerPatty(pattyPos.X, pattyPos.Y)
		spawnPoints = append(spawnPoints[:idx], spawnPoints[idx+1:]...)
	}

	var turnOrder []interface{}
	// Place Players around
	playerTypes := []config.PlayerType{config.TopBun, config.BottomBun, config.Cheese, config.Lettuce}
	for _, pt := range playerTypes {
		if len(spawnPoints) == 0 {
			break
		}
		idx := rand.IntN(len(spawnPoints))
		pos := spawnPoints[idx]
		player := entities.NewPlayer(pos.X, pos.Y, pt)
		turnOrder = append(turnOrder, player)
		spawnPoints = append(spawnPoints[:idx], spawnPoints[idx+1:]...)
	}

	// Place Enemies
	numEnemies := 2 + rand.IntN(4) // random from 2 to 5
	for i := 0; i < numEnemies; i++ {
		if len(spawnPoints) == 0 {
			break
		}
		idx := rand.IntN(len(spawnPoints))
		pos := spawnPoints[idx]
		enemy := entities.NewEnemy(pos.X, pos.Y)
		turnOrder = append(turnOrder, enemy)
		spawnPoints = append(spawnPoints[:idx], spawnPoints[idx+1:]...)
	}

	// Shuffle turn order for more randomness
	rand.Shuffle(len(turnOrder), func(i, j int) {
		turnOrder[i], turnOrder[j] = turnOrder[j], turnOrder[i]
	})
	lvl.TurnOrderPattern = turnOrder

	return lvl
}

func placeRandomWalls(lvl *Level) {
	wallShapes := []wallShape{
		{width: 5, height: 1}, {width: 1, height: 5},
		{width: 4, height: 1}, {width: 1, height: 4},
		{width: 3, height: 1}, {width: 1, height: 3},
		{width: 2, height: 1}, {width: 1, height: 2},
	}

	rand.Shuffle(len(wallShapes), func(i, j int) {
		wallShapes[i], wallShapes[j] = wallShapes[j], wallShapes[i]
	})

	wallsToPlace := 4
	for i := 0; i < wallsToPlace; i++ {
		shape := wallShapes[i]
		// tries 10 times to put a wall.
		for try := 0; try < 10; try++ {
			randX := 1 + rand.IntN(lvl.gridCols()-2-shape.width)
			randY := 1 + rand.IntN(lvl.gridRows()-2-shape.height)

			if lvl.GetCell(randX, randY).Type == CellTypeFloor {
				for y := 0; y < shape.height; y++ {
					for x := 0; x < shape.width; x++ {
						lvl.cells[randY+y][randX+x].Type = CellTypeWall
					}
				}
				break // wall placed, bye bye
			}
		}
	}
}

// getSpawnPoints collects all floor cells as valid positions for entities.
func getSpawnPoints(lvl *Level) []Position {
	var points []Position
	for r, row := range lvl.cells {
		for c, cell := range row {
			if cell.Type == CellTypeFloor {
				points = append(points, Position{X: c, Y: r})
			}
		}
	}
	return points
}
