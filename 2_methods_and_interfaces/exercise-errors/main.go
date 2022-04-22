package main

import "math"

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return ""
}

func Sqrt(x float64) (float64, error) {
	z := float64(1)
	var formerVal float64
	// https://en.wikipedia.org/wiki/Absolute_difference
	for absDiff := 1.; absDiff > 1e-15; absDiff = math.Abs(z - formerVal) {
		formerVal = z
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main() {

}
