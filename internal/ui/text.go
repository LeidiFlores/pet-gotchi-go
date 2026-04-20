package ui

import (
	"bytes"
	"image/color"
	"log"

	"golang.org/x/image/font/gofont/gobold"
	"golang.org/x/image/font/gofont/goregular"

	"github.com/hajimehoshi/ebiten/v2"
	etext "github.com/hajimehoshi/ebiten/v2/text/v2"
)

type FontWeight string

const (
	FontRegular FontWeight = "regular"
	FontBold    FontWeight = "bold"
)

type TextStyle struct {
	Size           float64
	Weight         FontWeight
	Color          color.Color
	ShadowColor    color.Color
	ShadowOffsetX  float64
	ShadowOffsetY  float64
	PrimaryAlign   etext.Align
	SecondaryAlign etext.Align
	LineSpacing    float64
}

var (
	regularFaceSource *etext.GoTextFaceSource
	boldFaceSource    *etext.GoTextFaceSource
)

func init() {
	var err error
	regularFaceSource, err = etext.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		log.Fatal(err)
	}
	boldFaceSource, err = etext.NewGoTextFaceSource(bytes.NewReader(gobold.TTF))
	if err != nil {
		log.Fatal(err)
	}
}

func DrawText(dst *ebiten.Image, value string, x, y float64, style TextStyle) {
	if value == "" {
		return
	}

	face := &etext.GoTextFace{
		Source: faceSourceFor(style.Weight),
		Size:   fontSize(style.Size),
	}
	op := &etext.DrawOptions{}
	op.GeoM.Translate(x, y)
	op.LineSpacing = lineSpacing(style.LineSpacing, style.Size)
	op.PrimaryAlign = style.PrimaryAlign
	op.SecondaryAlign = style.SecondaryAlign

	if style.ShadowColor != nil {
		shadow := *op
		shadow.GeoM.Translate(style.ShadowOffsetX, style.ShadowOffsetY)
		shadow.ColorScale.ScaleWithColor(style.ShadowColor)
		etext.Draw(dst, value, face, &shadow)
	}

	op.ColorScale.ScaleWithColor(fallbackColor(style.Color, InkSoft))
	etext.Draw(dst, value, face, op)
}

func MeasureText(value string, style TextStyle) (width, height float64) {
	if value == "" {
		return 0, 0
	}
	face := &etext.GoTextFace{
		Source: faceSourceFor(style.Weight),
		Size:   fontSize(style.Size),
	}
	return etext.Measure(value, face, lineSpacing(style.LineSpacing, style.Size))
}

func faceSourceFor(weight FontWeight) *etext.GoTextFaceSource {
	if weight == FontBold {
		return boldFaceSource
	}
	return regularFaceSource
}

func fontSize(size float64) float64 {
	if size <= 0 {
		return 12
	}
	return size
}

func lineSpacing(explicit, size float64) float64 {
	if explicit > 0 {
		return explicit
	}
	return fontSize(size) * 1.2
}
