package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	// Instead of using this syntax for creating a slice, will use make
	//bs := []byte{}

	bs := make([]byte, 99999)
	numBytes, _ := resp.Body.Read(bs)

	fmt.Println(string(bs))
	fmt.Println("Bytes returned:", numBytes)

	// NOTE: Above code was used to demonstrate interfaces and relies on declaring a byte slice
	// that is large enough to hold the response. If its too small not all response is printed.
	// Instead do this to print the body
	// bodyBytes, _ := io.ReadAll(resp.Body)
	// fmt.Println(string(bodyBytes))
	//
	// Or can do:
	// io.Copy(os.Stdout, resp.Body)

}
