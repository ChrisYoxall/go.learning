package main

import "fmt"

// In a generic function inside the [ ] is the type parameter (usually T) and the constraint that defines what
// the type parameter can be, which is 'any' in this example.  Cant do 'a + b' here, as while 'any' generally works
// as a type constraint for simple tasks it does not know how to deal with operators like +, -, *, /, etc.
func sum[T any](a, b T) {
	fmt.Println(a, b)
}

// SumConstraint defines a custom constraint.
type SumConstraint interface {
	int | float64 | string
}

func sumWithCustomConstraint[T SumConstraint](a, b T) T {
	return a + b
}

func main() {

	fmt.Println("Using 'any':")
	sum(1, 2)
	sum(4.0, 2.0)
	sum("a", "b")

	fmt.Println("Using custom constraint:")
	fmt.Println(sumWithCustomConstraint(1, 2))
	fmt.Println(sumWithCustomConstraint(4.0, 2.0))
	fmt.Println(sumWithCustomConstraint("a", "b"))

}
