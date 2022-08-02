/*
The expression x&(x-1) clears the rightmost non-zero bit of x. Write a version
of PopCountthat counts bits by using this fact, and assess its performance.
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
	for x != 0 {
		x = x & (x - 1)
		ret++
	}
	return
}

func main() {
	fmt.Println(PopCount(10)) // 8(2^3) + 2(2^1) = 10, so 2 set bits.
}
