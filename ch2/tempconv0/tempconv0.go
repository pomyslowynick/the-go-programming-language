package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsouluteZeroC Celsius = -273.15
	FreezingC      Celsius = 0
	Boiling        Celsius = 100
)

func main() {
	fmt.Println(CToF(AbsouluteZeroC))
	x := CToF(Boiling)
	fmt.Println(FToC(x))
	c := FToC(212.0)
	// fmt.Println(c.String())
	fmt.Printf("%v\n", c)
	fmt.Printf("%s\n", c)
	fmt.Println(c)
	fmt.Printf("%g\n", c)
	fmt.Println(float64(c))

}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// func (c Celsius) String() string {
// 	return fmt.Sprintf("%gC", c)
// }
