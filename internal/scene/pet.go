package scene

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Pet struct {
	switchTo func(Scene)
	// stats
	hunger float64 // 0..100 (0 = lleno, 100 = hambriento)
	sleep  float64 // 0..100 (0 = descansado, 100 = con sueño)
	mood   float64 // 0..100 (0 = feliz, 100 = triste)
	t      int
}

func NewPet(switchTo func(Scene)) *Pet {
	return &Pet{switchTo: switchTo, hunger: 50, sleep: 50, mood: 20}
}

func (p *Pet) Update() error {
	// Decaimiento suave con el tiempo
	p.t++
	if p.t%60 == 0 { // cada ~1s a 60fps
		p.hunger = clamp01(p.hunger + 1)
		p.sleep = clamp01(p.sleep + 1)
		p.mood = clamp01(p.mood + 0.5)
	}

	// Controles: F alimenta, N siesta, B cepillar
	if ebiten.IsKeyPressed(ebiten.KeyF) {
		p.hunger = clamp01(p.hunger - 2)
		p.mood = clamp01(p.mood - 0.5)
	}
	if ebiten.IsKeyPressed(ebiten.KeyN) {
		p.sleep = clamp01(p.sleep - 2)
	}
	if ebiten.IsKeyPressed(ebiten.KeyB) {
		p.mood = clamp01(p.mood - 2)
	}
	return nil
}

func (p *Pet) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{255, 247, 251, 255})

	// Mascota placeholder: un círculo que late
	r := 16.0 + 2.0*math.Sin(float64(p.t)/10.0)
	vector.DrawFilledCircle(screen, 160, 90, float32(r), color.RGBA{255, 198, 216, 255}, true)
	ebitenutil.DebugPrintAt(screen, "F: alimentar  B: cepillar  N: siesta", 10, 10)

	// HUD simple
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Hambre: %3.0f", p.hunger), 10, 40)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Sueño : %3.0f", p.sleep), 10, 55)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Ánimo : %3.0f", p.mood), 10, 70)
}

func clamp01(v float64) float64 {
	if v < 0 {
		return 0
	}
	if v > 100 {
		return 100
	}
	return v
}
