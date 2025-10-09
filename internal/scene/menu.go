package scene

import (
	"image/color"

	"petgotchi/internal/ui"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Scene interface {
	Update() error
	Draw(*ebiten.Image)
}

type Menu struct {
	btnStart *ui.Button
	switchTo func(Scene)
}

func NewMenu(switchTo func(Scene)) *Menu {
	m := &Menu{switchTo: switchTo}
	m.btnStart = &ui.Button{X: 100, Y: 100, W: 120, H: 32, Label: "Jugar ♥", OnClick: func() {
		m.switchTo(NewPet(m.switchTo))
	}}
	return m
}

func (m *Menu) Update() error {
	x, y := ebiten.CursorPosition()
	justPressed := inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)
	m.btnStart.Update(x, y, justPressed)
	return nil
}

func (m *Menu) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{255, 247, 251, 255})
	ebitenutil.DebugPrintAt(screen, "Pet-gotchi ♥", 10, 10)
	m.btnStart.Draw(screen)
}
