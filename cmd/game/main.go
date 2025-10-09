package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"petgotchi/internal/game"
)

func main() {
	ebiten.SetWindowTitle("Pet-gotchi ♥")
	ebiten.SetWindowResizable(true)
	g := game.New()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
