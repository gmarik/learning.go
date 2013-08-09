// Multiple results
// 
// A function can return any number of results.
// 
// This function returns two strings.


package main

import "fmt"

func swap(x, y string) (string, string) {
    return y, x
}

func main() {
    a, b := swap("hello", "world")
    fmt.Println(a, b)
}
