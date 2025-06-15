package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/mikelangelon/unibun/config"
	"golang.org/x/image/font/basicfont"
	"image/color"
)

func drawWinning(screen *ebiten.Image) {
	winTextStr := "YOU WIN!"
	textColor := color.White

	fontFace := basicfont.Face7x13
	textBounds := text.BoundString(fontFace, winTextStr)
	textW := textBounds.Dx()
	textH := textBounds.Dy()

	tempTextImg := ebiten.NewImage(textW, textH)
	text.Draw(tempTextImg, winTextStr, fontFace, -textBounds.Min.X, -textBounds.Min.Y, textColor)

	scaleFactor := 4.0
	opText := &ebiten.DrawImageOptions{}
	opText.GeoM.Translate(-float64(textW)/2, -float64(textH)/2)
	opText.GeoM.Scale(scaleFactor, scaleFactor)
	opText.GeoM.Translate(config.WindowWidth/2, config.WindowHeight/2)

	screen.DrawImage(tempTextImg, opText)
	messageText := "Nothing else is implemented after this :D"
	ebitenutil.DebugPrintAt(screen, messageText, config.WindowWidth/2-(len(messageText)*7/2), config.WindowHeight/2+int(float64(textH)*scaleFactor/2)+20)
}
