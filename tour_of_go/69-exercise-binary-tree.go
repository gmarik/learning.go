// Exercise: Equivalent Binary Trees
//
// 1. Implement the Walk function.
//
// 2. Test the Walk function.
//
// The function tree.New(k) constructs a randomly-structured binary tree holding the values k, 2k, 3k, ..., 10k.
//
// Create a new channel ch and kick off the walker:
//
// go Walk(tree.New(1), ch)
// Then read and print 10 values from the channel. It should be the numbers 1, 2, 3, ..., 10.
//
// 3. Implement the Same function using Walk to determine whether t1 and t2 store the same values.
//
// 4. Test the Same function.
//
// Same(tree.New(1), tree.New(1)) should return true, and Same(tree.New(1), tree.New(2)) should return false.

package main

import "code.google.com/p/go-tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)

	// requred "forward declaration" for recurive calls
	var walk func(t *tree.Tree)

	walk = func(t *tree.Tree) {
		if t == nil {
			return
		}
		walk(t.Left)
		ch <- t.Value
		walk(t.Right)
	}

	walk(t)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1, c2 := make(chan int), make(chan int)

	go Walk(t1, c1)
	go Walk(t2, c2)

	for {
		v1, cont1 := <-c1
		v2, cont2 := <-c2

		fmt.Println(v1, v2)

		if !(v1 == v2 && cont1 == cont2) {
			return false
		} else if !cont1 {
			return true
		}
	}

}

func main() {
	// c1 := make(chan int)
	// go Walk(tree.New(2), c1)

	// for v := range c1 {
	// 	fmt.Println(v)
	// }

	// select {
	// case v := <-c1:
	// 	fmt.Print(v)
	// }

	t1 := tree.New(1)
	t2 := tree.New(1)
	t3 := tree.New(3)

	// different trees but store same sequence
	fmt.Println(Same(t1, t2))
	fmt.Println(Same(t1, t3))
	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(t3)

}
