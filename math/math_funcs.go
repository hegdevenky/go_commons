package math

import (
	"golang.org/x/exp/constraints"
)

// Abs a generic function that returns the absolute value of a given signed number.
// Works on both signed & unsigned integers and floating point numbers.
func Abs[T constraints.Integer | constraints.Float](val T) T {
	if val < 0 {
		return -val
	}
	return val
}

// PercentageOf a generic function to calculate the percentage of a given number
// Returned result will be of type float64.
// Usage examples:
//   - to calculate 3% of 91 => PercentageOf(3, 91) or  PercentageOf[int](3, 91)
//   - to calculate 2.5% of 68.65 => PercentageOf(2.5, 68.65) or PercentageOf[float32](2.5, 68.65)
func PercentageOf[T constraints.Integer | constraints.Float](p, x T) float64 {
	return (float64(p) / 100) * float64(x)
}
