package main

import "fmt"

// Create 'deck' type which is a slice of strings.
type deck []string

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

// Create method by adding receiver before function name.
// By convention the reference (d in this case) is a very short abbreviation.
// While d is similar to 'this' or 'self' by convention go avoids those words.
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
