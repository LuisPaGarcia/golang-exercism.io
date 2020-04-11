package strand

import "strings"

// ToRNA is a function
func ToRNA(input string) string {
	return strings.NewReplacer(
		"G", "C",
		"C", "G",
		"T", "A",
		"A", "U",
	).Replace(input)
}
