package main

import "fmt"

// Exercise 5.19
func main() {
	v := panicker()
	fmt.Println(v)
}

// The exercise said that no return statement can be used. Since the function that was panicking returns
// normally so the declared value is returned and in main non zero value is returned.
func panicker() (ret interface{}) {
	defer func() {
		ret = recover()
	}()
	panic(2)
}
