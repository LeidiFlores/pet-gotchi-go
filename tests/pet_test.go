package tests

import (
	"petgotchi/internal/scene"
	"testing"
)

func TestNewPet(t *testing.T) {
	switchTo := func(scene.Scene) {} // dummy function
	pet := scene.NewPet(switchTo)

	if pet.Hunger != 50 {
		t.Errorf("Expected hunger 50, got %f", pet.Hunger)
	}
	if pet.Sleep != 50 {
		t.Errorf("Expected sleep 50, got %f", pet.Sleep)
	}
	if pet.Mood != 20 {
		t.Errorf("Expected mood 20, got %f", pet.Mood)
	}
	if pet.T != 0 {
		t.Errorf("Expected t 0, got %d", pet.T)
	}
	if pet.SwitchTo == nil {
		t.Error("Expected switchTo to be set")
	}
}

func TestClamp01(t *testing.T) {
	tests := []struct {
		input    float64
		expected float64
	}{
		{-5, 0},
		{0, 0},
		{50, 50},
		{100, 100},
		{105, 100},
	}

	for _, test := range tests {
		result := scene.Clamp01(test.input)
		if result != test.expected {
			t.Errorf("clamp01(%f) = %f; expected %f", test.input, result, test.expected)
		}
	}
}

func TestPetUpdate(t *testing.T) {
	switchTo := func(scene.Scene) {}
	pet := scene.NewPet(switchTo)

	// Initial values
	initialHunger := pet.Hunger
	initialSleep := pet.Sleep
	initialMood := pet.Mood

	// Update less than 60 times, stats should not change
	for i := 0; i < 59; i++ {
		pet.Update()
	}
	if pet.Hunger != initialHunger || pet.Sleep != initialSleep || pet.Mood != initialMood {
		t.Error("Stats changed before 60 ticks")
	}

	// Update to 60th tick
	pet.Update()
	if pet.Hunger != scene.Clamp01(initialHunger+1) {
		t.Errorf("Expected hunger %f, got %f", scene.Clamp01(initialHunger+1), pet.Hunger)
	}
	if pet.Sleep != scene.Clamp01(initialSleep+1) {
		t.Errorf("Expected sleep %f, got %f", scene.Clamp01(initialSleep+1), pet.Sleep)
	}
	if pet.Mood != scene.Clamp01(initialMood+0.5) {
		t.Errorf("Expected mood %f, got %f", scene.Clamp01(initialMood+0.5), pet.Mood)
	}
}
