// Named results
// 
// Functions take parameters. In Go, functions can return multiple "result parameters", not just a single value. They can be named and act just like variables.
// 
// If the result parameters are named, a return statement without arguments returns the current values of the results.


package main

import "fmt"

func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}

func main() {
    fmt.Println(split(17))
}
