// Slicing slices
// 
// Slices can be re-sliced, creating a new slice value that points to the same array.
// 
// The expression
// 
// s[lo:hi]
// evaluates to a slice of the elements from lo through hi-1, inclusive. Thus
// 
// s[lo:lo]
// is empty and
// 
// s[lo:lo+1]
// has one element.

package main

import "fmt"

func main() {
    p := []int{2, 3, 5, 7, 11, 13}
    fmt.Println("p ==", p)
    fmt.Println("p[1:4] ==", p[1:4])

    // missing low index implies 0
    fmt.Println("p[:3] ==", p[:3])

    // missing high index implies len(s)
    fmt.Println("p[4:] ==", p[4:])
}
