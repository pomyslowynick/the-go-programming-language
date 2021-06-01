package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type Flags uint

const (
	FlagUp Flags = 1 << iota
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

func main() {
	fmt.Printf("%d %d %d %d %d\n", FlagUp, FlagBroadcast, FlagLoopback, FlagPointToPoint, FlagMulticast)
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
// Exercise 3.11
// TODO: Issue when number is not divisible by 3.
func nonRecursiveComma(s string) string {
	var buf bytes.Buffer
	dotFlag := false
	for i, v := range s {
		if v == '.' {
			dotFlag = true
		}

		if i%3 == 0 && i != 0 && !dotFlag {
			buf.WriteString(",")
		}
		buf.WriteRune(v)
	}
	return buf.String()
}
