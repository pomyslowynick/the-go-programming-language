package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Exercise 1.12
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
	value := r.URL.Query().Get("cycles")
	parse_value, err := strconv.ParseFloat(value, 64)
	if err != nil {
		fmt.Println("Error %v\n", err)
	}
	lissajous(w, parse_value)
}
