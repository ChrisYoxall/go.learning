package main

func main() {

	cards := deck{"Two of Clubs", newCard()}

	cards = append(cards, "Seven of Diamonds")

	cards.print()

}

func newCard() string {
	return "Ace of Spaces"
}
