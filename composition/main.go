package main

import "fmt"

type sell interface {
	sell()
}

type show interface {
	show()
}

// Create a composite interface interface by embedding the sell and show interfaces
type salesman interface {
	sell
	show
}

type properties struct {
	genre  string
	rating string
}

// Embed the composite game struct to by embedding the properties struct
type game struct {
	properties
	name  string
	price int
}

// For a salseman to be able to sell and show a game to the customer then we need to implement the salesman inteface
// which is made up of the sell and show interafces

func (g game) sell() {
	fmt.Printf("Selling '%s' for '$%.2f'\n", g.name, float32(g.price/100))
}

func (g game) show() {
	fmt.Println("------ Game Details ------")
	fmt.Println("Name: ", g.name)
	fmt.Println("Genre: ", g.genre)
	fmt.Println("Rating: ", g.rating)
	fmt.Printf("Price: %.2f\n", float32(g.price/100))
	fmt.Println("--------------------------")
}

// This method takes the composite interface as a parameter so anything that has implemented that
// interface can be passed in
func shop(s salesman) {
	s.show()
	s.sell()
}

func main() {

	g := game{
		name:  "Ultima V",
		price: 12500,
		properties: properties{
			genre:  "Adventure",
			rating: "G",
		}}

	shop(g)
}
