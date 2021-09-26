package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	var formerVal float64
	// https://en.wikipedia.org/wiki/Absolute_difference
	for absDiff := 1.; absDiff > 1e-15; absDiff = math.Abs(z - formerVal) {
		formerVal = z
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	f := 32.
	fmt.Printf("math.Sqrt's result : %g\n", math.Sqrt(f))
	fmt.Printf("my result : %g\n", Sqrt(f))
}
