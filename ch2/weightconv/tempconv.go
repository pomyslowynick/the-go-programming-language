package weightconv

import "fmt"

// Exercise 2.1
type Kilo float64
type Pound float64

func (p Pound) String() string { return fmt.Sprintf("%g lbs", p) }
func (k Kilo) String() string  { return fmt.Sprintf("%g kilo", k) }
