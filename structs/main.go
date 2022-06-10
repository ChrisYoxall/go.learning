package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {

	// Several ways to create structs. This approach assumes first name will go to the
	// first field defined in the struct.
	//chris := person{"Chris", "Yoxall"}

	// Can specify which field a value should go to. Safer in case struct definition changes and
	// the order of the fields changes.
	chris := person{firstName: "Chris", lastName: "Yoxall"}

	// Can also define using long variable definition approach. Here bob is assigned an empty
	// string for firstName and lastName as that is the default zero value for string in Go.
	var bob person
	fmt.Println(bob)
	fmt.Printf("%+v\n", bob)
	bob.firstName = "Bob"
	bob.lastName = "The Great!"
	fmt.Println(bob)

	fmt.Println(chris)

}
