package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand/v2"
)

const (
	shakeDefaultDuration  = 30 //0.5 seconds at 60FPS
	shakeDefaultMagnitude = 4.0
)

type Shake struct {
	shakeTimer     int     // Timer to define when the shake is completed
	shakeMagnitude float64 // Max pixel offset for the current shake
	shakeOffsetX   float64 // Current random X offset for this frame
	shakeOffsetY   float64 // Current random Y offset for this frame
}

func newShake(durationFrames int, magnitude float64) *Shake {
	return &Shake{
		shakeTimer:     durationFrames,
		shakeMagnitude: magnitude,
		shakeOffsetX:   0,
		shakeOffsetY:   0,
	}
}

func (s *Shake) Update() {
	s.shakeTimer--
	if s.shakeTimer <= 0 {
		// it nils itself and leaves
		s = nil
		return
	} else {
		// Generate random offsets
		s.shakeOffsetX = (rand.Float64()*2 - 1) * s.shakeMagnitude
		s.shakeOffsetY = (rand.Float64()*2 - 1) * s.shakeMagnitude
	}
}

func (s Shake) shake(op *ebiten.DrawImageOptions) {
	op.GeoM.Translate(s.shakeOffsetX, s.shakeOffsetY)
}
