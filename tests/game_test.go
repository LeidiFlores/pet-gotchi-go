package tests

import (
	"testing"

	"petgotchi/internal/game"
	"petgotchi/internal/scene"

	"github.com/hajimehoshi/ebiten/v2"
)

func TestNewGameLayout(t *testing.T) {
	g := game.New()
	if g == nil {
		t.Fatal("expected non-nil game")
	}

	w, h := g.Layout(10, 20)
	if w != game.ScreenW || h != game.ScreenH {
		t.Fatalf("expected layout %dx%d, got %dx%d", game.ScreenW, game.ScreenH, w, h)
	}
}

func TestGameUpdateDrawNoPanic(t *testing.T) {
	g := game.New()

	if err := g.Update(); err != nil {
		t.Fatalf("Update returned error: %v", err)
	}

	screen := ebiten.NewImage(game.ScreenW, game.ScreenH)
	defer screen.Dispose()
	g.Draw(screen)
}

func TestNewMenuUpdateDraw(t *testing.T) {
	m := scene.NewMenu(func(scene.Scene) {})
	if m == nil {
		t.Fatal("expected non-nil menu")
	}

	next, err := m.Update()
	if err != nil {
		t.Fatalf("Menu.Update returned error: %v", err)
	}
	if next == nil {
		t.Fatal("Menu.Update returned nil scene")
	}

	screen := ebiten.NewImage(game.ScreenW, game.ScreenH)
	defer screen.Dispose()
	m.Draw(screen)
}
