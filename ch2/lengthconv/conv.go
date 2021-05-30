package lengthconv

// Exercise 2.2
func MToI(m Meter) Inch { return Inch(m / 0.0254) }
func IToM(i Inch) Meter { return Meter(i * 39.37007874) }
