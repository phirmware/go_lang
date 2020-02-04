package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 12 {
		t.Errorf("Lenght not right, got length of %v", len(d))
	}

	if d[0] != "ace of spades" {
		t.Errorf("First item not correct, got %v", d[0])
	}

	if d[len(d)-1] != "four of heart" {
		t.Errorf("Last Item not correct, got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndReadFile(t *testing.T) {
	fileName := "testing"
	os.Remove(fileName)

	d := newDeck()
	d.saveToFile(fileName)

	deckFromFile := readFile(fileName)
	length := len(deckFromFile)
	if length != 12 {
		t.Errorf("Expected length of 16 from file but got %v", length)
	}

	os.Remove(fileName)
}
