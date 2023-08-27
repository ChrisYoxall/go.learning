package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://amazon.com",
		"https://youtube.com",
		"https://wikipedia.org",
		"https://go.dev",
	}

	// Create channel.
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// Waiting for message from channel is a blocking call. Note how range call here uses
	// normal syntax, but not integrating over a known list - rather iterating on each response
	// from channel.
	for link := range c {

		// Example of function literal (similar to anonymous functions).
		go func(l string) {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}(link) // Extra () is to call the function as required by go keyword.

	}

}

func checkLink(link string, c chan string) {

	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down.")
		c <- link
		return
	}

	fmt.Println(link, "is up.")
	c <- link
}
