
// Exercise: Loops and Functions
// 
// As a simple way to play with functions and loops, implement the square root function using Newton's method.
// 
// In this case, Newton's method is to approximate Sqrt(x) by picking a starting point z and then repeating:
// 
// z = z - (z^2 - x)/2z
// 
// To begin with, just repeat that calculation 10 times and see how close you get to the answer for various values (1, 2, 3, ...).
// 
// Next, change the loop condition to stop once the value has stopped changing (or only changes by a very small delta). See if that's more or fewer iterations. How close are you to the math.Sqrt?
// 
// Hint: to declare and initialize a floating point value, give it floating point syntax or use a conversion:
// 
// z := float64(1)
// z := 1.0


package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64, prec float64) float64 {
    var z = x

    var next = func(x, z float64) float64 {
      return z - (math.Pow(z, 2) - x) / (2 * z)
    }

    for zz:= next(x, z); math.Abs(zz-z) > prec; {
      z = zz
      zz = next(x, z)
    }

    return z
}

func main() {
    var x = 2.0
    var r = Sqrt(x, 10e-13)
    var rr = math.Sqrt(x)
    fmt.Println(r, rr, rr - r)
}
