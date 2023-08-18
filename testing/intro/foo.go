package main

import "strconv"

// Foo returns "Foo" if the number is divisible by 3, otherwise returns the number as a string.
func Foo(input int) string {

	if (input % 3) == 0 {
		return "Foo"
	}

	return strconv.Itoa(input)
}
