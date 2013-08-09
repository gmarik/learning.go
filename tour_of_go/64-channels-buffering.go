// Buffered Channels
//
// Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel:
//
// ch := make(chan int, 100)
// Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.
//
// Modify the example to overfill the buffer and see what happens.

package main

import "fmt"

func main() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	// c <- 3
	fmt.Println(<-c)
	fmt.Println(<-c)
}
