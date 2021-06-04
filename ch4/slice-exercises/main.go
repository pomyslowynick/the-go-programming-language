package main

import "fmt"

func main() {
	//!+array
	a := [...]int{0, 1, 2, 3, 4, 5}
	sliceA := a[:]
	reverseWithPointers(&sliceA)
	fmt.Println(a) // "[5 4 3 2 1 0]"
	//!-array

	//!+slice
	s := []int{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions.
	sliceS := s[:]
	sliceSFirstHalf := s[:2]
	sliceSSecondHalf := s[2:]
	reverseWithPointers(&sliceSFirstHalf)
	reverseWithPointers(&sliceSSecondHalf)
	reverseWithPointers(&sliceS)
	fmt.Println(s) // "[2 3 4 5 0 1]"

}

// Exercise 4.3
func reverseWithPointers(s *[]int) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}
