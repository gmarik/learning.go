// Exercise: Maps
// 
// Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.
// 
// You might find strings.Fields helpful.

// Prereq: go get "code.google.com/p/go-tour"

package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    m := make(map[string]int)

    for _, w := range strings.Fields(s) {
      m[w] += 1
    }

    return m
}

func main() {
    wc.Test(WordCount)
}

// Running go run .//40-map-exercise.go
// PASS
//  f("I am learning Go!") = 
//   map[string]int{"I":1, "am":1, "learning":1, "Go!":1}
// PASS
//  f("The quick brown fox jumped over the lazy dog.") = 
//   map[string]int{"The":1, "quick":1, "fox":1, "jumped":1, "the":1, "brown":1, "over":1, "lazy":1, "dog.":1}
// PASS
//  f("I ate a donut. Then I ate another donut.") = 
//   map[string]int{"I":2, "ate":2, "a":1, "donut.":2, "Then":1, "another":1}
// PASS
//  f("A man a plan a canal panama.") = 
//   map[string]int{"A":1, "man":1, "a":2, "plan":1, "canal":1, "panama.":1}
