package level

import (
	"math"
	"math/rand/v2"

	"github.com/mikelangelon/unibun/config"
	"github.com/mikelangelon/unibun/entities"
)

type wallShape struct {
	width, height int
}

func getNumberEnemies(lvlIndex int) int {
	min, maxRand := 2, 3
	if lvlIndex >= 5 && lvlIndex < 10 {
		min, maxRand = 3, 4
	} else if lvlIndex >= 10 && lvlIndex < 15 {
		min, maxRand = 4, 6
	} else if lvlIndex >= 15 {
		min, maxRand = 4, 8
	}
	return min + rand.IntN(maxRand)
}

// NewEndlessLevel generates a new level with random elements, as walls (wallShape).
func NewEndlessLevel(lvlIndex int) *Level {
	lvl := NewEmptyLevel()
	lvl.tilesPatterns = randomTilePattern()

	// Clear existing entities
	lvl.TurnOrderPattern = []interface{}{}
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

	var availablePlayerTypes []config.PlayerType
	for _, char := range turnOrder {
		if p, ok := char.(entities.Player); ok {
			availablePlayerTypes = append(availablePlayerTypes, p.PlayerType)
		}
	}

	// Place Enemies
	numEnemies := getNumberEnemies(lvlIndex)
	for i := 0; i < numEnemies; i++ {
		if len(spawnPoints) == 0 {
			break
		}
		idx := rand.IntN(len(spawnPoints))
		pos := spawnPoints[idx]

		// max progression level is 20, from there, all the same.
		roofLevel := math.Min(20.0, float64(lvlIndex))
		progress := roofLevel / 20.0

		/**
		Distribute the 3 enemy types, leaving the path type outside because no idea how to randomize it properly
		Pigeon is the easiest, dash follower medium and follower the hardest.
		from 0 to 100
		- Pigeon would be between 0 to 50
		- DashFollower would be between Pigeon value and 60
		- Follower just the last threshold

		*/

		enemyThreshold := 70.0 - (50.0 * progress)
		dashThreshold := 20 + enemyThreshold + (20.0 * progress)

		var enemy entities.Enemier
		r := rand.Float64() * 100

		if len(availablePlayerTypes) == 0 || r < enemyThreshold {
			enemy = entities.NewPigeon(pos.X, pos.Y)
		} else {
			targetType := availablePlayerTypes[rand.IntN(len(availablePlayerTypes))]
			if r < dashThreshold {
				minTurns := int(3.0 - 2.0*progress)
				maxTurns := int(4.0 - 1.0*progress)
				turnsToDash := rand.IntN(maxTurns-minTurns+1) + minTurns
				enemy = entities.NewDashingFollowerEnemy(pos.X, pos.Y, targetType, turnsToDash)
			} else {
				enemy = entities.NewDuck(pos.X, pos.Y, targetType)
			}
		}
		turnOrder = append(turnOrder, enemy)
		spawnPoints = append(spawnPoints[:idx], spawnPoints[idx+1:]...)
	}
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
