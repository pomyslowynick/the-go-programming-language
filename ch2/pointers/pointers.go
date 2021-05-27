package main

import "fmt"

func main() {
	var p = f()
	var p2 = f()
	fmt.Println(p)
	fmt.Println(p2)
}
func f() *int {
	v := 1
	return &v
}
