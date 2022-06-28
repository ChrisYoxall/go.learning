package main

import "fmt"

// Create new type called bot.
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func (eb englishBot) getGreeting() string {
	// Would have a lot of logic here related to speaking english
	return "Hello"
}

func (sb spanishBot) getGreeting() string {
	// Would have a lot of logic here related to speaking spanish
	return "Hola"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
