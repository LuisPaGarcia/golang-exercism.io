// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import (
	"math"
)

// Kind KindFromSides() returns this type. Pick a suitable data type.
type Kind string

const (
	// NaT means not a triangle
	NaT = "NaT"
	// Equ means equilateral
	Equ = "Equ"
	// Iso means isosceles
	Iso = "Iso"
	// Sca means scalene
	Sca = "Sca"
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	// Validate if the entries are supperior to 0
	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}
	// Validate if any input is Infinite
	if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return NaT
	}
	// Validate if any input is Not A Number
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return NaT
	}

	// Validate the Triangle Inequality.
	if a+b < c || c+a < b || b+c < a {
		return NaT
	}
	// Validate if it's equilateral
	if a == b && b == c {
		return Equ
	}
	// Validate if it's isosceles
	if a == b || b == c || a == c {
		return Iso
	}
	// If any of the rules apply, it's a scalene
	return Sca
}
