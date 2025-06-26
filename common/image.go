package common

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"log/slog"
)

// GetImage returns an image for the provided bytes, or an empty one in case of errors
func GetImage(b []byte) *ebiten.Image {
	playerDecoded, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		slog.Error("unexpected error decoding player image", "error", err)
		return ebiten.NewImage(32, 32)
	}
	img := ebiten.NewImageFromImage(playerDecoded)
	return img
}
