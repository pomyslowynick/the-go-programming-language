package main

import "fmt"

// Exercise 5.19
func main() {
	v := panicker()
	fmt.Println(v)
}

func panicker() (ret interface{}) {
	defer func() {
		ret = recover()
	}()
	panic(2)
}
