package main

import (
	"log"

	"petgotchi/internal/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowTitle("Pet-gotchi ♥")
	ebiten.SetWindowResizable(true)
	g := game.New()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
