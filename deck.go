package main

import "fmt"

// Create 'deck' type which is a slice of strings.
type deck []string

// Create method by adding receiver before function name.
// By convention the reference (d in this case) is a very short abbreviation.
// While d is similar to 'this' or 'self' by convention go avoids those words.
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
