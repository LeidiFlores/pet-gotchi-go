package scene

import (
	"fmt"
	"image/color"
	"math"

	"petgotchi/internal/assets"
	"petgotchi/internal/ui"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	etext "github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Pet struct {
	SwitchTo func(Scene)
	btnFeed  *ui.Button
	btnBrush *ui.Button
	btnSleep *ui.Button
	// stats
	Hunger float64 // 0..100 (0 = full, 100 = hungry)
	Sleep  float64 // 0..100 (0 = awake, 100 = sleepy)
	Mood   float64 // 0..100 (0 = happy, 100 = sad)
	T      int

	feedbackKind  string
	feedbackTicks int
}

func NewPet(switchTo func(Scene)) *Pet {
	p := &Pet{SwitchTo: switchTo, Hunger: 50, Sleep: 50, Mood: 20}
	p.btnFeed = &ui.Button{
		X:            26,
		Y:            178,
		W:            82,
		H:            28,
		Label:        "Feed [F]",
		BaseColor:    ui.Peach,
		HoverColor:   ui.Blush,
		PressedColor: ui.BlushStrong,
		OnClick:      p.feed,
	}
	p.btnBrush = &ui.Button{
		X:            119,
		Y:            178,
		W:            82,
		H:            28,
		Label:        "Brush [B]",
		BaseColor:    ui.Blush,
		HoverColor:   ui.Peach,
		PressedColor: ui.BlushStrong,
		OnClick:      p.brush,
	}
	p.btnSleep = &ui.Button{
		X:            212,
		Y:            178,
		W:            82,
		H:            28,
		Label:        "Nap [N]",
		BaseColor:    ui.BabyBlue,
		HoverColor:   ui.Mint,
		PressedColor: ui.BlushStrong,
		OnClick:      p.rest,
	}
	return p
}

func (p *Pet) Update() (Scene, error) {
	p.T++
	if p.T%60 == 0 {
		p.Hunger = Clamp01(p.Hunger + 1)
		p.Sleep = Clamp01(p.Sleep + 1)
		p.Mood = Clamp01(p.Mood + 0.5)
	}
	if p.feedbackTicks > 0 {
		p.feedbackTicks--
	}

	pointer := readPointerState()
	p.btnFeed.Update(pointer.x, pointer.y, pointer.down, pointer.justPressed, pointer.justReleased)
	p.btnBrush.Update(pointer.x, pointer.y, pointer.down, pointer.justPressed, pointer.justReleased)
	p.btnSleep.Update(pointer.x, pointer.y, pointer.down, pointer.justPressed, pointer.justReleased)

	p.UpdateInputs(
		inpututil.IsKeyJustPressed(ebiten.KeyF),
		inpututil.IsKeyJustPressed(ebiten.KeyN),
		inpututil.IsKeyJustPressed(ebiten.KeyB),
	)
	return p, nil
}

func (p *Pet) UpdateInputs(feed, sleep, brush bool) {
	if feed {
		p.feed()
	}
	if sleep {
		p.rest()
	}
	if brush {
		p.brush()
	}
}

func (p *Pet) Draw(screen *ebiten.Image) {
	screen.Fill(ui.Cream)
	ebitenutil.DrawRect(screen, 10, 10, 300, 220, ui.BabyBlue)
	ebitenutil.DrawRect(screen, 18, 18, 284, 204, ui.Mint)
	ebitenutil.DrawRect(screen, 26, 26, 268, 188, ui.Cream)
	ebitenutil.DrawRect(screen, 34, 34, 74, 38, ui.Peach)
	ebitenutil.DrawRect(screen, 123, 34, 74, 38, ui.Blush)
	ebitenutil.DrawRect(screen, 212, 34, 74, 38, ui.BabyBlue)
	ebitenutil.DrawRect(screen, 26, 170, 268, 44, ui.Blush)
	drawFrame(screen, 10, 10, 300, 220)
	drawFrame(screen, 26, 170, 268, 44)
	drawFrame(screen, 34, 34, 74, 38)
	drawFrame(screen, 123, 34, 74, 38)
	drawFrame(screen, 212, 34, 74, 38)

	drawPetBuddy(screen, p)
	ui.DrawText(screen, "Care for your pocket pal", 160, 18, ui.TextStyle{
		Size:           12,
		Weight:         ui.FontBold,
		Color:          ui.InkDark,
		ShadowColor:    color.RGBA{255, 255, 255, 200},
		ShadowOffsetX:  0,
		ShadowOffsetY:  1,
		PrimaryAlign:   etext.AlignCenter,
		SecondaryAlign: etext.AlignStart,
	})
	ui.DrawText(screen, "Tap a button or press its key.", 160, 159, ui.TextStyle{
		Size:           7.5,
		Weight:         ui.FontRegular,
		Color:          ui.InkDark,
		PrimaryAlign:   etext.AlignCenter,
		SecondaryAlign: etext.AlignStart,
	})

	drawStatMeter(screen, 34, 34, 74, 38, "Hunger", p.Hunger, ui.Peach, "drumstick")
	drawStatMeter(screen, 123, 34, 74, 38, "Mood", p.Mood, ui.Blush, "heart")
	drawStatMeter(screen, 212, 34, 74, 38, "Sleep", p.Sleep, ui.BabyBlue, "moon")

	p.btnFeed.Draw(screen)
	p.btnBrush.Draw(screen)
	p.btnSleep.Draw(screen)
}

func (p *Pet) feed() {
	p.Hunger = Clamp01(p.Hunger - 2)
	p.Mood = Clamp01(p.Mood - 0.5)
	p.startFeedback("feed")
}

func (p *Pet) rest() {
	p.Sleep = Clamp01(p.Sleep - 2)
	p.startFeedback("sleep")
}

func (p *Pet) brush() {
	p.Mood = Clamp01(p.Mood - 2)
	p.startFeedback("brush")
}

func (p *Pet) startFeedback(kind string) {
	p.feedbackKind = kind
	p.feedbackTicks = 28
}

func Clamp01(v float64) float64 {
	if v < 0 {
		return 0
	}
	if v > 100 {
		return 100
	}
	return v
}

func drawStatMeter(screen *ebiten.Image, x, y, w, h float64, label string, rawValue float64, panelColor color.Color, icon string) {
	ebitenutil.DrawRect(screen, x, y, w, h, panelColor)

	fillValue := 100 - Clamp01(rawValue)
	fillWidth := math.Round((w - 12) * (fillValue / 100))
	fillColor := statFillColor(fillValue)

	ebitenutil.DrawRect(screen, x+6, y+18, w-12, 8, ui.Cream)
	if fillWidth > 0 {
		ebitenutil.DrawRect(screen, x+6, y+18, fillWidth, 8, fillColor)
	}

	switch icon {
	case "heart":
		drawMeterHeart(screen, float32(x+12), float32(y+10), 4, fillColor)
	case "moon":
		drawMoon(screen, x+12, y+10, 5, fillColor)
	default:
		drawDrumstick(screen, x+12, y+10, fillColor)
	}

	ui.DrawText(screen, label, x+22, y+6, ui.TextStyle{
		Size:           8,
		Weight:         ui.FontBold,
		Color:          ui.InkDark,
		PrimaryAlign:   etext.AlignStart,
		SecondaryAlign: etext.AlignStart,
	})
	ui.DrawText(screen, statText(fillValue), x+w/2, y+29, ui.TextStyle{
		Size:           7.5,
		Weight:         ui.FontBold,
		Color:          ui.InkDark,
		PrimaryAlign:   etext.AlignCenter,
		SecondaryAlign: etext.AlignCenter,
	})
}

func statFillColor(value float64) color.Color {
	if value <= 25 {
		return ui.Alert
	}
	if value <= 60 {
		return ui.Peach
	}
	return ui.Success
}

func statText(value float64) string {
	switch {
	case value <= 15:
		return "CHECK NOW"
	case value <= 35:
		return "LOW"
	case value <= 65:
		return fmt.Sprintf("%d%%", int(math.Round(value)))
	default:
		return "GOOD"
	}
}

func drawMeterHeart(screen *ebiten.Image, x, y float32, size float32, fill color.Color) {
	vector.DrawFilledCircle(screen, x-size*0.45, y-size*0.2, size*0.45, fill, true)
	vector.DrawFilledCircle(screen, x+size*0.45, y-size*0.2, size*0.45, fill, true)
	vector.StrokeLine(screen, x-size*0.86, y, x, y+size, size*0.62, fill, true)
	vector.StrokeLine(screen, x+size*0.86, y, x, y+size, size*0.62, fill, true)
}

func drawMoon(screen *ebiten.Image, x, y, r float64, fill color.Color) {
	vector.DrawFilledCircle(screen, float32(x), float32(y), float32(r), fill, true)
	vector.DrawFilledCircle(screen, float32(x+2), float32(y-1), float32(r-1), ui.Cream, true)
}

func drawDrumstick(screen *ebiten.Image, x, y float64, fill color.Color) {
	vector.DrawFilledCircle(screen, float32(x+2), float32(y), 4.2, fill, true)
	ebitenutil.DrawRect(screen, x+4, y-1.3, 5, 2.6, fill)
	vector.DrawFilledCircle(screen, float32(x+10), float32(y-2), 1.6, ui.Cream, true)
	vector.DrawFilledCircle(screen, float32(x+10), float32(y+2), 1.6, ui.Cream, true)
}

func isBlink(t int) bool {
	phase := t % 200
	return phase >= 196
}

func drawPetTail(screen *ebiten.Image, x, y float32) {
	vector.DrawFilledCircle(screen, x, y, 7, ui.BlushStrong, true)
	vector.DrawFilledCircle(screen, x-1, y-3.5, 5.5, ui.Blush, true)
	vector.DrawFilledCircle(screen, x, y-6.5, 3, ui.Cream, true)
}

func drawPetBuddy(screen *ebiten.Image, p *Pet) {
	x := float32(160)
	bounce := math.Sin(float64(p.T)/11.0) * 1.8
	if p.feedbackTicks > 0 {
		progress := float64(p.feedbackTicks) / 28.0
		bounce -= math.Sin(progress*math.Pi) * 5
	}
	y := float32(121 + bounce)
	face := petFaceState(p.Hunger, p.Sleep, p.Mood)

	drawPetShadow(screen, x, y+29)
	drawPetSprite(screen, x, y)
	drawPetStateBadge(screen, x+50, y-26, face)
	drawPetFeedback(screen, x, y-32, p.feedbackKind, p.feedbackTicks)
}

func drawPetSprite(screen *ebiten.Image, x, y float32) {
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
	op.GeoM.Translate(float64(x)-w/2, float64(y)-h/2+2)
	screen.DrawImage(assets.PetSprite, op)
}

func drawPetStateBadge(screen *ebiten.Image, x, y float32, face string) {
	vector.DrawFilledCircle(screen, x, y, 10, ui.Cream, true)
	vector.DrawFilledCircle(screen, x, y, 8.8, ui.BabyBlue, true)

	switch face {
	case "sleepy":
		drawMoon(screen, float64(x-4), float64(y), 4, ui.InkDark)
		drawSleepMarks(screen, x+1, y-5)
	case "hungry":
		drawDrumstick(screen, float64(x-5), float64(y), ui.InkDark)
	case "gloomy":
		vector.StrokeLine(screen, x-4, y+2, x, y-1, 1.2, ui.InkDark, true)
		vector.StrokeLine(screen, x, y-1, x+4, y+2, 1.2, ui.InkDark, true)
	case "delighted":
		drawHeart(screen, x, y, 4, ui.BlushStrong)
	default:
		drawHeart(screen, x, y, 3.8, ui.Success)
	}
}

func petFaceState(hunger, sleep, mood float64) string {
	switch {
	case sleep >= 75:
		return "sleepy"
	case mood >= 72:
		return "gloomy"
	case hunger >= 72:
		return "hungry"
	case mood <= 30 && hunger <= 45 && sleep <= 45:
		return "delighted"
	default:
		return "content"
	}
}

func drawPetShadow(screen *ebiten.Image, x, y float32) {
	vector.DrawFilledCircle(screen, x, y, 20, ui.Blush, true)
	vector.DrawFilledCircle(screen, x, y, 14, ui.Cream, true)
}

func drawPetEar(screen *ebiten.Image, x, y, size float32, outer color.Color, inner color.Color) {
	vector.DrawFilledCircle(screen, x, y, size+1.5, ui.BlushStrong, true)
	vector.DrawFilledCircle(screen, x, y, size, outer, true)
	vector.DrawFilledCircle(screen, x, y, size*0.45, inner, true)
}

func drawPetBody(screen *ebiten.Image, x, y, bodyRadius float32) {
	vector.DrawFilledCircle(screen, x, y+3, bodyRadius+5, ui.BlushStrong, true)
	vector.DrawFilledCircle(screen, x, y, bodyRadius, ui.Blush, true)
	vector.DrawFilledCircle(screen, x+2, y+6, 11, ui.Cream, true)
	vector.DrawFilledCircle(screen, x-7, y-10, 5, color.RGBA{255, 255, 255, 90}, true)
}

func drawPetPaws(screen *ebiten.Image, x, y float32) {
	vector.DrawFilledCircle(screen, x-10, y, 5.5, ui.Peach, true)
	vector.DrawFilledCircle(screen, x-13, y-0.5, 1.8, ui.BlushStrong, true)
	vector.DrawFilledCircle(screen, x-10, y-2.5, 1.8, ui.BlushStrong, true)
	vector.DrawFilledCircle(screen, x-7, y-0.5, 1.8, ui.BlushStrong, true)
	vector.DrawFilledCircle(screen, x+10, y, 5.5, ui.Peach, true)
	vector.DrawFilledCircle(screen, x+7, y-0.5, 1.8, ui.BlushStrong, true)
	vector.DrawFilledCircle(screen, x+10, y-2.5, 1.8, ui.BlushStrong, true)
	vector.DrawFilledCircle(screen, x+13, y-0.5, 1.8, ui.BlushStrong, true)
}

func drawPetFace(screen *ebiten.Image, x, y float32, face string, blink bool) {
	leftEyeX := x - 8
	rightEyeX := x + 8
	eyeY := y - 5
	browY := eyeY - 5

	switch face {
	case "gloomy":
		vector.StrokeLine(screen, leftEyeX-3, browY+1.5, leftEyeX+3, browY-1.5, 1.5, ui.InkSoft, true)
		vector.StrokeLine(screen, rightEyeX-3, browY-1.5, rightEyeX+3, browY+1.5, 1.5, ui.InkSoft, true)
	case "delighted":
		vector.StrokeLine(screen, leftEyeX-3, browY-1, leftEyeX+3, browY-2, 1.3, ui.InkSoft, true)
		vector.StrokeLine(screen, rightEyeX-3, browY-2, rightEyeX+3, browY-1, 1.3, ui.InkSoft, true)
	default:
		vector.StrokeLine(screen, leftEyeX-3, browY, leftEyeX+3, browY, 1.3, ui.InkSoft, true)
		vector.StrokeLine(screen, rightEyeX-3, browY, rightEyeX+3, browY, 1.3, ui.InkSoft, true)
	}

	switch face {
	case "sleepy":
		vector.StrokeLine(screen, leftEyeX-3, eyeY, leftEyeX+3, eyeY, 1.8, ui.InkSoft, true)
		vector.StrokeLine(screen, rightEyeX-3, eyeY, rightEyeX+3, eyeY, 1.8, ui.InkSoft, true)
		drawSleepMarks(screen, x+20, y-20)
	case "gloomy":
		vector.StrokeLine(screen, leftEyeX-3, eyeY+1, leftEyeX+3, eyeY-1, 1.8, ui.InkSoft, true)
		vector.StrokeLine(screen, rightEyeX-3, eyeY-1, rightEyeX+3, eyeY+1, 1.8, ui.InkSoft, true)
	case "hungry":
		vector.DrawFilledCircle(screen, leftEyeX, eyeY, 2.8, ui.InkSoft, true)
		vector.DrawFilledCircle(screen, rightEyeX, eyeY, 2.8, ui.InkSoft, true)
		vector.DrawFilledCircle(screen, leftEyeX+0.8, eyeY-0.8, 1.0, color.RGBA{255, 255, 255, 200}, true)
		vector.DrawFilledCircle(screen, rightEyeX+0.8, eyeY-0.8, 1.0, color.RGBA{255, 255, 255, 200}, true)
	default:
		if blink {
			vector.StrokeLine(screen, leftEyeX-3, eyeY, leftEyeX+3, eyeY, 1.8, ui.InkSoft, true)
			vector.StrokeLine(screen, rightEyeX-3, eyeY, rightEyeX+3, eyeY, 1.8, ui.InkSoft, true)
		} else {
			vector.DrawFilledCircle(screen, leftEyeX, eyeY, 2.4, ui.InkSoft, true)
			vector.DrawFilledCircle(screen, rightEyeX, eyeY, 2.4, ui.InkSoft, true)
			vector.DrawFilledCircle(screen, leftEyeX+1, eyeY-1, 0.8, color.RGBA{255, 255, 255, 220}, true)
			vector.DrawFilledCircle(screen, rightEyeX+1, eyeY-1, 0.8, color.RGBA{255, 255, 255, 220}, true)
		}
	}

	vector.DrawFilledCircle(screen, x-14, y+3, 4.5, color.RGBA{247, 187, 206, 150}, true)
	vector.DrawFilledCircle(screen, x+14, y+3, 4.5, color.RGBA{247, 187, 206, 150}, true)
	vector.DrawFilledCircle(screen, x, y+2, 1.8, ui.Peach, true)

	switch face {
	case "sleepy":
		vector.StrokeLine(screen, x-4, y+11, x+4, y+11, 1.5, ui.InkSoft, true)
	case "gloomy":
		vector.StrokeLine(screen, x-5, y+13, x, y+10, 1.5, ui.InkSoft, true)
		vector.StrokeLine(screen, x, y+10, x+5, y+13, 1.5, ui.InkSoft, true)
	case "hungry":
		vector.DrawFilledCircle(screen, x, y+11, 3.5, ui.Peach, true)
		vector.DrawFilledCircle(screen, x, y+9.8, 2.5, ui.Cream, true)
	case "delighted":
		vector.StrokeLine(screen, x-6, y+9, x, y+14, 1.8, ui.InkSoft, true)
		vector.StrokeLine(screen, x, y+14, x+6, y+9, 1.8, ui.InkSoft, true)
	default:
		vector.StrokeLine(screen, x-5, y+10, x, y+12, 1.5, ui.InkSoft, true)
		vector.StrokeLine(screen, x, y+12, x+5, y+10, 1.5, ui.InkSoft, true)
	}
}

func drawSleepMarks(screen *ebiten.Image, x, y float32) {
	vector.StrokeLine(screen, x, y, x+5, y, 1.2, ui.InkSoft, true)
	vector.StrokeLine(screen, x+5, y, x, y+6, 1.2, ui.InkSoft, true)
	vector.StrokeLine(screen, x, y+6, x+5, y+6, 1.2, ui.InkSoft, true)
	vector.StrokeLine(screen, x+6, y-5, x+10, y-5, 1.1, ui.InkSoft, true)
	vector.StrokeLine(screen, x+10, y-5, x+6, y, 1.1, ui.InkSoft, true)
	vector.StrokeLine(screen, x+6, y, x+10, y, 1.1, ui.InkSoft, true)
}

func drawPetFeedback(screen *ebiten.Image, x, y float32, kind string, ticks int) {
	if ticks <= 0 {
		return
	}

	progress := 1 - float32(ticks)/28
	rise := progress * 12

	switch kind {
	case "feed":
		drawSnackFeedback(screen, x, y-rise)
	case "brush":
		drawBrushFeedback(screen, x, y-rise)
	case "sleep":
		drawSleepFeedback(screen, x, y-rise)
	}
}

func drawSnackFeedback(screen *ebiten.Image, x, y float32) {
	drawDrumstick(screen, float64(x-14), float64(y+5), ui.Peach)
	vector.DrawFilledCircle(screen, x+4, y+3, 2.3, ui.Peach, true)
	vector.DrawFilledCircle(screen, x+11, y-3, 1.8, ui.BlushStrong, true)
	vector.DrawFilledCircle(screen, x+17, y+4, 1.5, ui.Mint, true)
}

func drawBrushFeedback(screen *ebiten.Image, x, y float32) {
	drawHeart(screen, x-10, y+3, 5, ui.BlushStrong)
	drawHeart(screen, x+8, y-2, 4, ui.Peach)
	drawSparkle(screen, x+18, y+5, 4, ui.Mint)
}

func drawSleepFeedback(screen *ebiten.Image, x, y float32) {
	drawMoon(screen, float64(x-10), float64(y+2), 5, ui.BabyBlue)
	vector.DrawFilledCircle(screen, x+6, y+3, 3.2, ui.Cream, true)
	vector.DrawFilledCircle(screen, x+13, y-3, 2.2, ui.Cream, true)
	vector.DrawFilledCircle(screen, x+18, y+5, 1.7, ui.Cream, true)
	drawSleepMarks(screen, x+2, y-10)
}

func drawSparkle(screen *ebiten.Image, x, y, size float32, fill color.Color) {
	vector.StrokeLine(screen, x-size, y, x+size, y, 1.2, fill, true)
	vector.StrokeLine(screen, x, y-size, x, y+size, 1.2, fill, true)
	vector.StrokeLine(screen, x-size*0.6, y-size*0.6, x+size*0.6, y+size*0.6, 1.0, fill, true)
	vector.StrokeLine(screen, x-size*0.6, y+size*0.6, x+size*0.6, y-size*0.6, 1.0, fill, true)
}
