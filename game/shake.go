package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand/v2"
)

const (
	shakeDefaultDuration  = 30
	shakeDefaultMagnitude = 4.0
)

// shake contains the params for the shake effect
type shake struct {
	timer     int
	magnitude float64
	offsetX   float64
	offsetY   float64
}

func newShake(durationFrames int, magnitude float64) *shake {
	return &shake{
		timer:     durationFrames,
		magnitude: magnitude,
		offsetX:   0,
		offsetY:   0,
	}
}

func (s *shake) Update() {
	s.timer--
	if s.timer <= 0 {
		// it nils itself and leaves
		s = nil
		return
	} else {
		// Generate random offsets
		s.offsetX = (rand.Float64()*2 - 1) * s.magnitude
		s.offsetY = (rand.Float64()*2 - 1) * s.magnitude
	}
}

func (s shake) shake(op *ebiten.DrawImageOptions) {
	op.GeoM.Translate(s.offsetX, s.offsetY)
}
