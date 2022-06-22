package main

import "fmt"

type contactInfo struct {
	email    string
	postcode string
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo // Can just put contactInfo here to create field called contactInfo
}

func main() {

	// Several ways to create structs. This approach assumes first name will go to the
	// first field defined in the struct.
	//chris := person{"Chris", "Yoxall"}

	// Can specify which field a value should go to. Safer in case struct definition changes and
	// the order of the fields changes. Note extra comments when using multi-line format.
	chris := person{
		firstName: "Chris",
		lastName:  "Yoxall",
		contact: contactInfo{
			email:    "chris.yoxall@gmail.com",
			postcode: "1024",
		},
	}

	fmt.Println(chris)
	fmt.Printf("%+v\n", chris)

	// Can also define using long variable definition approach. Here bob is assigned an empty
	// string for firstName and lastName as that is the default zero value for string in Go.
	// Note you don't need to supply contactInfo
	var bob person
	fmt.Println(bob)
	fmt.Printf("%+v\n", bob)
	bob.firstName = "Bob"
	bob.lastName = "The Great!"
	fmt.Println(bob)

	// Note use of pointer syntax
	dani := person{firstName: "dani"}
	daniPointer := &dani
	daniPointer.updateFirstName("Dani")
	fmt.Println(dani)

	// Go has a shorcut, so that as long as the function receives a value of type pointer
	// you can just pass the variable without having to create the pointer and Go creates
	// the pointer for you
	dani.updateFirstName("Dani created via shortcut")
	fmt.Println(dani)
}

func (pointerToPerson *person) updateFirstName(firstName string) {
	(*pointerToPerson).firstName = firstName
}
