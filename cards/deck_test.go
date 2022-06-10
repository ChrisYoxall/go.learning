package main

import (
	"os"
	"testing"
)

// Note how test functions start with upper case.

func TestNewDeck(t *testing.T) {

	d := newDeck()

	size := 52

	if len(d) != size {
		t.Errorf("Expected a deck size of %v but got %v", size, len(d))
	}

	if d[0] != "Ace of Hearts" {
		t.Errorf("Expected first card to be 'Ace of Hearts' but got %v'", d[0])
	}

	if d[size-1] != "King of Spades" {
		t.Errorf("Expected last card to be 'King of Spades' but got %v'", d[size-1])
	}

}

func TestSaveDeckToFileAndNewDeckFromFile(t *testing.T) {

	file := "_decktesting"

	// Make sure any test files left from previous runs are removed.
	os.Remove(file)

	d := newDeck()
	d.saveToFile(file)

	loadedDeck := newDeckFromFile(file)

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards but got %v", len(loadedDeck))
	}

	// Remove tets file.
	os.Remove(file)
}
