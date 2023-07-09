package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		log.Println("Need to provide filename!")
		os.Exit(1)
	}

	// Open the first argument as a file.
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// Create a scanner to read from the file.
	scanner := bufio.NewScanner(file)

	// Change the split function to split on words instead of lines.
	scanner.Split(bufio.ScanWords)

	// Counter to track the running total.
	var wordCount int

	// Add 1 to the count for each word.
	for scanner.Scan() {
		wordCount++
	}

	if scanner.Err() != nil {
		log.Println(scanner.Err())
		os.Exit(1)
	}

	fmt.Println("Found", wordCount, "words.")
}
