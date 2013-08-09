
// A function can take zero or more arguments.
// 
// In this example, add takes two parameters of type int.
// 
// Notice that the type comes after the variable name.
// 
// (For more about why types look the way they do, see the article on Go's declaration syntax http://golang.org/doc/articles/gos_declaration_syntax.html)

package main

import "fmt"

func add(x int, y int) int {
    return x + y
}

func main() {
    fmt.Println(add(42, 13))
}
