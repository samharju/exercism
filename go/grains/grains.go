// Package grains helps filthy peasants to fool their king.
package grains

import (
	"errors"
)

// Returns number of grains put on the given square.
// Input value must be in the range from 1 to 64.
func Square(s int) (uint64, error) {
	if s < 1 || s > 64 {
		return 0, errors.New("Invalid index")
	}
	return 1 << (s - 1), nil
}

// Returns the total number of grains if each square holds
// double the amount that was put in the previous square.
func Total() uint64 {
	return 1<<64 - 1
}
