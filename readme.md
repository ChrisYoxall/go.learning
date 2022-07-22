
Go course on Udemy: https://www.udemy.com/course/go-the-complete-developers-guide/learn/

https://www.youtube.com/watch?v=1MXIGYrMk80

freeCodeCamp.org learn go: https://www.youtube.com/watch?v=YS4e4q9oBaU

Go resources:

    - Code organisation (packages): https://go.dev/doc/code
    - Go packages: https://go.dev/pkg

    
Good tutorial on unit testing web requests: https://codeburst.io/unit-testing-for-rest-apis-in-go-86c70dada52d


To run:  Doing 'go run main.go' executes program.

To build: Doing 'go build main.go' will generate an executable.

Use of the package name 'main' (reserved word) means that will generate an executable. The 'main.go' file needs a function called 'main'

To enable 'go test' to work created a package by doing 'go mod init example/cards'


An array has a fixed size. A slice (data type in its own right) is a dynamically-sized flexible view into the elements of an array. In practice, slices are much more common than arrays.

A struct is a collection of fields.

Go passes by value which means a copy will be created when passing something to a function. For primitive data types (bool, integer, float, string, array and struct) when they are passed to a 
function that needs to change the values you must use pointers where:

    - The & operator applied to a variable will return the address in memory of that variable (known as a pointer)
    - The * operator:
        - When used on a type means the type is a pointer to something of that type.
        - When used on a pointer will access the value in the memory address held by the pointer.

Types are either value or reference types, where reference means the type has a reference to another type.

A map in Go is a collection of key-value pairs. All the keys must be of the same type and all the values must be of the same type however the type
used for keys can be different to the type used for values.

As a general guide use maps to represent a collection of very similar items. Use structs when the items can be very different. Note that for 
a struct you need to know the fields at compile time, and it's harder to iterate over the field values inside a struct.


Go Routines:

    - Always get one main go routine started when the program starts.
    - Can create additional go routines by using the keyword 'go' in front of a function call.
    - The Go scheduler manages go routines and by default only uses 1 CPU. This means only 1 go routine actively runs at once no matter how many exist (concurrent processing).
    - Can utilise multiple CPU cores in which case there can be a go routine actively running on each CPU (parallel processing). Using 1 CPU is usually recommended.
    - Channels allow different go routines to communicate via the '<-' operator.



Go is not an Object Oriented language. Patterns that get used in go are:

    - Define methods on types by adding a receiver in argument lists between the func keyword and method name. Functions with recievers 
    are referred to as methods.

    - Create interfaces. Interfaces are collections of the signature of methods or other interfaces. Types that implement everything in the interface have implemented the interface and can
    be referred to using the interface (polymorphism).

