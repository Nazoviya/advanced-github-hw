/*
Write a function that reports whether two strings are anagrams of each other,
that is, they contain the same letters in a different order.
*/

package main

import "fmt"

func main() {
	fmt.Println(anagram("dictionary", "indicatory"))
}

func anagram(s1, s2 string) bool {
	m := make(map[rune]int)
	for _, r := range s1 {
		m[r]++
	}
	for _, r := range s2 {
		if i, ok := m[r]; ok {
			if i > 1 {
				m[r]--
			} else {
				delete(m, r)
			}
		} else {
			return false
		}
	}
	return len(m) == 0
}
