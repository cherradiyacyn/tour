package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	var formerVal float64
	// The delta means the difference between two numbers. It's mathematical stuff...
	for delta := 1.; delta > 1e-15; delta = math.Abs(z - formerVal) {
		formerVal = z
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Printf("math.Sqrt's result : %g\n", math.Sqrt(2))
	fmt.Printf("my result : %g\n", Sqrt(2))
}
