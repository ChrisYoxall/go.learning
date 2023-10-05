package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	fmt.Println("Starting webserver.")

	// Example from https://www.sohamkamani.com/golang/context/
	// A request will return a response after 3 seconds. If the request is cancelled before that it's detected
	// by the context and a message is written to stdout.
	err := http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		select {
		case <-time.After(3 * time.Second):
			_, err := w.Write([]byte("Request processed."))
			if err != nil {
				log.Printf("Error: %v\n", err)
			}
		case <-ctx.Done():
			_, err := fmt.Fprint(os.Stdout, "Request cancelled.\n")
			if err != nil {
				log.Printf("Error: %v\n", err)
			}
		}
	}))

	if err != nil {
		panic(err)
	}
}
