// Package isogram has one function to validate isograms.
package isogram

import "strings"

// IsIsogram checks if given string is an isogram.
func IsIsogram(s string) bool {
	ss := strings.ToLower(s)
	for _, c := range ss {
		if c == '-' || c == ' ' {
			continue
		}
		if strings.Count(ss, string(c)) > 1 {
			return false
		}
	}
	return true
}
