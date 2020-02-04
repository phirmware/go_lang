package main

import "testing"

func TestGame(t *testing.T) {
	input := readInput()
	if input != "paper" {
		t.Errorf("Expected Paper but got %v", input)
	}
}
