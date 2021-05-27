package main

import "fmt"

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
func fib(n int) int {
	var x, y = 0, 1
	for n != 0 {
		x, y = y, x+y
		n--
	}
	return x
}

func main() {
	fmt.Println(fib(10))
}
