package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
		fmt.Println(s)
	}

	// Exercise 1.1
	// fmt.Println(strings.Join(os.Args[0:], " "))

	// Exercise 1.2
	// for index, arg := range os.Args[0:] {
	// 	// s := fmt.Sprint(index) + " " + arg
	// 	s := []string{fmt.Sprint(index), arg}
	// 	fmt.Println(s)
	// }
}
