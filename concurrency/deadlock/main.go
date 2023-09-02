/*

Running this will result in "fatal error: all goroutines are asleep - deadlock!"

Remember that both reading and writing to a channel are blocking operations in Go. When a goroutine is blocked
control will pass to another goroutine in the hope of being able to do some work, but if all goroutines are blocked
then you will end up with a deadlock.

In this example the main goroutine is blocked on when a value is sent to the channel, and as are no other goroutines
available to read the value it is effectively a deadlock.

*/

package main

import "fmt"

func main() {

	c := make(chan string)
	c <- "Chris"
	fmt.Println(<-c)
	close(c)

}
