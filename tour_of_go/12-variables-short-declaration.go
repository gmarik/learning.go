// Short variable declarations
// 
// Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
// 
// Outside a function, every construct begins with a keyword (var, func, and so on) and the := construct is not available.

package main

import "fmt"

func main() {
    var x, y, z int = 1, 2, 3
    c, python, java := true, false, "no!"

    fmt.Println(x, y, z, c, python, java)
}
