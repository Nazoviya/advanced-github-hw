package main

import "fmt"

func main() {
	const (
		Jan = 1 + iota // 1 + 0 = 1
		Feb            // 1 + 1 = 2
		Mar            // 1 + 2 = 3
		Apr            // 1 + 3 = 4
		May            // ..
		Jun            // .
		Jul
		Aug
		Sep
		Oct
		Nov
		Dec
	)

	fmt.Printf("%v, %v, %v, %v, %v, %v", Feb, Mar, Jun, Aug, Sep, Dec)
}
