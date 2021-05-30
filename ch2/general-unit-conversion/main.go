package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"example.com/lengthconv"
	"example.com/tempconv"
	"example.com/weightconv"
)

// Exercise 2.2

var temperatureFlag = flag.Bool("t", false, "convert the provided value to temperature")
var weightFlag = flag.Bool("w", false, "convert the provided value to weight in pounds")
var lengthFlag = flag.Bool("l", false, "convert the provided value to temperature")

func main() {
	flag.Parse()
	if *temperatureFlag {
		argNumber, err := strconv.ParseFloat(flag.Arg(0), 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error happened when parsing the cli arg %v", err)
		}
		fmt.Println(tempconv.CToF(tempconv.Celsius(argNumber)))
	}

	if *weightFlag {
		argNumber, err := strconv.ParseFloat(flag.Arg(0), 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error happened when parsing the cli arg %v", err)
		}
		fmt.Println(weightconv.KToP(weightconv.Kilo(argNumber)))

	}
	if *lengthFlag {
		argNumber, err := strconv.ParseFloat(flag.Arg(0), 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error happened when parsing the cli arg %v", err)
		}
		fmt.Println(lengthconv.MToI(lengthconv.Meter(argNumber)))
	}

}
