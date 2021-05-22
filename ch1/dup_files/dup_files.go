package main

import (
	"bufio"
	"fmt"
	"os"
)

// func main() {
// 	counts := make(map[string]int)
// 	files := os.Args[1:]
// 	if len(files) == 0 {
// 		countLines(os.Stdin, counts)
// 	} else {
// 		for _, arg := range files {
// 			f, err := os.Open(arg)
// 			if err != nil {
// 				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
// 				continue
// 			}
// 			countLines(f, counts)
// 			f.Close()
// 		}
// 	}
// 	for line, n := range counts {
// 		if n > 1 {
// 			fmt.Printf("%d\t%s\n", n, line)
// 		}
// 	}
// }

// func countLines(f *os.File, counts map[string]int) {
// 	input := bufio.NewScanner(f)
// 	for input.Scan() {
// 		counts[]++
// 	}
// }

// Exercise 1.4
func main() {
	counts := make(map[string]int)
	file_appearances := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, file_appearances)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, file_appearances)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, file_appearances[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, file_appearances map[string]string) {
	input := bufio.NewScanner(f)
	has_appeared := make(map[string]bool)
	for input.Scan() {
		counts[input.Text()]++
		if !has_appeared[input.Text()] {
			file_appearances[input.Text()] += f.Name() + " "
			has_appeared[input.Text()] = true
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}
