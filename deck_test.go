package main

import "testing"

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
