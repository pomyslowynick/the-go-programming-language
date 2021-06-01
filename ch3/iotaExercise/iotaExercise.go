package main

import (
	"fmt"
)

type ByteSize uint

const (
	KB = ByteSize(iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func (b ByteSize) String() string {
	v := exponent(b, 1000)
	return v
}

// Exercise 3.13
func main() {
	fmt.Printf("%s\n", GB)

}

// Return exponent of ten for given n
func exponent(n ByteSize, exp int) string {
	if n == 0 {
		return fmt.Sprint(exp)
	}
	return exponent(n-1, exp*1000)
}
