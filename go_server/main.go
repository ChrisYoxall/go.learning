// Simple web server. For anything more than a few responses there are plenty
// of web frameworks to use such as Gin or Fiber.

// Note that Go 1.22 introduced enhanced routing capabilities to the ServeMux handler.
// Refer https://go.dev/blog/routing-enhancements

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Hello World!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "ParseForm() error: %v", err)
	}

	fmt.Fprintf(w, "POST request successful\n")

	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "Address: %s\n", address)

}

func main() {

	// Using the default mux (i.e. specify nil as the second argument to ListenAndServe which is often seen in simple
	// examples) is a security concern as the default 'ServeMux' is stored in a global variable where any third party
	// package can access it and register a route (i.e. a handlers). Instead, create a locally scoped ServeMux.
	mux := http.NewServeMux()
	mux.Handle("GET /", http.FileServer(http.Dir("./static"))) // called via http://localhost:8080 which will serve index.html by default
	mux.HandleFunc("POST /form", formHandler)                  // called when submitting the form  http://localhost:8080/form.html
	mux.HandleFunc("GET /hello", helloHandler)                 // called via http://localhost:8080/hello

	port := 8080

	// Could now start a server by doing 'http.ListenAndServe' as per the next commented out block. However, it is
	// generally recommended to use the http.Server struct which provides more control over the server, and configure
	// some time-outs to prevent slowloris attacks. Refer  https://pkg.go.dev/net/http#hdr-Servers

	//fmt.Printf("Starting webserver on port %v\n", port)
	//err := http.ListenAndServe(":"+port, mux)
	//if err != nil {
	//	log.Fatal(err)
	//}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	fmt.Printf("Starting webserver at %v\n", srv.Addr)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Error starting server: %v\n\n", err)
	}
}
