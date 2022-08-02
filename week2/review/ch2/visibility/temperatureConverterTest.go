// Run: go mod edit -replace  temperature/converter=./converter

package main

import (
	"fmt"
	"temperature/converter"
)

func main() {

	cDegree := 36.7
	fDegree := converter.Convert('c', 'f', cDegree)
	fmt.Printf("%f celcius is %f fahrenheit.", cDegree, fDegree)

	//fDegree = converter.convertFromCelsiusToFahrenheit(fDegree)
}
