package main

import "fmt"

func main() {
	a, b := 10, 5.5
	fmt.Println(float64(a) /* int to float */ + b)

	a, b = 10, 5.5
	a = int(b) // float to int
	fmt.Println(float64(a) /* int to float */ + b)

	age := 2
	fmt.Println(7.5 + float32(age) /* int to float32 */)

	min := int8(127)
	max := int16(1000)
	fmt.Println(max + int16(min) /* int8 to int16 */)
}
