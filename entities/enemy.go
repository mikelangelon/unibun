package entities

import (
	"math"
	"math/rand/v2"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mikelangelon/unibun/config"
)

type animationMode int

const (
	AnimateOnMove animationMode = iota
	AnimateContinuously

	animationFrameDuration = int(200 * time.Millisecond / (time.Second / 60))
	moveAnimationDuration  = 30
)

// TODO Rename this monstruosity name
type Enemier interface {
	Draw(screen *ebiten.Image)
	Update(level Level) bool
	Collision(player *Player) bool
	Image() *ebiten.Image
	Icon() *ebiten.Image
	Reset()
	Position() (int, int)
}

type Enemy struct {
	gridX, gridY               int
	initialGridX, initialGridY int
	images                     []*ebiten.Image
	facingDirection            int // 1 for right, -1 for left

	animationTimer     int
	currentFrameIndex  int
	animationDirection int
	animationMode      animationMode
	animationTrigger   int
	pulseOffset        float64
}

func NewEnemy(startX, startY int, img []*ebiten.Image) *Enemy {
	return &Enemy{
		gridX:              startX,
		gridY:              startY,
		initialGridX:       startX,
		initialGridY:       startY,
		images:             img,
		facingDirection:    1,
		animationTimer:     animationFrameDuration,
		currentFrameIndex:  0,
		animationDirection: 1,
		animationMode:      AnimateOnMove,
		animationTrigger:   0,
		pulseOffset:        rand.Float64() * 2 * math.Pi,
	}
}
func (e *Enemy) Collision(player *Player) bool {
	return e.gridX == player.GridX && e.gridY == player.GridY
}

func (e *Enemy) Update(level Level) bool {
	oldX, oldY := e.gridX, e.gridY
	newX, newY := nextRandomMove(level, e)
	return updatePosDirection(e, oldX, oldY, newX, newY)
}

func (e *Enemy) Draw(screen *ebiten.Image) {
	if len(e.images) > 1 {
		shouldAnimate := false
		switch e.animationMode {
		case AnimateContinuously:
			shouldAnimate = true
		case AnimateOnMove:
			if e.animationTrigger > 0 {
				shouldAnimate = true
				e.animationTrigger--
			}
		}
		// TODO Improve this if-statement madness
		if shouldAnimate {
			e.animationTimer--
			if e.animationTimer <= 0 {
				e.animationTimer = animationFrameDuration
				e.currentFrameIndex += e.animationDirection
				if e.currentFrameIndex <= 0 || e.currentFrameIndex >= len(e.images)-1 {
					e.animationDirection *= -1
				}
			}
		}
	}

	op := &ebiten.DrawImageOptions{}
	w, h := e.images[e.currentFrameIndex].Bounds().Dx(), e.images[e.currentFrameIndex].Bounds().Dy()
	centerX, centerY := float64(w)/2, float64(h)/2

	op.GeoM.Translate(-centerX, -centerY)

	if e.animationMode == AnimateOnMove && e.animationTrigger <= 0 {
		scaleY := 1.0 + 0.05*math.Sin(e.pulseOffset/20.0)
		op.GeoM.Scale(1, scaleY)
		translateY := -(float64(h)*scaleY - float64(h)) / 2
		op.GeoM.Translate(0, translateY)
	}
	e.pulseOffset++

	if e.facingDirection == -1 {
		op.GeoM.Scale(-1, 1)
	}

	op.GeoM.Translate(centerX, centerY)
	pixelX := float64(e.gridX * config.TileSize)
	pixelY := float64(e.gridY * config.TileSize)
	op.GeoM.Translate(pixelX, pixelY)
	screen.DrawImage(e.images[e.currentFrameIndex], op)
}

func (e *Enemy) Image() *ebiten.Image {
	return e.images[e.currentFrameIndex]
}

func (e *Enemy) Icon() *ebiten.Image {
	return e.images[0]
}
func (e *Enemy) Reset() {
	e.gridX = e.initialGridX
	e.gridY = e.initialGridY
	e.facingDirection = 1
	e.animationTimer = animationFrameDuration
	e.currentFrameIndex = 0
	e.animationDirection = 1
}

func (e *Enemy) Position() (int, int) {
	return e.gridX, e.gridY
}

func (e *Enemy) TriggerMoveAnimation() {
	if e.animationMode == AnimateOnMove {
		e.animationTrigger = moveAnimationDuration
	}
}

func updatePosDirection(e *Enemy, oldX, oldY, newX, newY int) bool {
	if newX > oldX {
		e.facingDirection = 1
	} else if newX < oldX {
		e.facingDirection = -1
	}
	e.gridX, e.gridY = newX, newY
	if oldX != newX || oldY != newY {
		e.TriggerMoveAnimation()
	}
	return true
}
