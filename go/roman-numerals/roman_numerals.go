package romannumerals

import (
	"errors"
)

type conversion struct {
	integer int
	alpha   string
}

var numerals = []conversion{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ToRomanNumeral(i int) (out string, err error) {

	if i < 1 || i > 3000 {
		return "", errors.New("Please dont")
	}
	for _, conv := range numerals {
		for i >= conv.integer {
			out += conv.alpha
			i -= conv.integer
		}
	}

	return
}
