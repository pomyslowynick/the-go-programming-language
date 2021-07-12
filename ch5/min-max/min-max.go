package main

import (
	"fmt"
	"math"
	"os"
)

// Exercise 5.15
func main() {
	max, ok := max([]int{1, 2, 3, 4, 12, 4, 12, 421, 12, 3, 21, 321, 412}...)
	if !ok {
		fmt.Fprintf(os.Stderr, "max wasn't given at least one argument")
	}
	fmt.Println(max)

	min, ok := min([]int{1, 2, 3, 4, 12, 4, 12, 421, 12, 3, 21, 321, 412, -234}...)
	if !ok {
		fmt.Fprintf(os.Stderr, "min wasn't given at least one argument")
	}
	fmt.Println(min)
}

func max(numbers ...int) (int, bool) {
	if len(numbers) == 0 {
		return -1, false
	}
	tempMax := math.MinInt32
	for _, number := range numbers {
		if number > tempMax {
			tempMax = number
		}
	}
	return tempMax, true
}

func min(numbers ...int) (int, bool) {
	if len(numbers) == 0 {
		return -1, false
	}
	tempMin := math.MaxUint32
	for _, number := range numbers {
		if number < tempMin {
			tempMin = number
		}
	}
	return tempMin, true
}
