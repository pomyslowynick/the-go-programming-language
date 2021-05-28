package main

import (
	"fmt"
	"os"
	"strconv"

	"example.com/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		temperatureC, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		temperatureF := tempconv.CToF(tempconv.Celsius(temperatureC))
		fmt.Println(temperatureF)
	}
}
