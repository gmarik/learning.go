// Channels
//
// Channels are a typed conduit through which you can send and receive values with the channel operator, <-.
//
// ch <- v    // Send v to channel ch.
// v := <-ch  // Receive from ch, and
//            // assign value to v.
// (The data flows in the direction of the arrow.)
//
// Like maps and slices, channels must be created before use:
//
// ch := make(chan int)
// By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.

package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
