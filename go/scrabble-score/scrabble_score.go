// Package scrabble provides score calculator for Scrabble.
package scrabble

import "strings"

type point struct {
	letters string
	value   int
}

var points = []point{
	{"aeioulnrst", 1},
	{"dg", 2},
	{"bcmp", 3},
	{"fhvwy", 4},
	{"k", 5},
	{"jx", 8},
	{"qz", 10},
}

// Returns total score for formed word.
func Score(s string) int {
	score := 0
	for _, c := range strings.ToLower(s) {
		for _, point := range points {
			count := strings.Count(point.letters, string(c))
			if count != 0 {
				score += count * point.value
				break
			}
		}
	}
	return score
}
