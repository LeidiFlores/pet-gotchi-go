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

	btn.Update(5, 5, false)
	if !btn.Hover {
		t.Error("expected button to be hovered when pointer is inside")
	}
	if clicked {
		t.Error("expected button not to click when justPressed is false")
	}

	btn.Update(5, 5, true)
	if !clicked {
		t.Error("expected button to click when hovered and justPressed is true")
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

	btn.Update(20, 20, true)
	if btn.Hover {
		t.Error("expected button not to be hovered when pointer is outside")
	}
	if clicked {
		t.Error("expected button not to click when pointer is outside")
	}
}
