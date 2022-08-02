package main

import "fmt"

func main() {
	number := getNumber()
	primeFactors(number)
}

func getNumber() int {
	// get input from user
	fmt.Print("Give a number to find prime factors: ")
	var number int
	fmt.Scanln(&number)
	if number < 2 {
		fmt.Println("Prime numbers are bigger than or equal to two.")
	}
	return number
}

// prime factors of taken number
func primeFactors(n int) {
	c := 2
	for n > 1 {
		if n%c == 0 {
			fmt.Print(c, " ")
			n /= c
		} else {
			c++
		}
	}
}
