package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(nonRecursiveComma(input.Text()))
	}
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

// Exercise 3.10
// TODO: Issue when number is not divisible by 3.
func nonRecursiveComma(s string) string {
	var buf bytes.Buffer
	for i, v := range s {
		if i%3 == 0 && i != 0 {
			buf.WriteString(",")
		}
		buf.WriteRune(v)
	}
	return buf.String()
}
