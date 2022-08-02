package main

import (
	"bufio"
	"fmt"
	"os"
)

// global variable s.
var s string = "I love Go!"

func main() {
	fmt.Printf("%s\n", s)

	// local variable s that belongs to main().
	var s string = "I really love Go!"
	fmt.Printf("%s\n", s)

	//var s string = "I said you, I really love Go!"
	//fmt.Printf("%s\n", s)
	//x := "hellooww"
	for i, s := range s { // That's so strange!
		// local variable s that belongs to for loop.
		s := s
		fmt.Printf("%d:  %c\n", i, s)
	}

	fmt.Printf("%s\n", s)

	fmt.Println("*****************")

	length := len(s)
	fmt.Printf("%d\n", length)

	if length := f(); length > 0 {
		fmt.Printf("In if %d\n", length)
	} else {
		fmt.Printf("In else %d\n", length)
	}
	fmt.Printf("%d\n", length)

	fmt.Println("*****************")

	fmt.Println(g1())
	fmt.Println(g2())
}

func f() int {
	return 20
}

// GOPL p. 48
func g1() error {
	//if f, err := os.Open("scope.go"); err != nil { // compile error: unused: f
	//	return err
	//}
	//f.Stat() // compile error: undefined f
	//f.Close()    // compile error: undefined f
	return nil
}

func g2() error {
	f, err := os.Open("scope/read_example.txt")
	if err != nil {
		return err
	}

	b := bufio.NewReader(f)
	c, err := b.Peek(b.Size())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(c))
	f.Stat()  // compile error: undefined f
	f.Close() // compile error: undefined f
	return nil
}
