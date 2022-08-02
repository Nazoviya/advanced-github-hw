package main

import (
	"fmt"
	"strconv"
)

func main() {
	// string to int
	a, _ := strconv.Atoi("29")
	fmt.Printf("%d, %T\n", a, a)

	// string to float64
	b, _ := strconv.ParseFloat("6.3", 64)
	fmt.Printf("%f, %T\n", b, b)

	// int to string
	c := strconv.Itoa(26)
	fmt.Printf("%v, %T\n", c, c)

	// string to int64
	d, _ := strconv.ParseInt("85", 10, 64)
	fmt.Printf("%d, %T\n", d, d)

	// string to []byte
	e := []byte("hey!")
	fmt.Printf("%b, %T\n", e, e)

	// byte to string
	fmt.Printf("%s, %T\n", string(e), string(e))
}
