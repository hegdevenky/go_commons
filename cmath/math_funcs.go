package cmath

import (
	"golang.org/x/exp/constraints"
	"iter"
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

// Ints function returns the sequence of integers from start to end (inclusive)
// similar to Java's IntStream.rangeClosed in concept.
//
// It produces a lazy, read-only, ordered sequence of integers that can be
// iterated using Go's native `range` loop syntax. The sequence is ascending
// if start <= end, or descending if start > end.
//
// This function is particularly useful for functional-style programming,
// pipelines, or concise iteration over a known range of integers.
//
// Example usage:
//
//	for i := range Ints(1, 5) {
//	    fmt.Print(i, " ") // Output: 1 2 3 4 5
//	}
//
//	for i := range Ints(5, 1) {
//	    fmt.Print(i, " ") // Output: 5 4 3 2 1
//	}
func Ints[T constraints.Integer](start, end T) iter.Seq[T] {
	if start <= end {
		return func(yield func(T) bool) {
			for ; start <= end; start++ {
				if !yield(start) {
					return
				}
			}
		}
	} else {
		return func(yield func(T) bool) {
			for ; start >= end; start-- {
				if !yield(start) {
					return
				}
			}
		}
	}
}
