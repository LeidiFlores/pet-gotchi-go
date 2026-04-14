package scene

import "github.com/hajimehoshi/ebiten/v2"

// Scene defines the contract every scene must implement.
type Scene interface {
	Update() (Scene, error)
	Draw(screen *ebiten.Image)
}
