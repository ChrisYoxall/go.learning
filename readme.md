
## Resources

Go resources:

- Go documentation: https://go.dev/doc/     There is a lot of stuff here and can be used as a resource for leaning Go.
- Effective Go (see note near top about being outdated): https://go.dev/doc/effective_go
- Go style guide: https://google.github.io/styleguide/go/index
- Code organisation (creating packages): https://go.dev/doc/code.
- Go packages: https://go.dev/pkg.
- Go Vulnerability Management: https://pkg.go.dev/golang.org/x/vuln

GoLand Documentation: https://www.jetbrains.com/help/go/getting-started.html


Frameworks to experiment with:
- GoFr is an Opinionated Go Framework for accelerated microservice development: https://gofr.dev/ (1.3k stars on GitHub)
- Go kit is a toolkit for microservices https://gokit.io/ (26k stars on GitHub)


## Packages

- Collections of source files within the same directory that are compiled together.
- Code in a package can use all types, constants, variables, functions etc. defined within the package even when declared in a different file.
- Types, functions & variables in a package are available to other packages if they start with a capital letter.
- Every Go file must be a part of a package. The package it is part of is declared at the top of the file.
- Can have init() functions that will be run when the package is initialised.
- Standard library packages installed with go are at https://pkg.go.dev/std.
- To use code in a package include it in the current file by using the 'import' keyword.
- To include packages you create as part of an application specify the module name before the package name (i.e. import "[module-name]/[package-name]")
- For non-main packages the directory should have the same name as the package.
- A Go executable needs an entrypoint main() function defined in a special package called main (by convention in a main.go file).

The introduction of modules changed how packages work. I don't think you can reference other user packages without using a module now. Use to have to use $GOPATH prior to modules.

The internal folder is a special folder that can only be imported by packages within the tree rooted at the parent of the internal directory


## Modules

- The unit of code distribution and versioning in Go.
- Consist of one or more packages with a 'go.mod' file at its root.
- Create go.mod file by something like 'go mod init example/cards'.
- The module path acts as a canonical identifier for a module, and is typically an URL where the module is hosted.
- When referencing packages in the same module the import statement is 'import {module path}/{path to the package relative to your go.mod file}'
- To publish a module, the module path in go.mod must be a location from which Go tools can download your module. 
- For more, refer: https://go.dev/doc/modules/managing-dependencies & https://go.dev/blog/using-go-modules & https://go.dev/doc/code#ImportingRemote


Workspace:
- Default location is $HOME/go.
- All installed packages (i.e. from doing go install) are here. This is the GOPATH.


## Build, Run & Test

The 'run' command compiles and runs the named go package. Can specify the package by either
- Supplying a list of files, for example 'go run main.go deck.go'
- Pattern matching a single package, for example 'go run .', 'go run *.go' or 'go run example/cards'
- Refer https://pkg.go.dev/cmd/go#hdr-Compile_and_run_Go_program (which is part of the standard library documentation on the Go tool)

Doing 'go build' compiles but does not install the result and ignores files that end in '_test.go'. See https://pkg.go.dev/cmd/go#hdr-Compile_packages_and_dependencies

Doing 'go test' (include -v for verbose) recompiles and runs tests. See https://pkg.go.dev/cmd/go#hdr-Test_packages



## TYPES

An array has a fixed size. A slice (data type in its own right) is a dynamically-sized flexible view into the elements of an array. 

A struct is a collection of fields.

A map in Go is a collection of key-value pairs. All the keys must be of the same type and all the values must be of the same type however the type
used for keys can be different to the type used for values.

As a general rule use maps to represent a collection of very similar items and structs when the items can be very different. Note that for 
a struct you need to know the fields at compile time, and it's harder to iterate over the field values inside a struct.

Types are either value or reference types, where reference means the type has a reference to another type:
- Value types: bool, integer, float, string, array and struct.
- Reference types: slice, map, channel, pointer, interface and functions.

When those reference types above are not set to a value, nil (not null) is used.

Can create anonymous types which are useful if you need a data structure that won't be used again, perhaps when marshaling/unmarshalling data or
passing data to a template. For example:

    person := struct {
        fname string
        lname string
    }{
        fname: "Chris",
        lname: "Yoxall",
    }

Go passes by value which means a copy will be created when passing something to a function. To change a value type in a function
need to use pointers so the actual type is changed rather than the copy (reference types use pointers by default).


## POINTERS

- The & operator applied to a variable will return the address in memory of that variable (known as a pointer)
- The * operator:
    - When used on a type means the type is a pointer to something of that type.
    - When used on a pointer will access the value in the memory address held by the pointer.


## GO ROUTINES

A goroutine is a lightweight thread of execution managed by the Go runtime. They are not actual OS threads and the main
routine itself is a goroutine. Create additional goroutines by using the keyword 'go' in front of a function call.

The Go scheduler manages go routines and by default only uses 1 CPU. This means only 1 go routine actively runs at once no matter
how many exist (concurrent processing).

Concurrency is not Parallelism. A quote from Rob Pike: "Concurrency is about dealing with lots of things at once. Parallelism
is about doing lots of things at once"

Can utilise multiple CPU cores in which case there can be a go routine actively running on each CPU (parallel processing). Using 1 CPU
is usually recommended.

Channels allow different go routines to communicate via the '<-' operator. Can:

- Send data to a channel: channel <- 5
- Receive data from a channel. Don't need the variable. (blocking statement): myVar := <- channel

Can specify a buffer for channels and also specify if a channel is either send or receive only (the default is both).

Receiving from a channel will block until there is data to receive. If the channel is unbuffered, sending to the channel blocks
until a receiver has received the value. If the channel has a buffer, the sender blocks only until the value has been copied to the
buffer and if the buffer is full this blocks until a value has been received which frees up a spot in the buffer to write the value.

The 'select' statement in go is similar in some way to switch but for channels. It allows you to wait on multiple channel operations.

Remember to close the channel when finished with it.

The sync package https://pkg.go.dev/sync provides a number of useful functions for synchronising go routines.

Search for 'Go Concurrency Patterns' when doing something non-trivial.

In concurrent programs, it's often necessary to preempt operations because of timeouts, cancellation, or failure of another portion of
the system. The context package makes it easy to pass request-scoped values, cancellation signals, and deadlines across API boundaries to
all the goroutines involved in handling a request.

Beware of using a loop variable in a closure running as a goroutine. Refer https://go.dev/doc/faq#closures_and_goroutines

To send a signal via a channel it's common to use a channel of type 'struct{}' which is zero size. Can create a 'struct{}'
by doing 'struct{}{}' which is a composite literal. See 'Done()' in side the context package https://pkg.go.dev/context#Context
for an example.



## INTERFACES AND PATTERNS

Go is not an Object-Oriented language. Patterns that get used in go are:

- Define methods on types by adding a receiver in argument lists between the func keyword and method name. Functions with receivers 
    are referred to as methods.

- Create interfaces. Interfaces are collections of method signatures or other interfaces. Types that implement everything in the
    interface have implemented the interface and can be referred to using the interface (polymorphism).

- If a type requires initialisation instead of a constructor:
  - Crete a function of the form NewT(<args>) *T
  - Use the built-in new function to return a pointer to the desired type and instantiate everything within that type to zero-values.

See the examples under the 'patterns' directory.

### The IO Interface

The 'io' package defines a number of interface types:
- Reader
- Writer
- Seeker
- Closer



## Static Compilation

A lot of this section came from:

- https://www.arp242.net/static-go.html
- https://mt165.co.uk/blog/static-link-go


By default, Go creates static binaries unless cgo (https://pkg.go.dev/cmd/cgo) is used to call C code which results in a dynamically linked
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

This still results in a working binary as there are pure Go versions that wil be used instead, but they are not totally feature compatible
with the C implementations. They are close however so its usually fine.

If you want to use cgo and have a statically linked binary can do:

    go build -ldflags "-linkmode 'external' -extldflags '-static'"

That command above tells the Go toolchain to use an external linker and gets that linker to produce a static binary. A lot of sites say that
specifying an external linker is not needed.

However, you will usually get some warnings when using glibc (most common implementation of C standard libraries) as glibc itself has dependencies
on other libraries. If your binary doesn't rely on the functions in the warnings, again you should be fine.

You can use musl (https://wiki.musl-libc.org/) to replace glibc which may help solve this. To use musl, you can either install it and build your
software using musl-gcc, or you can use a Linux distribution that uses musl, e.g. Alpine Linux, as a build environment.

In conclusion - relying on a static binary has risks.....
