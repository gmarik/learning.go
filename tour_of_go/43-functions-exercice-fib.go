package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    n1,n2, n := 0, 1, 0

    return func() int {
        n, n1, n2 = n1, n2, n1 + n2

        return n
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
