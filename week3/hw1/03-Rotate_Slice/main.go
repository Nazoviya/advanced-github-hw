package main

import "fmt"

func main() {
	rotation, err := getRotation()
	if err != nil {
		return
	}

	// declaration of slice that will be rotated.
	realS := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("\n-------Rotate to left--------")
	fmt.Printf("input:  %v\n", realS)
	fmt.Printf("output: %v\n", leftRotation(realS, len(realS), rotation))

	fmt.Println("\n-------Rotate to right-------")
	fmt.Printf("input:  %v\n", realS)
	fmt.Printf("output: %v\n", rightRotation(realS, len(realS), rotation))
}

func getRotation() (rotation int, err error) {
	// get rotation step from user
	fmt.Print("Give a rotation step to see rotated slice: ")
	fmt.Scanln(&rotation)
	return rotation, err
}

// rotate slice to the left with taken step
func leftRotation(realS []int, size int, rotation int) []int {
	var tempS []int
	for i := 0; i < rotation; i++ {
		tempS = realS[1:size]
		tempS = append(tempS, realS[0])
		realS = tempS
	}
	return realS
}

// rotate slice to the right with taken step
func rightRotation(realS []int, size int, rotation int) []int {
	var tempS []int
	for i := 0; i < rotation; i++ {
		tempS = append([]int{realS[size-1]}, realS[0:size-1]...)
		realS = tempS
		tempS = nil
	}
	return realS
}
