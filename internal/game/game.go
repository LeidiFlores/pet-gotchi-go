package game

import (
	"petgotchi/internal/scene"

	"github.com/hajimehoshi/ebiten/v2"
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
	g.current = scene.NewMenu(func(next scene.Scene) { gSwitch(g, next) })
	return g
}

func gSwitch(g *Game, s scene.Scene) {
	g.current = s
}

func (g *Game) Update() error {
	nextScene, err := g.current.Update()

	if err != nil {
		return err
	}

	g.current = nextScene
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.current.Draw(screen)
}

func (g *Game) Layout(w, h int) (int, int) {
	return ScreenW, ScreenH
}
