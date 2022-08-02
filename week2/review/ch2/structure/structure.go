// That's the main package.
package main

import (
	"errors"
	"fmt"
	"log"
)

//const lowestC = -273.0
//const lowestF = -460.0

// These are the constants representing the lowest possible degrees.
const (
	lowestC = -273.0
	lowestF = -460.0
)

// C2F Converts from celsius to fahrenheit
func C2F(cDegree float64) (float64, error) {
	if cDegree < lowestC {
		return 0, errors.New("Invalid degree passed: " + fmt.Sprintf("%f", cDegree))
	} else {
		return (9 * cDegree / 5) + 32, nil
	}
}

// Converts from fahrenheit to celsius.
func f2C(fDegree float64) (float64, error) {
	if fDegree < lowestF {
		return 0, errors.New("invalid degree passed: " + fmt.Sprintf("%f", fDegree))
	}
	return (fDegree - 32) * 5 / 9, nil
}

/*
That's a longer comment.
Using this format makes writing longer comments easier.
*/
func main() {
	fmt.Println(desc)
	cDegree := -40.0 // Try -340
	fDegree, err := C2F(cDegree)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%f celcius is %f fahrenheit\n", cDegree, fDegree)
	}

	fDegree = 98.0 // Try -498
	cDegree, _ = f2C(fDegree)
	fmt.Printf("%f fahrenheit is %f celcius", fDegree, cDegree)
}

var desc string = "This is a Celsius-Fahrenheit converter."
