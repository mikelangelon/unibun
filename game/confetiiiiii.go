package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/mikelangelon/unibun/config"
	"github.com/mikelangelon/unibun/entities"
	"image"
	"image/color"
	"math"
	"math/rand"
)

const gravity = 0.05

type confettiParticle struct {
	x, y, vx, vy float64
	color        color.Color
	life         int
}

type mergeAnimation struct {
	isActive          bool
	timer             int
	duration          int
	topBunStartPos    image.Point
	bottomBunStartPos image.Point
	pattyPos          image.Point
	confetti          []*confettiParticle
}

func (m *mergeAnimation) Update() {
	m.timer--

	var confettiAlive []*confettiParticle
	for _, c := range m.confetti {
		c.x += c.vx
		c.y += c.vy
		c.vy += gravity
		c.life--
		if c.life > 0 {
			confettiAlive = append(confettiAlive, c)
		}
	}
	m.confetti = confettiAlive
}

func (m *mergeAnimation) draw(screen *ebiten.Image, patty *entities.BurgerPatty, topBun, bottomBun *entities.Player) {
	progress := 1.0 - float64(m.timer)/float64(m.duration)
	if progress > 1.0 {
		progress = 1.0
	}

	topBunPixelX := (float64(m.topBunStartPos.X)*(1-progress) + float64(m.pattyPos.X)*progress) * float64(config.TileSize)
	topBunPixelY := (float64(m.topBunStartPos.Y)*(1-progress) + float64(m.pattyPos.Y)*progress) * float64(config.TileSize)

	bottomBunPixelX := (float64(m.bottomBunStartPos.X)*(1-progress) + float64(m.pattyPos.X)*progress) * float64(config.TileSize)
	bottomBunPixelY := (float64(m.bottomBunStartPos.Y)*(1-progress) + float64(m.pattyPos.Y)*progress) * float64(config.TileSize)

	// Draw the 3 elements unite
	patty.Draw(screen)
	opTop := &ebiten.DrawImageOptions{}
	opTop.GeoM.Translate(topBunPixelX, topBunPixelY)
	screen.DrawImage(topBun.Image, opTop)
	opBottom := &ebiten.DrawImageOptions{}
	opBottom.GeoM.Translate(bottomBunPixelX, bottomBunPixelY)
	screen.DrawImage(bottomBun.Image, opBottom)

	for _, p := range m.confetti {
		ebitenutil.DrawRect(screen, p.x, p.y, 2, 2, p.color)
	}
}

func (m *mergeAnimation) activate(patty *entities.BurgerPatty, topBun, bottomBun *entities.Player) {
	m.isActive = true
	m.duration = 60
	m.timer = m.duration

	m.topBunStartPos = image.Point{X: topBun.GridX, Y: topBun.GridY}
	m.bottomBunStartPos = image.Point{X: bottomBun.GridX, Y: bottomBun.GridY}
	m.pattyPos = image.Point{X: patty.GridX, Y: patty.GridY}

	m.confetti = createConfetti(patty.GridX, patty.GridY)
}

func createConfetti(gridX, gridY int) []*confettiParticle {
	particles := make([]*confettiParticle, 100)
	centerX := float64(gridX*config.TileSize + config.TileSize/2)
	centerY := float64(gridY*config.TileSize + config.TileSize/2)

	for i := range particles {
		angle := rand.Float64() * 2 * math.Pi
		speed := 2 + rand.Float64()*2
		particles[i] = &confettiParticle{
			x:     centerX,
			y:     centerY,
			vx:    math.Cos(angle) * speed,
			vy:    math.Sin(angle) * speed,
			life:  30 + rand.Intn(30),
			color: color.RGBA{R: uint8(rand.Intn(256)), G: uint8(rand.Intn(256)), B: uint8(rand.Intn(256)), A: 0xff},
		}
	}
	return particles
}
