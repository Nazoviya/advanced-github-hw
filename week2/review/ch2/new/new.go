package main

import "fmt"

func main() {
	// new is a built-in function that allocates memory with zero value.
	p := new(int)
	fmt.Printf("%d %d\n", p, *p)

	q := new(int)
	fmt.Printf("%d %d\n", q, *q)

	var r int
	fmt.Printf("%d %d\n", r, &r)

	fmt.Printf("%t\n", p == q)
	fmt.Printf("%t\n", *p == *q)
}
