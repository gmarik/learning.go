// For
// 
// Go has only one looping construct, the for loop.
// 
// The basic for loop looks as it does in C or Java, except that the ( ) are gone (they are not even optional) and the { } are required.


package main

import "fmt"

func main() {
    sum := 0
    for i := 0; i < 10; i++ {
        sum += i
    }
    fmt.Println(sum)
}
