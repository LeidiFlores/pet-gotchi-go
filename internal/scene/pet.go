package scene

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Pet struct {
	SwitchTo func(Scene)
	// stats
	Hunger float64 // 0..100 (0 = full, 100 = hungry)
	Sleep  float64 // 0..100 (0 = awake, 100 = sleepy)
	Mood   float64 // 0..100 (0 = happy, 100 = sad)
	T      int
}

func NewPet(switchTo func(Scene)) *Pet {
	return &Pet{SwitchTo: switchTo, Hunger: 50, Sleep: 50, Mood: 20}
}

func (p *Pet) Update() error {

	p.T++
	if p.T%60 == 0 {
		p.Hunger = Clamp01(p.Hunger + 1)
		p.Sleep = Clamp01(p.Sleep + 1)
		p.Mood = Clamp01(p.Mood + 0.5)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		p.Hunger = Clamp01(p.Hunger - 2)
		p.Mood = Clamp01(p.Mood - 0.5)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyN) {
		p.Sleep = Clamp01(p.Sleep - 2)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyB) {
		p.Mood = Clamp01(p.Mood - 2)
	}
	return nil
}

func (p *Pet) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{255, 247, 251, 255})

	r := 16.0 + 2.0*math.Sin(float64(p.T)/10.0)
	vector.DrawFilledCircle(screen, 160, 90, float32(r), color.RGBA{255, 198, 216, 255}, true)
	ebitenutil.DebugPrintAt(screen, "F: feed  B: brush  N: sleep", 10, 10)

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Hunger: %3.0f", p.Hunger), 10, 40)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Sleep: %3.0f", p.Sleep), 10, 55)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Mood: %3.0f", p.Mood), 10, 70)
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
