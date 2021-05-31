// Basename1 reads file names from stdin and prints the base name of each one.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(basename_2(input.Text()))
	}
	// NOTE: ignoring potential errors from input.Err()
}

// basename removes directory components and a .suffix.
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
func basename(s string) string {
	// Discard last '/' and everything before.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// Preserve everything before last '.'.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func basename_2(s string) string {
	lastSlash := strings.LastIndex(s, "/")
	cutS := s[lastSlash:]
	if lastDot := strings.LastIndex(cutS, "."); lastDot >= 0 {
		s = cutS[:lastDot]
	}
	return s
}
