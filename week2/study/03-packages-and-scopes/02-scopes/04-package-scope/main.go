package main

import "fmt"

func main() {
	fmt.Println("Hello!")

	// two files belong to the same package
	// calling `bye()` of bye.go here
	hey()
	bye()
}
