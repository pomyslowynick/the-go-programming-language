package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(isAnagram(os.Args[1], os.Args[2]))
}

// Exercise 3.12
func isAnagram(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for _, v := range s1 {
		if strings.IndexRune(s2, v) == -1 {
			return false
		}
	}

	return true
}
