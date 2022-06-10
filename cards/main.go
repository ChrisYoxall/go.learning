package main

import "fmt"

func main() {

	cards := newDeck()

	cards.shuffle()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()

	cards.saveToFile("myCards")
	cardsFromDisk := newDeckFromFile("myCards")
	fmt.Println(cardsFromDisk)

}
