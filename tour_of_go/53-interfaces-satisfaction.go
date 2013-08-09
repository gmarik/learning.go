// Interfaces are satisfied implicitly
//
// A type implements an interface by implementing the methods.
//
// There is no explicit declaration of intent.
//
// Implicit interfaces decouple implementation packages from the packages that define the interfaces: neither depends on the other.
//
// It also encourages the definition of precise interfaces, because you don't have to find every implementation and tag it with the new interface name.
//
// Package io defines Reader and Writer; you don't have to.

package main

import (
	"fmt"
	"os"
)

type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

func main() {
	var w Writer

	// os.Stdout implements Writer
	w = os.Stdout

	fmt.Fprintf(w, "hello, writer\n")
}
