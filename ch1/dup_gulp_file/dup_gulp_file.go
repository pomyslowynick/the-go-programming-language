package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	chars := make(map[string]int)
	for _, arg := range os.Args[1:] {
		data, err := ioutil.ReadFile(arg)

		if err != nil {
			fmt.Fprintf(os.Stderr, "oopsie doo %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			chars[line] += 1
		}
	}

	for line, num := range chars {
		if num > 1 {
			fmt.Printf("%d \t%s\n", num, line)
		}
	}
}
