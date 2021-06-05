// The sha256 command computes the SHA256 hash (an array) of a string.
package main

import (
	"crypto/sha256"
	"fmt"
)

//!+

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	for i := range c1 {
		fmt.Printf("XOR of c1 and c2: %b\n", c1[i]&c2[i])
		fmt.Printf("c1 %b\n", c1[i])
		fmt.Printf("c2 %b\n", c2[i])

	}
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8
}

//!-