package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"petgotchi/internal/scene"
)

const (
	ScreenW = 320
	ScreenH = 180
)

type Game struct {
	current scene.Scene
}

func New() *Game {
	g := &Game{}
	g.current = scene.NewMenu(func(next scene.Scene){ gSwitch(g, next) })
	return g
}

func gSwitch(g *Game, s scene.Scene) {
	g.current = s
}

func (g *Game) Update() error {
	return g.current.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.current.Draw(screen)
}

func (g *Game) Layout(w, h int) (int, int) {
	return ScreenW, ScreenH
}
