package hamming

import (
	"errors"
)

// Distance function returns the Hamming distance is only defined for sequences of equal length, so an attempt to calculate it between sequences of different lengths should not work
func Distance(a, b string) (int, error) {
	// Validate if lenght of both entry are equal
	if len(a) != len(b) {
		// If length are different, we return an error
		err := "Both of the sequences must have the same lenght"
		return -1, errors.New(err)
	}

	// If they are equal, we do not need to iterate them and return this scenario
	if a == b {
		return 0, nil
	}

	// diff will be the output of how many chars are different
	diff := 0
	// Iterate one string and validate char by char if there are equal.
	for index := 0; index < len(a); index++ {
		// For each char, we compare both strins using runes
		if a[index] != b[index] {
			// If they are different, we will increment the return value by 1
			diff = diff + 1
		}
	}
	// Return the diff result
	return diff, nil
}
