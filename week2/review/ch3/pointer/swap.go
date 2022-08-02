package main

import "fmt"

func main() {
	x := 5
	y := 7
	fmt.Printf("x = %d y = %d\n", x, y)

	x, y = swap1(x, y)
	fmt.Printf("x = %d y = %d  -> swap1\n", x, y)

	x = 5
	y = 7
	swap2(x, y)
	fmt.Printf("x = %d y = %d  -> swap2\n", x, y)

	x = 5
	y = 7
	swap3(&x, &y)
	fmt.Printf("x = %d y = %d  -> swap3\n", x, y)
}

// real x and y values are affacted, because x and y values
// returned to be assigned.
func swap1(x, y int) (int, int) {
	temp := x
	x = y
	y = temp
	return x, y
}

// in this form it makes copy of the x and y variables
// so, real x and y values did not change.
func swap2(x, y int) {
	temp := x
	x = y
	y = temp
}

// real x and y values are affacted, because the values on
// specific memory addresses are changed via using pointers.
func swap3(x, y *int) {
	temp := *x
	*x = *y
	*y = temp
}
