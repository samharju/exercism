// Package diffsquares solves sum of squares and square of sum difference.
package diffsquares

// SquareOfSum returns square of sum of first n natural numbers.
func SquareOfSum(n int) int {
	out := 0
	for i := 1; i <= n; i++ {
		out += i
	}
	return out * out
}

// SumOfSquares returns sum of squares of first n natural numbers.
func SumOfSquares(n int) int {
	out := 0
	for i := 1; i <= n; i++ {
		out += i * i
	}
	return out
}

// Difference returns the difference of square of sum vs sum of squares
// for first n natural numbers.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
