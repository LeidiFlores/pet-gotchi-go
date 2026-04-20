package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type pointerState struct {
	x            int
	y            int
	down         bool
	justPressed  bool
	justReleased bool
}

func readPointerState() pointerState {
	touchIDs := ebiten.AppendTouchIDs(nil)
	if len(touchIDs) > 0 {
		x, y := ebiten.TouchPosition(touchIDs[0])
		return pointerState{x: x, y: y, down: true}
	}

	releasedTouchIDs := inpututil.AppendJustReleasedTouchIDs(nil)
	if len(releasedTouchIDs) > 0 {
		x, y := inpututil.TouchPositionInPreviousTick(releasedTouchIDs[0])
		return pointerState{x: x, y: y, justReleased: true}
	}

	x, y := ebiten.CursorPosition()
	return pointerState{
		x:            x,
		y:            y,
		down:         ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft),
		justPressed:  inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft),
		justReleased: inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft),
	}
}
