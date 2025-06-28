package entities

import (
	"image"
	"math"

	"github.com/mikelangelon/unibun/common"

	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mikelangelon/unibun/assets"
	"github.com/mikelangelon/unibun/config"
)

type Player struct {
	GridX, GridY        int
	PlayerType          config.PlayerType
	Image               *ebiten.Image
	CanDash             bool
	CanWalkThroughWalls bool
	IsActiveTurn        bool
	pulseOffset         float64
	Speed               int

	initialGridX, initialGridY int
	initialCanDash             bool
	initialCanWalkThroughWalls bool

	dashState *DashState
}

type Level interface {
	IsWalkable(gridX, gridY int) bool
}

func NewPlayer(startX, startY int, playerType config.PlayerType) Player {
	var b = []byte{}
	switch playerType {
	case config.TopBun:
		b = assets.TopBun
	case config.BottomBun:
		b = assets.BottomBun
	case config.Cheese:
		b = assets.Cheese
	case config.Lettuce:
		b = assets.Lettuce
	case config.BurguerPatty:
		b = assets.BurgerPatty

	}
	img := common.GetImage(b)
	// calculate offsets for centering
	offsettedImg := ebiten.NewImage(32, 32)
	offsetY := float64(32-18) / 2.0
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, offsetY)
	offsettedImg.DrawImage(img, op)

	p := Player{
		GridX:                      startX,
		GridY:                      startY,
		Image:                      offsettedImg,
		PlayerType:                 playerType,
		CanDash:                    false,
		CanWalkThroughWalls:        false,
		Speed:                      2,
		initialGridX:               startX,
		initialGridY:               startY,
		initialCanDash:             false,
		initialCanWalkThroughWalls: false,
		dashState:                  NewDashState(),
	}
	return p
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	if p.IsDashing() {
		op.ColorScale.Scale(1.5, 1.8, 0.5, 0.8)
	}
	if p.IsActiveTurn {
		brightness := 1.25 + 0.25*math.Sin(p.pulseOffset/11.0)
		op.ColorScale.Scale(float32(brightness), float32(brightness), float32(brightness), 1.0)
		p.pulseOffset++
	}
	pixelX := float64(p.GridX * config.TileSize)
	pixelY := float64(p.GridY * config.TileSize)
	op.GeoM.Translate(pixelX, pixelY)
	screen.DrawImage(p.Image, op)
}
func (p *Player) CollisionTo(gridX, gridY int) bool {
	return p.GridX == gridX && p.GridY == gridY
}
func (p *Player) Update(level Level) bool {
	if p.dashState != nil && p.dashState.IsActive() {
		// If currently dashing, process the next step of the dash.
		newX, newY, moved, finished := p.dashState.Update(p.GridX, p.GridY, level, p.CanWalkThroughWalls)
		p.GridX = newX
		p.GridY = newY

		return moved || finished
	}
	return true
}

func (p *Player) GetMoveInput() (dx, dy int, isMoving, isDashing bool) {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
		dx = -1
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
		dx = 1
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		dy = -1
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		dy = 1
	}

	isMoving = dx != 0 || dy != 0
	isDashing = isMoving && (ebiten.IsKeyPressed(ebiten.KeyZ) || ebiten.IsKeyPressed(ebiten.KeyX))
	return
}

func (p *Player) IsDashing() bool {
	return p.dashState != nil && p.dashState.IsActive()
}

func (p *Player) Reset() {
	p.GridX = p.initialGridX
	p.GridY = p.initialGridY
	p.CanDash = p.initialCanDash
	p.CanWalkThroughWalls = p.initialCanWalkThroughWalls
	if p.dashState != nil {
		p.dashState.Reset()
	}
	p.pulseOffset = 0.0
	p.IsActiveTurn = false
}

func (p *Player) CalculateMovePath(level Level, dx, dy int) []image.Point {
	var path []image.Point

	for i := 1; i <= p.Speed; i++ {
		nextX, nextY := p.GridX+dx*i, p.GridY+dy*i

		if p.CanWalkThroughWalls {
			path = append(path, image.Point{X: nextX, Y: nextY})
			continue
		}

		if !level.IsWalkable(nextX, nextY) {
			break
		}
		path = append(path, image.Point{X: nextX, Y: nextY})
	}
	return path
}

// StartDash initiates a dash movement for the player.
func (p *Player) StartDash(level Level, dx, dy int) bool {
	return p.dashState.Start(p.GridX, p.GridY, dx, dy, 10, level, p.CanWalkThroughWalls)
}
