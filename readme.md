
## Resources

General resources:

- Udemy Go course: https://www.udemy.com/course/go-the-complete-developers-guide/learn
- Udemy web developer course: https://www.udemy.com/course/go-programming-language and repo at https://github.com/GoesToEleven/golang-web-dev
- Pluralsight: https://app.pluralsight.com/paths/skill/go-core-language
- play-with-go https://play-with-go.dev/


Tutorials & Blog Posts:
- Unit testing web requests: https://codeburst.io/unit-testing-for-rest-apis-in-go-86c70dada52d
- Packages & modules: https://www.alexedwards.net/blog/an-introduction-to-packages-imports-and-modules


Go resources:
- Getting started covers adding packages and links to modules tutorial https://go.dev/doc/tutorial/getting-started.
- Code organisation (creating packages): https://go.dev/doc/code.
- Go packages: https://go.dev/pkg.
- Effective Go (see note near top about being outdated): https://go.dev/doc/effective_go.


Packages:
- Collections of source files in the same directory that are compiled together.
- Each go file declares which package it is part of.
- Types, functions & variables in a package are available to other packages if they start with a capital letter.
- To use code in a package include it via the import keyword.
- Must have a main package that contains a main() function (by convention this is in a main.go file).
- For non-main packages the directory should have the same name as the package.
- Standard library packages installed with go are at https://pkg.go.dev/std.


Modules:
- Are Go's Dependency management system.
- Consist of one or more packages with with a 'go.mod' file at its root.
- Are tracked through the projects module file called go.mod.
- Create go.mod file by something like 'go mod init example/cards'.
- The module path acts as a canonical identifier for a module.
- To publish a module, the module path in go.mod must be a location from which Go tools can download your module. 
- For more, refer: https://go.dev/doc/modules/managing-dependencies & https://go.dev/blog/using-go-modules


Workspace:
- Default location is $HOME/go.
- All installed packages (i.e. from doing go install) are here. This is the GOPATH.



## Work laptop

- Cant run own build binaries unless they are in the c:\dev folder.
- Doing 'go run' builds and runs from a temp directory.  To see which directory do 'go run -work main.go'.
- To change directory used set the GOTMPDIR environment variable. For example do :$env:GOTMPDIR = "C:\dev" (or set permanaently).
- When debugging found that I had to also set the output directory to a location within the c:\dev hierarchy as well.

## SSH

Have enabled SSH for this repo on GitHub. Are good instructions at https://docs.github.com/en/authentication/connecting-to-github-with-ssh/about-ssh

Currently on work laptop will need to:

1. Start the ssh-agent in the background: eval "$(ssh-agent -s)"
2. Add the private SSH key to the agent: ssh-add ~/.ssh/id_rsa_personal_github
3. Will need to push, pull etc via the command line rather than use VS Code running in WSL.


## Build and Run

To run: Doing 'go run main.go' (or replace 'main.go' with '.') executes the program.

A go file using the reserved name 'main' is requied to generate an executable. The 'main.go' file needs a function called 'main'.

To build: Doing 'go build main.go' will generate an executable.

To run tests: Do 'go test'. Needed a go.mod file to exist.


## TYPES

An array has a fixed size. A slice (data type in its own right) is a dynamically-sized flexible view into the elements of an array. 

A struct is a collection of fields.

A map in Go is a collection of key-value pairs. All the keys must be of the same type and all the values must be of the same type however the type
used for keys can be different to the type used for values.

As a general rule use maps to represent a collection of very similar items and structs when the items can be very different. Note that for 
a struct you need to know the fields at compile time, and it's harder to iterate over the field values inside a struct.

Types are either value or reference types, where reference means the type has a reference to another type. The data types of bool, integer,
float, string, array and struct are value types.

Can create anonymous types which are useful if you need a data structure that won't be used again, perhaps when marshaling/unmarshalling data or
passing data to a template. For example:

    person := struct {
        fname string
        lname string
    }{
        fname: "Chris",
        lname: "Yoxall",
    }

Go passes by value which means a copy will be created when passing something to a function. To change a value type in a function need
to use pointers so the actual type is changed rather than the copy (referency types use pointers by default).


## POINTERS

- The & operator applied to a variable will return the address in memory of that variable (known as a pointer)
- The * operator:
    - When used on a type means the type is a pointer to something of that type.
    - When used on a pointer will access the value in the memory address held by the pointer.

## GO ROUTINES

Always get one main go routine started when the program starts.

Can create additional go routines by using the keyword 'go' in front of a function call.

The Go scheduler manages go routines and by default only uses 1 CPU. This means only 1 go routine actively runs at once no matter
how many exist (concurrent processing).

Can utilise multiple CPU cores in which case there can be a go routine actively running on each CPU (parallel processing). Using 1 CPU
is usually recommended.

Channels allow different go routines to communicate via the '<-' operator.



## INTERFACES AND PATTERNS

Go is not an Object Oriented language. Patterns that get used in go are:

- Define methods on types by adding a receiver in argument lists between the func keyword and method name. Functions with recievers 
    are referred to as methods.

- Create interfaces. Interfaces are collections of method signatures or other interfaces. Types that implement everything in the
    interface have implemented the interface and can be referred to using the interface (polymorphism).


## Static Compilation

A lot of this section came from:

- https://www.arp242.net/static-go.html
- https://mt165.co.uk/blog/static-link-go


By default Go creates static binaries unless cgo (https://pkg.go.dev/cmd/cgo) is used to call C code which results in a dynamically linked
binary. To check whether the binary is statically compiled use 'file' or 'ldd'

    file FILENAME

or

    ldd FILENAME

Two packages in the Go standard library use cgo:

- os/user: User account lookups. On Windows there is only the Go implementation so this does not apply.
- net: Interface for network I/O, including TCP/IP, UDP, domain name resolution, and Unix domain sockets. 

The Go binary will not be statically linked if it imports one of those two packages above, either directly or through a dependency. Can
use the 'osuergo' and 'netgo' build tags to skip building the cgo parts:

    go build -tags osusergo,netgo

Or disable cgo completely:

    CGO_ENABLED=0 go build

This still results in a working binary as there are pure Go versions that wil be used instead but they are not totally feature compatible
with the C implementations. They are close however so its usually fine.

If you want to use cgo and have a statically linked binary can do:

    go build -ldflags "-linkmode 'external' -extldflags '-static'"

That command above tells the Go toolchain to use an external linker and gets that linker to produce a static binary. A lot of sites say that
specifying an external linker is not needed.

However you will usually get some warnings when using glibc (most common implementation of C statndard libraries) as glibc itself has dependencies
on other libraries. If your binary doesn't rely on the functions in the warnings, again you should be fine.

You can use musl (https://wiki.musl-libc.org/) to replace glibc which may help solve this. To use musl, you can either install it and build your
software using musl-gcc, or you can use a Linux distribution that uses musl, e.g. Alpine Linux, as a build environment.

In conclusion - relying on a static binary has risks.....
