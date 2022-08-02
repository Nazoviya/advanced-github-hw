package main

import (
	"fmt"
	"math"
)

func main() {
	pi := 3.14
	var i int = int(pi)
	fmt.Println(i)

	// i = int(6.28)
	// Cannot convert an expression of the type 'float64' to the type 'int'

	i = int(math.Floor(6.28))
	fmt.Println(i)

	type yazi string
	var y yazi = "I like Go!"

	type text string
	var t text = "I love Go!"

	type exValue yazi
	var s3 exValue = "It is also possible!"

	var s1 string = string(y)
	var s2 string = string(t)

	fmt.Println(s1, s2, s3)
}
