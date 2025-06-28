package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/mikelangelon/unibun/assets"
	"github.com/mikelangelon/unibun/common"
	"github.com/mikelangelon/unibun/config"
)

func drawTutorial(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 20, G: 20, B: 20, A: 255}) // Dark background

	yPos := 40
	drawText := func(s string, x, y int) {
		ebitenutil.DebugPrintAt(screen, s, x, y)
	}
	drawImage := func(img *ebiten.Image, x, y float64) {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(x, y)
		screen.DrawImage(img, op)
	}

	// Title
	title := "Bun-structions"
	titleX := (config.WindowWidth - len(title)*6) / 2
	drawText(title, titleX, yPos)
	yPos += 40

	// Buns
	topBunImg := common.GetImage(assets.TopBun)
	bottomBunImg := common.GetImage(assets.BottomBun)
	drawImage(topBunImg, 50, float64(yPos))
	drawImage(bottomBunImg, 50+32+10, float64(yPos))
	drawText("You control these Buns.", 50+32*2+20, yPos+8)
	yPos += 50

	// Patty
	pattyImg := common.GetImage(assets.BurgerPatty)
	drawImage(pattyImg, 50, float64(yPos))
	drawText("Your goal is to unite the Buns with the Patty.", 50+32+10, yPos+8)
	yPos += 50

	// Burger Combination Guide
	drawText("Form the Burger like this:", 50, yPos)
	yPos += 25
	drawImage(topBunImg, 50, float64(yPos))
	drawImage(pattyImg, 50+32, float64(yPos))
	drawImage(bottomBunImg, 50+32*2, float64(yPos))
	drawImage(bottomBunImg, 200, float64(yPos))
	drawImage(pattyImg, 200+32, float64(yPos))
	drawImage(topBunImg, 200+32*2, float64(yPos))
	drawImage(topBunImg, 350, float64(yPos))
	drawImage(pattyImg, 350, float64(yPos+32))
	drawImage(bottomBunImg, 350, float64(yPos+32*2))
	yPos += 120

	drawText("You also control:", 50, yPos)
	yPos += 25
	lettuceImg := common.GetImage(assets.Lettuce)
	drawImage(lettuceImg, 50, float64(yPos))
	drawText("Lettuce: unite it with a Bun to move across walls.", 50+32+10, yPos+8)
	yPos += 40
	cheeseImg := common.GetImage(assets.Cheese)
	drawImage(cheeseImg, 50, float64(yPos))
	drawText("Cheese: unite it with a Bun to Dash. Press Z/X and an arrow for it.", 50+32+10, yPos+8)
	yPos += 50

	// Winning
	clientImg := common.GetImage(assets.Client)
	drawImage(clientImg, 50, float64(yPos))
	drawText("When the Burger is completed, go to the client!", 50+32+10, yPos+8)
	yPos += 50

	// Controls
	drawText("Controls:", 50, yPos)
	yPos += 20
	drawText("  - Arrows/WASD --> Move", 50, yPos)
	yPos += 20
	drawText("  - Z/X + Arrows/WASD --> Dash (if you have cheese)", 50, yPos)
	yPos += 20
	drawText("  - Space --> Pause", 50, yPos)
	yPos += 20
	drawText("  - Enter --> Select options & to leave this tutorial", 50, yPos)
	yPos += 40
}
