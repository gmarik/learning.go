// Exported names
// 
// After importing a package, you can refer to the names it exports.
// 
// In Go, a name is exported if it begins with a capital letter.
// 
// Foo is an exported name, as is FOO. The name foo is not exported.
// 
// Run the code. Then rename math.pi to math.Pi and try it again.

package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println(math.Pi)
}
