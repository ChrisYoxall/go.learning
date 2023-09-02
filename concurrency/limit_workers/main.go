/* This is an attempt to process a queue with a limited number of workers, where the number of workers
is increased or decreased over time. Seems to work, however if it was a remote queue that was being
processed I wouldn't want to do this as the workers check the length of the queue every iteration.

Most solutions online seem to use a WaitGroup to start batches of workers.
*/

package main

import (
	"fmt"
	"runtime"
	"time"
)

const QueueLength = 50
const QueueTriggerLevel = 10
const MaxWorkers = 15

func writeToQueue(c chan int) {

	id := 0

	// A buffered channel can be used like a semaphore. In this case it limits the number of items in the queue
	// and when full adding another value to the queue will block until there is space to add the value.
	for {
		time.Sleep(250 * time.Millisecond)
		c <- id
		id++
	}

}

func readFromQueue(c chan int) {

	for {
		id := <-c

		time.Sleep(2 * time.Second) // Simulate processing the item.

		fmt.Printf("Finished processing message %d.\n", id)

		if len(c) < QueueTriggerLevel {
			fmt.Printf("Worker is quitting.\n")
			break
		}
	}
}

func main() {

	// Create a buffered channel. The channel will hold the number of values specified by its size at which point it
	// will start blocking if another value is added. Can only get the number of items in a channel if it is buffered.
	c := make(chan int, QueueLength)

	go writeToQueue(c)

	for {
		time.Sleep(2 * time.Second)

		queueLength := len(c)
		fmt.Printf("There are %d goroutines running and %d items in the queue.\n", runtime.NumGoroutine(), queueLength)

		if queueLength > QueueTriggerLevel && runtime.NumGoroutine() < MaxWorkers {
			fmt.Printf("Starting new worker.\n")
			go readFromQueue(c)
		}
	}
}
