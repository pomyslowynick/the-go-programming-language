package main

import "fmt"

func namedResult() (output string) {
	output = "Hello"
	return output
}

func main() {
	fmt.Println(namedResult())
}
