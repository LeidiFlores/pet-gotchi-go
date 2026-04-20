package tests

import (
	"testing"

	"petgotchi/internal/ui"
)

func TestButtonUpdateHoverAndClick(t *testing.T) {
	clicked := false
	btn := &ui.Button{
		X: 0,
		Y: 0,
		W: 10,
		H: 10,
		OnClick: func() {
			clicked = true
		},
	}

	btn.Update(5, 5, false, false, false)
	if !btn.Hover {
		t.Error("expected button to be hovered when pointer is inside")
	}
	if btn.Pressed {
		t.Error("expected button not to be pressed when mouseDown is false")
	}
	if clicked {
		t.Error("expected button not to click before a press")
	}

	btn.Update(5, 5, true, true, false)
	if !btn.Pressed {
		t.Error("expected button to be pressed when hovered and mouseDown is true")
	}
	if clicked {
		t.Error("expected button not to click until mouse release")
	}

	btn.Update(5, 5, false, false, true)
	if !clicked {
		t.Error("expected button to click when released over the button")
	}
	if btn.Pressed {
		t.Error("expected button not to stay pressed after release")
	}
}

func TestButtonFastClick(t *testing.T) {
	clicked := false
	btn := &ui.Button{
		X: 0, Y: 0, W: 10, H: 10,
		OnClick: func() { clicked = true },
	}

	// press and release in the same tick (justPressed=true, justReleased=true)
	btn.Update(5, 5, false, true, true)
	if !clicked {
		t.Error("expected button to click on fast click (same-tick press and release)")
	}
}

func TestButtonUpdateOutside(t *testing.T) {
	clicked := false
	btn := &ui.Button{
		X: 0,
		Y: 0,
		W: 10,
		H: 10,
		OnClick: func() {
			clicked = true
		},
	}

	btn.Update(20, 20, true, true, false)
	if btn.Hover {
		t.Error("expected button not to be hovered when pointer is outside")
	}
	if btn.Pressed {
		t.Error("expected button not to be pressed when pointer is outside")
	}
	if clicked {
		t.Error("expected button not to click when pointer is outside")
	}

	btn.Update(20, 20, false, false, true)
	if clicked {
		t.Error("expected button not to click when released outside")
	}
}
