// Package scrabble provides score calculator for Scrabble.
package scrabble

import "strings"

// Returns total score for formed word.
func Score(s string) int {
	score := 0
	for _, c := range strings.ToUpper(s) {
		switch c {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			score++

		case 'D', 'G':
			score += 2

		case 'B', 'C', 'M', 'P':
			score += 3

		case 'F', 'H', 'V', 'W', 'Y':
			score += 4

		case 'K':
			score += 5

		case 'J', 'X':
			score += 8

		case 'Q', 'Z':
			score += 10
		}
	}
	return score
}
