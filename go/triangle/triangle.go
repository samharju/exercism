// Package triangle for analyzing triangles.
package triangle

import "math"

type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides returns type of triangle with given measures.
func KindFromSides(a, b, c float64) Kind {

	for _, v := range []float64{a, b, c} {
		if !(v > 0) || math.IsInf(v, 0) {
			return NaT
		}
	}
	if a+b < c || a+c < b || c+b < a {
		return NaT
	}
	if a == b && b == c {
		return Equ
	}
	if a == b || a == c || b == c {
		return Iso
	}
	return Sca
}
