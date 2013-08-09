// Advanced Exercise: Complex cube roots
// 
// Let's explore Go's built-in support for complex numbers via the complex64 and complex128 types. For cube roots, Newton's method amounts to repeating:
// 
// 
// Find the cube root of 2, just to make sure the algorithm works. There is a Pow function in the math/cmplx package.

package main

import "fmt"
import "math/cmplx"

func Cbrt(x complex128) complex128 {

    iter := func(z complex128) complex128 {
      return z - (z * z * z - x) / (3 * z * z)
    }

    isGoodEnough := func(z complex128) bool {
      return cmplx.Abs(z - iter(z)) > 1e-15
    }

    z := x;

    for ; isGoodEnough(z); z = iter(z) {
      fmt.Println(z, cmplx.Abs(z-iter(z)) )
    }

    return z
}

func main() {
    fmt.Println(Cbrt(2))
}
