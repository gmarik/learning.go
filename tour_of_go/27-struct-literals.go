// Struct Literals
// 
// A struct literal denotes a newly allocated struct value by listing the values of its fields.
// 
// You can list just a subset of fields by using the Name: syntax. (And the order of named fields is irrelevant.)
// 
// The special prefix & constructs a pointer to a newly allocated struct.package main

package main

import "fmt"

type Vertex struct {
    X, Y int
}

var (
    p = Vertex{1, 2}  // has type Vertex
    q = &Vertex{1, 2} // has type *Vertex
    r = Vertex{X: 1}  // Y:0 is implicit
    s = Vertex{}      // X:0 and Y:0
)

func main() {
    fmt.Println(p, q, r, s)
}
