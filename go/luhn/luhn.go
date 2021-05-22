// Package luhn implements char sequence validation with luhn algorithm.
package luhn

import (
	"strings"
)

// Valid checks string validity against luhn checksum formula.
func Valid(in string) bool {
	sum := 0
	cc := strings.Replace(in, " ", "", -1)
	if !(len(cc) > 1) {
		return false
	}
	for i, r := range cc {
		val := int(r - '0')
		if val > 9 {
			return false
		}
		if (len(cc)-i)%2 == 0 {
			val = val * 2
		}
		if val > 9 {
			val -= 9
		}
		sum += val
	}
	return sum%10 == 0

}
