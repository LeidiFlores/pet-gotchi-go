package scene

import (
	"image/color"
	"petgotchi/internal/assets"
	"petgotchi/internal/ui"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	etext "github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Menu struct {
	btnStart *ui.Button
	switchTo func(Scene)
	next     Scene
}

func NewMenu(switchTo func(Scene)) *Menu {
	m := &Menu{switchTo: switchTo}
	m.btnStart = &ui.Button{X: 80, Y: 124, W: 160, H: 34, Label: "Start Caring", OnClick: func() {
		m.next = NewPet(m.switchTo)
	}}
	return m
}

func (m *Menu) Update() (Scene, error) {
	pointer := readPointerState()
	m.btnStart.Update(pointer.x, pointer.y, pointer.down, pointer.justPressed, pointer.justReleased)

	if m.next != nil {
		return m.next, nil
	}

	return m, nil
}

func (m *Menu) Draw(screen *ebiten.Image) {
	screen.Fill(ui.Cream)
	ebitenutil.DrawRect(screen, 12, 12, 296, 156, ui.Blush)
	ebitenutil.DrawRect(screen, 18, 18, 284, 144, ui.Mint)
	ebitenutil.DrawRect(screen, 24, 24, 272, 132, ui.Cream)
	ebitenutil.DrawRect(screen, 42, 28, 236, 20, ui.BabyBlue)
	ebitenutil.DrawRect(screen, 50, 40, 220, 78, ui.Blush)
	ebitenutil.DrawRect(screen, 58, 48, 204, 62, ui.Cream)
	drawFrame(screen, 12, 12, 296, 156)
	drawFrame(screen, 58, 48, 204, 62)

	drawCloud(screen, 52, 38, 18)
	drawCloud(screen, 252, 46, 14)
	drawHeart(screen, 72, 44, 6, ui.Peach)
	drawHeart(screen, 248, 34, 5, ui.BabyBlue)
	drawMascotPreview(screen, 160, 81)

	ui.DrawText(screen, "PET GOTCHI GO", 160, 39, ui.TextStyle{
		Size:           18,
		Weight:         ui.FontBold,
		Color:          ui.InkDark,
		ShadowColor:    color.RGBA{255, 255, 255, 200},
		ShadowOffsetX:  0,
		ShadowOffsetY:  2,
		PrimaryAlign:   etext.AlignCenter,
		SecondaryAlign: etext.AlignCenter,
	})
	ui.DrawText(screen, "Adopt a pocket pal,\nthen keep it fed, rested, and happy.", 160, 110, ui.TextStyle{
		Size:           10.5,
		Weight:         ui.FontRegular,
		Color:          ui.InkDark,
		PrimaryAlign:   etext.AlignCenter,
		SecondaryAlign: etext.AlignCenter,
		LineSpacing:    14,
	})
	ui.DrawText(screen, "Tap or click to begin", 160, 112+38, ui.TextStyle{
		Size:           8.5,
		Weight:         ui.FontBold,
		Color:          ui.InkDark,
		PrimaryAlign:   etext.AlignCenter,
		SecondaryAlign: etext.AlignCenter,
	})
	m.btnStart.Draw(screen)
}

func drawFrame(screen *ebiten.Image, x, y, w, h float64) {
	ebitenutil.DrawRect(screen, x, y, w, 1, ui.PanelLine)
	ebitenutil.DrawRect(screen, x, y, 1, h, ui.PanelLine)
	ebitenutil.DrawRect(screen, x+w-1, y, 1, h, ui.InkDark)
	ebitenutil.DrawRect(screen, x, y+h-1, w, 1, ui.InkDark)
}

func drawCloud(screen *ebiten.Image, x, y, size float32) {
	vector.DrawFilledCircle(screen, x-size*0.6, y, size*0.55, ui.Cream, true)
	vector.DrawFilledCircle(screen, x, y-size*0.2, size*0.75, ui.Cream, true)
	vector.DrawFilledCircle(screen, x+size*0.65, y, size*0.5, ui.Cream, true)
	ebitenutil.DrawRect(screen, float64(x-size), float64(y), float64(size*2), float64(size*0.55), ui.Cream)
}

func drawHeart(screen *ebiten.Image, x, y, size float32, fill color.Color) {
	vector.DrawFilledCircle(screen, x-size*0.45, y-size*0.2, size*0.45, fill, true)
	vector.DrawFilledCircle(screen, x+size*0.45, y-size*0.2, size*0.45, fill, true)
	vector.StrokeLine(screen, x-size*0.86, y, x, y+size, size*0.62, fill, true)
	vector.StrokeLine(screen, x+size*0.86, y, x, y+size, size*0.62, fill, true)
}

func drawMascotPreview(screen *ebiten.Image, x, y float32) {
	if assets.PetSprite == nil {
		return
	}
	bounds := assets.PetSprite.Bounds()
	scale := 0.09
	w := float64(bounds.Dx()) * scale
	h := float64(bounds.Dy()) * scale
	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterNearest
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(float64(x)-w/2, float64(y)-h/2)
	screen.DrawImage(assets.PetSprite, op)
}
