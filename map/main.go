package main

import "fmt"

func main() {

	// Declare a variale of type map
	var myMap map[string]string
	fmt.Println(myMap)

	// Use a built in function to create a map
	names := make(map[string]string)

	// Add value to map
	names["chris"] = "Chris Yoxall"
	fmt.Println(names)

	// Delete a value in a map
	delete(names, "chris")
	fmt.Println(names)

	// Create and add values to map
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	fmt.Println(colors)

	printColors(colors)
}

func printColors(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
