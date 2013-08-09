// Nil slices
// 
// The zero value of a slice is nil.
// 
// A nil slice has a length and capacity of 0.
// 
// (To learn more about slices, read the Slices: usage and internals article http://golang.org/doc/articles/slices_usage_and_internals.html.)

package main

import "fmt"

func main() {
    var z []int
    fmt.Println(z, len(z), cap(z))
    if z == nil {
        fmt.Println("nil!")
    }
}
