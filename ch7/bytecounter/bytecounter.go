// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 173.

// Bytecounter demonstrates an implementation of io.Writer that counts bytes.
package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

//!+bytecounter

type ByteCounter int

type WordCounter int

type LineCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}
func (w *WordCounter) Write(p string) (int, error) {
	var isAlphaNum = regexp.MustCompile(`^[a-zA-Z0-9!,.?]+$`).MatchString

	for _, word := range strings.Split(p, " ") {
		if isAlphaNum(word) {
			*w += 1
		}
	}
	return int(*w), nil
}

//!-bytecounter

func main() {
	//!+main
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // "5", = len("hello")

	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, Dolly")

	// Exercise 7.1
	var w WordCounter
	w.Write("Hello gentlemen, how was your day?")
	fmt.Println(w)

	var fW io.Writer
	fW = os.Stdout
	//!-main
}
