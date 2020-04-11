// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package leap should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package leap

// IsLeapYear should have a comment documenting it.
func IsLeapYear(year int) bool {
	// Validate if the year is exctly divisible by 400
	if year%400 == 0 {
		return true
	}

	// Validate if the year is divisible by 4 but not by 100
	if year%4 == 0 && year%100 != 0 {
		return true
	}

	// Any other case, will be false
	return false
}
