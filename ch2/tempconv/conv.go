package tempconv

// Exercise 2.1
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*5/9 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 9 / 5) }
func KToC(c Celsius) Kelvin     { return Kelvin(c + 273.15) }
