package main

import (
	"io"
)

// Exercise 7.2
// https://github.com/ray-g/gopl/blob/master/ch07/ex7.02/countingwriter.go
// Couldn't figure it out, took solution from the above website.

// The struct keeps the wrapper writer and the counter int
type CounterWriter struct {
	writer io.Writer
	count  int64
}

// Method implements the io.Writer interface, it just wraps the original method
// then takes the amount of bytes written and adds them to the count.
func (w *CounterWriter) Write(p []byte) (n int, err error) {
	n, err = w.writer.Write(p)
	w.count += int64(n)
	return
}

// The function creates a new "instance?" of CounterWriter struct with a type implementing the io.Writer
// and starts the count at 0. Then it returns the address of the new struct and the address of it's count property
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := &CounterWriter{
		writer: w,
		count:  0,
	}
	return cw, &cw.count
}

// I didn't think about the way in which I could just pass the variables within one method to another one
// I was trying to work out some solution in which I implement my own Write method.
// It turns out I had to do that but this method would be only passing arguments to the originally provided one
