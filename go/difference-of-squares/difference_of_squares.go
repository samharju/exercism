// Package diffsquares solves sum of squares and square of sum difference.
package diffsquares

// SquareOfSum returns square of sum of first n natural numbers.
func SquareOfSum(n int) int {
	sum := n * (n + 1) / 2
	return sum * sum
}

// SumOfSquares returns sum of squares of first n natural numbers.
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference returns the difference of square of sum vs sum of squares
// for first n natural numbers.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
