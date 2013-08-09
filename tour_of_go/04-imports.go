
// Imports
// 
// This code groups the imports into a parenthesized, "factored" import statement. You can also write multiple import statements, like:
// 
// import "fmt"
// import "math"


package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Printf("Now you have %g problems.\n", math.Nextafter(2, 3))
}
