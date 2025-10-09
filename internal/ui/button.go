package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Button struct {
	X, Y int
	W, H int
	Label string
	OnClick func()
	Hover bool
}

func (b *Button) Update(mx, my int, justPressed bool) {
	b.Hover = mx >= b.X && mx <= b.X+b.W && my >= b.Y && my <= b.Y+b.H
	if b.Hover && justPressed && b.OnClick != nil {
		b.OnClick()
	}
}

func (b *Button) Draw(dst *ebiten.Image) {
	col := color.RGBA{255, 214, 231, 255} // pastel
	if b.Hover { col = color.RGBA{255, 196, 224, 255} }
	ebitenutil.DrawRect(dst, float64(b.X), float64(b.Y), float64(b.W), float64(b.H), col)
	// Texto simple (placeholder): usar ebiten/text para tipografía real
	ebitenutil.DebugPrintAt(dst, b.Label, b.X+10, b.Y+10)
}
