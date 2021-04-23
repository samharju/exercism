// Package raindrops provides functions for making noise with numbers.
package raindrops

import "strconv"

var conversions = []struct {
	value  int
	output string
}{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}

// Convert returns the noise that given number of raindrops create.
func Convert(i int) string {
	out := ""
	for _, conv := range conversions {
		if i%conv.value == 0 {
			out += conv.output
		}
	}
	if out == "" {
		return strconv.Itoa(i)
	}
	return out
}
