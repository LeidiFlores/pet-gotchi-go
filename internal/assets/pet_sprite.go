package assets

import (
	"bytes"
	_ "embed"
	"image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed perrito-pixel.png
var petSpritePNG []byte

var PetSprite *ebiten.Image

func init() {
	img, err := png.Decode(bytes.NewReader(petSpritePNG))
	if err != nil {
		log.Fatal(err)
	}
	PetSprite = ebiten.NewImageFromImage(img)
}
