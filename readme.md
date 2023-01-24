
General resources:

    - Udemy Go course: https://www.udemy.com/course/go-the-complete-developers-guide/learn
    - Udemy web developer course: https://www.udemy.com/course/go-programming-language and repo at https://github.com/GoesToEleven/golang-web-dev
    - Learn Go Programming by Building 11 Projects (freecodecamp): https://www.youtube.com/watch?v=jFfo23yIWac
    - Pluralsight: https://app.pluralsight.com/paths/skill/go-core-language


Good tutorial on unit testing web requests: https://codeburst.io/unit-testing-for-rest-apis-in-go-86c70dada52d


Go resources:

    - Getting started covers adding packages and links to modules tutorial https://go.dev/doc/tutorial/getting-started
    - Code organisation (creating packages): https://go.dev/doc/code
    - Go packages: https://go.dev/pkg
    - Effective Go (see note near top about being outdated): https://go.dev/doc/effective_go


Packages:
    - collections of source files in the same directory that are compiled together
    - each go file declares which package it is part of
    - a program can be made up of multiple packages
    - to use code in a package include it via the import keyword
    - standard library packages installed with go are at https://pkg.go.dev/std


Dependencies:
    - are tracked through the projects module file called go.mod.
    - create go.mod file by something like 'go mod init example/cards'.
    - to publish a module, the module path in go.mod must be a location from which Go tools can download your module. 
    - for more, refer: https://go.dev/doc/modules/managing-dependencies


Workspace:
    - by default it is at $HOME/go.
    - all installed packages (i.e. from doing go install) are here. This is the GOPATH.





To run: Doing 'go run main.go' (or replace 'main.go' with '.') executes the program.

A go file using the reserved name 'main' is requied to generate an executable. The 'main.go' file needs a function called 'main'.

To build: Doing 'go build main.go' will generate an executable.

To run tests: Do 'go test'. Needed a go.mod file to exist.

On work laptop:

    Cant run own build binaries unless they are in the c:\dev folder.
    Doing 'go run' builds and runs from a temp directory.  To see which directory do 'go run -work main.go'.
    To change directory used set the GOTMPDIR environment variable. For example do :$env:GOTMPDIR = "C:\dev" (or set permanaently)


TYPES & POINTERS

An array has a fixed size. A slice (data type in its own right) is a dynamically-sized flexible view into the elements of an array. 

A struct is a collection of fields.

Types are either value or reference types, where reference means the type has a reference to another type.

Go passes by value which means a copy will be created when passing something to a function.

The data types of bool, integer, float, string, array and struct are value types. To change in a function will need to use pointers to avoid copying:

    - The & operator applied to a variable will return the address in memory of that variable (known as a pointer)
    - The * operator:
        - When used on a type means the type is a pointer to something of that type.
        - When used on a pointer will access the value in the memory address held by the pointer.

A map in Go is a collection of key-value pairs. All the keys must be of the same type and all the values must be of the same type however the type
used for keys can be different to the type used for values.

As a general rule use maps to represent a collection of very similar items and structs when the items can be very different. Note that for 
a struct you need to know the fields at compile time, and it's harder to iterate over the field values inside a struct.

Can create anonymous types which are useful if you need a data structure that won't be used again, perhaps when marshaling/unmarshalling data or
passing data to a template. For example:

        person := struct {
            fname string
            lname string
        }{
            fname: "Chris",
            lname: "Yoxall",
        }


GO ROUTINES

Go Routines:

    - Always get one main go routine started when the program starts.
    - Can create additional go routines by using the keyword 'go' in front of a function call.
    - The Go scheduler manages go routines and by default only uses 1 CPU. This means only 1 go routine actively runs at once no matter
        how many exist (concurrent processing).
    - Can utilise multiple CPU cores in which case there can be a go routine actively running on each CPU (parallel processing). Using 1 CPU
        is usually recommended.
    - Channels allow different go routines to communicate via the '<-' operator.



INTERFACES AND PATTERNS

Go is not an Object Oriented language. Patterns that get used in go are:

    - Define methods on types by adding a receiver in argument lists between the func keyword and method name. Functions with recievers 
        are referred to as methods.

    - Create interfaces. Interfaces are collections of method signatures or other interfaces. Types that implement everything in the
        interface have implemented the interface and can be referred to using the interface (polymorphism).

