
// If
// 
// The if statement looks as it does in C or Java, except that the ( ) are gone and the { } are required.
// 
// (Sound familiar?)


package main

import (
    "fmt"
    "math"
)

func sqrt(x float64) string {
    if x < 0 {
        return sqrt(-x) + "i"
    }
    return fmt.Sprint(math.Sqrt(x))
}

func main() {
    fmt.Println(sqrt(2), sqrt(-4))
}
