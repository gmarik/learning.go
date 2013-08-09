// Methods with pointer receivers
//
// Methods can be associated with a named type or a pointer to a named type.
//
// We just saw two Abs methods. One on the *Vertex pointer type and the other on the MyFloat value type.
//
// There are two reasons to use a pointer receiver. First, to avoid copying the value on each method call (more efficient if the value type is a large struct). Second, so that the method can modify the value that its receiver points to.
//
// Try changing the declarations of the Abs and Scale methods to use Vertex as the receiver, instead of *Vertex.
//
// The Scale method has no effect when v is a Vertex. Scale mutates v. When v is a value (non-pointer) type, the method sees a copy of the Vertex and cannot mutate the original value.
//
// Abs works either way. It only reads v. It doesn't matter whether it is reading the original value (through a pointer) or a copy of that value.

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	// There are two reasons to use a pointer receiver.
	// First, to avoid copying the value on each method call (more efficient if the value type is a large struct). Second, so that the method can modify the value that its receiver points to.
	v := &Vertex{3, 4}
	v.Scale(5)
	fmt.Println(v, v.Abs())
}
