// For continued
//
// As in C or Java, you can leave the pre and post statements empty.


package main

import "fmt"

func main() {
    sum := 1
    for ; sum < 1000; {
        sum += sum
    }
    fmt.Println(sum)
}
