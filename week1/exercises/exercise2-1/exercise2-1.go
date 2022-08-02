/*
Add types, constants, and functions to tempconv for processing temperatures in
the Kelvin scale, where zero Kelvin is −273.15°C and a difference of 1K has the
same magnitude as 1°C.
*/

// Package tempconv performs Celsius and Fahrenheit conversions.
// package tempconv
package main

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }

func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

func printKToC() {
	degree := 10
	fmt.Printf("%s is %s", Kelvin(degree), KToC(Kelvin(degree)))
}

func main() {
	printKToC()
}
