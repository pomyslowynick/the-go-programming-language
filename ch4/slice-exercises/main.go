package main

import "fmt"

func main() {
	//!+array
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverseWithPointers(&a)
	fmt.Println(a) // "[5 4 3 2 1 0]"
	//!-array

	//!+slice
	s := []int{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions.
	rotate(s)
	fmt.Println(s)

	stringsArr := []string{"Word", "Word", "Word", "newWord", "oldWord", "Word", "newWord"}
	stringArr := removeAdjacent(stringsArr)
	fmt.Println(stringArr)

}

// Exercise 4.3
// https://stackoverflow.com/questions/2439453/using-a-pointer-to-array
// Looks like it might not be possible anymore?
func reverseWithPointers(s *[6]int) {
	sliceA := s[:]
	for i, j := 0, len(sliceA)-1; i < j; i, j = i+1, j-1 {
		sliceA[i], sliceA[j] = sliceA[j], sliceA[i]
	}
}

// Exercise 4.4
func rotate(s []int) {
	sliceS := s[:]
	sliceSFirstHalf := s[:2]
	sliceSSecondHalf := s[2:]
	reverse(sliceS)
	reverse(sliceSSecondHalf)
	reverse(sliceSFirstHalf)
}

// Exercise 4.5
func removeAdjacent(s []string) []string {
	for index, item := range s {
		fmt.Printf("%d index, %s item\n", index, item)
		if index != len(s)-1 && item == s[index+1] {
			copy(s[index:], s[index+1:])
			s = s[:len(s)-1]
		}
	}
	return s
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
