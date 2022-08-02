package main

import "fmt"

// firstly global value assignments are calculated, function call is possible.
// secondly init() functions will be executed top to bottom.
// lastly main() function will be executed.

var j int = i + 1

func createB() bool {
	return true
}

var i int = createI()

func main() {
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(b)
}

func createI() int {
	fmt.Println("In createI")
	//fmt.Printf("%d %d \n", i, j)
	return 5
}

func init() {
	fmt.Println("In init I")
	fmt.Printf("%d %d \n", i, j)
}

func init() {
	fmt.Println("In init II")
	fmt.Printf("%d %d \n", i, j)
}

var b bool = createB()
