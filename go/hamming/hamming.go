// Package hamming provides functions for analysing text differencies
package hamming

import "errors"

// Distance returns hamming distance of input strings.
// Input strings must be of same length.
func Distance(a, b string) (int, error) {
	runesA, runesB := []rune(a), []rune(b)
	if len(runesA) != len(runesB) {
		return 0, errors.New("input strings are not of equal length")
	}
	dist := 0
	for i := 0; i < len(runesA); i++ {
		if runesA[i] != runesB[i] {
			dist++
		}
	}
	return dist, nil

}
