package proverb

// Proverb should have a comment documenting it.
func Proverb(input []string) []string {
	// Create an output array
	output := []string{}
	// Get the length of the input array
	inputLen := len(input)

	// Iterate the input array
	for i := 0; i < inputLen; i++ {
		// We need the actual and the next index
		actualIndex := i
		nextIndex := i + 1

		proverb := ""
		// Validate if the next index is valid compared with the input length
		if nextIndex > inputLen-1 {
			// If the next index is invalid, insert the short version of the proverb
			proverb = "And all for the want of a " + input[0] + "."
		} else {
			// If the next index is valid, insert the long version of the proverb
			proverb = "For want of a " + input[actualIndex] + " the " + input[nextIndex] + " was lost."
		}
		// Push the proverb into the output array
		output = append(output, proverb)
	}

	return output
}
