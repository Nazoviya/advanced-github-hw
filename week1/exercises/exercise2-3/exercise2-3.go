/*
Rewrite PopCount to use a loop instead of a single expression. Compare the
performance of the two versions. (Section 11.4 shows how to compare the
performance of different implementations systematically.)
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
	for i := 0; i < 8; i++ {
		ret += int(pc[byte(x>>uint(i*8))])
	}
	return
}

func main() {
	fmt.Println(PopCount(10)) // 8(2^3) + 2(2^1) = 10, so 2 set bits.
}
