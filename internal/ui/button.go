package ui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	etext "github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Button struct {
	X, Y         int
	W, H         int
	Label        string
	OnClick      func()
	BaseColor    color.Color
	HoverColor   color.Color
	PressedColor color.Color
	Hover        bool
	Pressed      bool
}

func (b *Button) Update(mx, my int, pointerDown, justPressed, justReleased bool) {
	b.Hover = mx >= b.X && mx <= b.X+b.W && my >= b.Y && my <= b.Y+b.H
	if (pointerDown || justPressed) && b.Hover {
		b.Pressed = true
	}
	if justReleased {
		if b.Pressed && b.Hover && b.OnClick != nil {
			b.OnClick()
		}
		b.Pressed = false
	}
	if !pointerDown && !justReleased {
		b.Pressed = false
	}
}

func (b *Button) Draw(dst *ebiten.Image) {
	col := fallbackColor(b.BaseColor, Blush)
	if b.Hover {
		col = fallbackColor(b.HoverColor, Peach)
	}
	if b.Pressed {
		col = fallbackColor(b.PressedColor, BlushStrong)
	}
	shadow := fallbackColor(b.PressedColor, BlushStrong)
	if b.Pressed {
		shadow = fallbackColor(b.HoverColor, Peach)
	}

	y := float64(b.Y)
	if b.Pressed {
		y += 2
	}

	ebitenutil.DrawRect(dst, float64(b.X), y+3, float64(b.W), float64(b.H-1), shadow)
	ebitenutil.DrawRect(dst, float64(b.X), y, float64(b.W), float64(b.H-3), col)
	ebitenutil.DrawRect(dst, float64(b.X+1), y+1, float64(b.W-2), 1, PanelLine)
	ebitenutil.DrawRect(dst, float64(b.X+1), y+1, 1, float64(b.H-5), PanelLine)
	ebitenutil.DrawRect(dst, float64(b.X+b.W-2), y+1, 1, float64(b.H-5), InkDark)
	ebitenutil.DrawRect(dst, float64(b.X+1), y+float64(b.H-4), float64(b.W-2), 1, InkDark)
	ebitenutil.DrawRect(dst, float64(b.X+3), y+3, float64(b.W-6), float64(b.H-9), Cream)

	labelY := y + float64(b.H-3)/2
	if b.Pressed {
		labelY++
	}
	DrawText(dst, b.Label, float64(b.X+b.W/2), labelY, TextStyle{
		Size:           10,
		Weight:         FontBold,
		Color:          InkDark,
		ShadowColor:    color.RGBA{255, 255, 255, 180},
		ShadowOffsetX:  0,
		ShadowOffsetY:  1,
		PrimaryAlign:   etext.AlignCenter,
		SecondaryAlign: etext.AlignCenter,
	})
}

func fallbackColor(raw color.Color, fallback color.Color) color.Color {
	if raw == nil {
		return fallback
	}
	return raw
}
