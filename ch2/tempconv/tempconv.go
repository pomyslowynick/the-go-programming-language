package tempconv

import "fmt"

// Exercise 2.1
type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsouluteZeroC = -273.10
	FreezingC      = 0
	BoilingC       = 100
	AbsouluteZeroK = 0
	FreezingK      = 273.15
	BoilingK       = 373.15
)

func (c Celsius) String() string    { return fmt.Sprintf("%g C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g K", k) }
