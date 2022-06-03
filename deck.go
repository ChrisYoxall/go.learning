package main

import (
	"fmt"
	"os"
	"strings"
)

// Create 'deck' type which is a slice of strings.
type deck []string

// Return string representation of deck. This is an example of a method
// in Go, where a method is a function with a receiver argument placed
// before the function name.
// By convention the reference (d in this case) is a very short abbreviation.
// While d is similar to 'this' or 'self' by convention go avoids those words.
func (d deck) toString() string {

	// Do a type conversion from deck to slice of strings
	return strings.Join([]string(d), ", ")
}

// Create a deck of 52 cards.
func newDeck() deck {

	cardSuits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six",
		"Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	cards := deck{}

	// Note how _ is used instead of the index as the index is not used.
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards

}

// Method to print out values in deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Return two values, both of type deck. Use slice notation to create new decks to return.
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Save a deck to disk
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

// Read a deck from a file
func newDeckFromFile(filenmae string) deck {

	bs, err := os.ReadFile(filenmae)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// Convert bs to string, use split to return a slice of strings then convert to deck
	return deck(strings.Split(string(bs), ", "))
}
