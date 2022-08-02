package main

import (
	"fmt"
	"math"
)

func main() {
	number := getNumber()
	printer(number, sieveOfEratosthenes(number))
}

func getNumber() int {
	// get input from user
	fmt.Print("Give a number to see all primes up to: ")
	var number int
	fmt.Scanln(&number)
	if number < 2 {
		fmt.Println("Prime numbers are bigger than or equal to two.")
	}
	return number
}

func sieveOfEratosthenes(number int) (allNumbers map[int]bool) {
	// create map
	allNumbers = make(map[int]bool)

	// declare zero values
	for k := 2; k <= number; k++ {
		allNumbers[k] = false
	}

	// check if divisible
	k := 2
	for i := 2; i <= number; i++ {
		if !allNumbers[k] {
			squareOfk := int(math.Pow(float64(k), 2))
			for j := squareOfk; j <= number; j++ {
				if j%i == 0 {
					allNumbers[j] = true
				}
			}
		}
		k++
	}
	return allNumbers
}

func printer(number int, allNumbers map[int]bool) {
	// print primes
	for i := 2; i <= number; i++ {
		if !allNumbers[i] {
			fmt.Print(i, " ")
		}
	}
}
