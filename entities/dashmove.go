package entities

import "github.com/mikelangelon/unibun/config"

type dashMove struct {
	dashPath        []struct{ X, Y int }
	dashCurrentStep int
	dashStepTimer   int
}

// StartDash calculates the path for a dash and initializes the dashing state.
func (p *Player) StartDash(level Level, dx, dy int) bool {
	p.dashMove = nil

	tempPath := []struct{ X, Y int }{}
	currentX, currentY := p.GridX, p.GridY

	for {
		nextX, nextY := currentX+dx, currentY+dy
		if level.IsWalkable(nextX, nextY) {
			tempPath = append(tempPath, struct{ X, Y int }{nextX, nextY})
			currentX = nextX
			currentY = nextY
		} else {
			// Hit an obstacle or edge of the map
			break
		}
	}

	if len(tempPath) == 0 {
		return false // No valid path for dash (already on wall)
	}

	p.dashMove = &dashMove{
		dashPath:        tempPath,
		dashCurrentStep: 0,
		dashStepTimer:   0,
	}
	return true
}

// processDashStep handles the animation of an ongoing dash.
// Returns true when the dash is complete
func (p *Player) processDashStep() bool {
	if p.dashMove.dashStepTimer > 0 {
		p.dashMove.dashStepTimer--
		return false
	}

	// Moving to next tile
	if p.dashMove.dashCurrentStep < len(p.dashMove.dashPath) {
		nextPos := p.dashMove.dashPath[p.dashMove.dashCurrentStep]
		p.GridX = nextPos.X
		p.GridY = nextPos.Y

		p.dashMove.dashCurrentStep++
		p.dashMove.dashStepTimer = config.DashStepDelay

		if p.dashMove.dashCurrentStep >= len(p.dashMove.dashPath) {
			p.dashMove = nil
			return true
		}
		return false // Moved one step in dash, but dash is not over yet
	}

	// TODO: can arrive here?
	p.dashMove = nil
	return true
}
