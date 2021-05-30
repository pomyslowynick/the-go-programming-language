package weightconv

// Exercise 2.2
func KToP(k Kilo) Pound { return Pound(k * 2.205) }
func PToK(p Pound) Kilo { return Kilo(p * 0.453592) }
