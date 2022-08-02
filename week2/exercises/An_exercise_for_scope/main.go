package main

import (
	"fmt"
)

// global variable s.
var s string = "I love Go! -> global"

func main() {
	fmt.Printf("%s\n", s)

	// local variable s that belongs to main().
	var s string = "I love Go! -> local"
	fmt.Printf("%s\n", s)

	for i, s := range s {
		// local variable s that belongs to for loop.
		s := s
		fmt.Printf("%d:  %c\n", i, s)
	}
}
