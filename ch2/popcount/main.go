package main

import (
	"fmt"
	"math/rand"
	"time"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
// Exercise 2.3
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func main() {
	start := time.Now()
	for i := 0; i < 1000000000; i++ {
		PopCount(uint64(rand.Int63()))
	}
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Printf("Time it took: %v\n", elapsed)
}

//!-
