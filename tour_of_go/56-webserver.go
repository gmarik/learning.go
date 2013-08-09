// Web servers
//
// Package http serves HTTP requests using any value that implements http.Handler:
//
// package http
//
// type Handler interface {
//     ServeHTTP(w ResponseWriter, r *Request)
// }
// In this example, the type Hello implements http.Handler.
//
// Visit http://localhost:4000/ to see the greeting.
//
// Note: This example won't run through the web-based tour user interface. To try writing web servers you may want to Install Go.

package main

import (
	"fmt"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func main() {
	var h Hello
	http.ListenAndServe("localhost:4000", h)
}
