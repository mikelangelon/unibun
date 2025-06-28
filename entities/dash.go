package entities

import (
	"github.com/mikelangelon/unibun/config"
	"image"
)

type DashState struct {
	path        []image.Point
	currentStep int
	stepTimer   int

	// isActive checks if the dash still is active.... TODO Maybe could this be defined if DashState is not nil?
	isActive bool
}

func NewDashState() *DashState {
	return &DashState{}
}

func (ds *DashState) Start(startX, startY, dx, dy int, speed int, level Level, canWalkThroughWalls bool) bool {
	ds.path = nil
	ds.currentStep = 0
	ds.stepTimer = 0
	ds.isActive = false

	tempPath := []image.Point{}
	currentX, currentY := startX, startY

	for i := 1; i <= speed; i++ {
		nextX, nextY := currentX+dx, currentY+dy
		// TODO code is duplicated in Dash & normal movement --> Extract common parts
		if canWalkThroughWalls && !level.OutOfBounds(nextX, nextY) {
			tempPath = append(tempPath, image.Point{X: nextX, Y: nextY})
			currentX = nextX
			currentY = nextY
			continue
		}

		if !level.IsWalkable(nextX, nextY) {
			break
		}
		tempPath = append(tempPath, image.Point{X: nextX, Y: nextY})
		currentX = nextX
		currentY = nextY
	}

	if len(tempPath) == 0 {
		return false
	}

	ds.path = tempPath
	ds.isActive = true
	return true
}

// TODO Extend logic to combinate it with walk through walls
func (ds *DashState) Update(currentX, currentY int, level Level, canWalkThroughWalls bool) (newX, newY int, movedThisFrame bool, finished bool) {
	if !ds.isActive {
		return currentX, currentY, false, false
	}

	if ds.stepTimer > 0 {
		ds.stepTimer--
		return currentX, currentY, false, false
	}

	if ds.currentStep < len(ds.path) {
		nextPos := ds.path[ds.currentStep]
		newX, newY = nextPos.X, nextPos.Y
		ds.currentStep++
		ds.stepTimer = config.DashStepDelay
		movedThisFrame = true

		if ds.currentStep >= len(ds.path) {
			ds.isActive = false
			finished = true
		}
		return newX, newY, movedThisFrame, finished
	}

	ds.isActive = false
	return currentX, currentY, false, true
}

func (ds *DashState) IsActive() bool {
	return ds.isActive
}

func (ds *DashState) Reset() {
	ds.path = nil
	ds.currentStep = 0
	ds.stepTimer = 0
	ds.isActive = false
}
