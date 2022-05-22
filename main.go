package main

import "fmt"

func main() {

	// declare a slice and set some values
	cards := []string{"Two of Clubs", newCard()}
	cards = append(cards, "Seven of Diamonds")

	// Note the use of := which means that index and card is thrown away and re-created each interation.
	for index, card := range cards {
		fmt.Println(index, card)
	}

}

func newCard() string {
	return "Ace of Spaces"
}
