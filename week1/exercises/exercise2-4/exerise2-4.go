/*
Write a version of PopCount that counts bits by shifting its argument through
64 bit positions, testing the right most bit each time. Compare its performance
to the table-lookup version.
*/

package main

import "fmt"

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) (ret int) {
	for i := 0; i < 64; i++ {
		if x&1 == 1 {
			ret++
		}
		x >>= 1
	}
	return ret
}

func main() {
	fmt.Println(PopCount(10)) // 8(2^3) + 2(2^1) = 10, so 2 set bits.
}
