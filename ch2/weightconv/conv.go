package weightconv

// Exercise 2.2
func KToP(k Kilo) Kilo   { return Pound(p * 2.205) }
func PToK(p Pound) Pound { return Kilo(p * 0.453592) }
