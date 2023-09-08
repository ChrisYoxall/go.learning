package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

const QueueLength = 50
const MaxWorkers = 15

type message struct {
	id int
}

func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func writeToQueue(c chan message) {

	id := 0

	// A buffered channel can be used like a semaphore. In this case it limits the number of items in the queue
	// and when full adding another value to the queue will block until there is space to add the value.
	for {
		r := randInt(200, 500)
		time.Sleep(time.Duration(r) * time.Millisecond)
		c <- message{id: id}
		id++
	}

}

func processMessage(m message, wg *sync.WaitGroup) {

	defer wg.Done()

	// simulate doing some work
	r := randInt(2, 5)
	time.Sleep(time.Duration(r) * time.Second)

	fmt.Printf("Finished processing message %d.\n", m.id)

}

func main() {

	// Create a buffered channel. The channel will hold the number of values specified by its size at which point it
	// will start blocking if another value is added. Only buffered channels will have a length.
	c := make(chan message, QueueLength)

	go writeToQueue(c)

	fmt.Println("Sleeping while some items get added to the queue..")
	time.Sleep(2 * time.Second)

	for {
		queueLength := len(c)
		fmt.Printf("There are %d goroutines running and %d items in the queue.\n", runtime.NumGoroutine(), queueLength)

		wg := sync.WaitGroup{}

		var requiredWorkers int

		if queueLength >= MaxWorkers {
			requiredWorkers = MaxWorkers
		} else {
			requiredWorkers = queueLength
		}

		fmt.Printf("Starting %d workers.\n", requiredWorkers)
		for i := 0; i < requiredWorkers; i++ {
			wg.Add(1)
			go processMessage(<-c, &wg)
		}

		fmt.Printf("There are %d goroutines running. Waiting for workers to finish...\n", runtime.NumGoroutine())

		wg.Wait()
	}
}
