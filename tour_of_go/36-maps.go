// Maps
// 
// A map maps keys to values.
// 
// Maps must be created with make (not new) before use; the nil map is empty and cannot be assigned to.package main

package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}

var m map[string]Vertex

func main() {
    m = make(map[string]Vertex)
    m["Bell Labs"] = Vertex{
        40.68433, -74.39967,
    }
    fmt.Println(m["Bell Labs"])
    fmt.Println(m)
}
