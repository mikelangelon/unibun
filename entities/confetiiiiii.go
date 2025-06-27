package entities

import (
	"image"
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/mikelangelon/unibun/config"
)

const gravity = 0.05

type ConfettiParticle struct {
	x, y, vx, vy float64
	color        color.Color
	life         int
}

type Confetti []*ConfettiParticle

func (c *Confetti) Update() {
	var confettiAlive Confetti
	for _, cp := range *c {
		cp.x += cp.vx
		cp.y += cp.vy
		cp.vy += gravity
		cp.life--
		if cp.life > 0 {
			confettiAlive = append(confettiAlive, cp)
		}
	}
	c = &confettiAlive
}

func (c *Confetti) Draw(screen *ebiten.Image) {
	for _, p := range *c {
		ebitenutil.DrawRect(screen, p.x, p.y, 2, 2, p.color)
	}
}

type MergeAnimation struct {
	IsActive          bool
	timer             int
	duration          int
	topBunStartPos    image.Point
	bottomBunStartPos image.Point
	pattyPos          image.Point
	Confetti          Confetti
}

func (m *MergeAnimation) Update() {
	m.timer--

	var confettiAlive []*ConfettiParticle
	for _, c := range m.Confetti {
		c.x += c.vx
		c.y += c.vy
		c.vy += gravity
		c.life--
		if c.life > 0 {
			confettiAlive = append(confettiAlive, c)
		}
	}
	m.Confetti = confettiAlive
	if m.timer <= 0 {
		m.Deactivate()
	}
}

func (m *MergeAnimation) DrawMergeAnimation(screen *ebiten.Image, patty *BurgerPatty, topBun, bottomBun *Player) {
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

	m.Draw(screen)
}

func (m *MergeAnimation) Draw(screen *ebiten.Image) {
	m.Confetti.Draw(screen)
}

func (m *MergeAnimation) Activate(duration int) {
	m.IsActive = true
	m.duration = 60
	m.timer = m.duration
}
func (m *MergeAnimation) Deactivate() {
	m.IsActive = false
}
func (m *MergeAnimation) ActivateMerge(patty *BurgerPatty, topBun, bottomBun *Player) {
	m.Activate(60)

	m.topBunStartPos = image.Point{X: topBun.GridX, Y: topBun.GridY}
	m.bottomBunStartPos = image.Point{X: bottomBun.GridX, Y: bottomBun.GridY}
	m.pattyPos = image.Point{X: patty.GridX, Y: patty.GridY}

	m.Confetti = CreateConfetti(patty.GridX, patty.GridY)
}

func CreateConfetti(gridX, gridY int) []*ConfettiParticle {
	particles := make([]*ConfettiParticle, 100)
	centerX := float64(gridX*config.TileSize + config.TileSize/2)
	centerY := float64(gridY*config.TileSize + config.TileSize/2)

	for i := range particles {
		angle := rand.Float64() * 2 * math.Pi
		speed := 2 + rand.Float64()*2
		particles[i] = &ConfettiParticle{
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
