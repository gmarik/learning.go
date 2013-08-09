// Errors
//
// An error is anything that can describe itself as an error string. The idea is captured by the predefined, built-in interface type, error, with its single method, Error, returning a string:
//
// type error interface {
//     Error() string
// }
// The fmt package's various print routines automatically know to call the method when asked to print an error.

package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

// define Error() method so MyError type now is_a error type
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
