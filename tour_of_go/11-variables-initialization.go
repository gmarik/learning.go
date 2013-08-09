
// Variables with initializers
// 
// A var declaration can include initializers, one per variable.
// 
// If an initializer is present, the type can be omitted; the variable will take the type of the initializer.


package main

import "fmt"

var x, y, z int = 1, 2, 3
var c, python, java = true, false, "no!"

func main() {
    fmt.Println(x, y, z, c, python, java)
}
