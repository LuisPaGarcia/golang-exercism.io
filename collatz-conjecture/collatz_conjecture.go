package collatzconjecture

import "errors"

// CollatzConjecture returns the number of steps to reach 1 using Collatz Conjecture algorithm
func CollatzConjecture(number int) (int, error) {
	// Throw an error if is lower than 0
	if number < 1 {
		return 0, errors.New("The input must be an positive integer")
	}
	// Step counter setted as zero
	steps := 0
	// Init a while loop, until number to be zero
	for number != 1 {
		// Increment the steps if enters
		steps++
		// If is even
		if number%2 == 0 {
			number = number / 2
			// Continue ignores the next statements for this loop
			continue
		}
		// If is not even, is odd
		number = (number * 3) + 1
	}
	// Return the step count and no error
	return steps, nil
}
