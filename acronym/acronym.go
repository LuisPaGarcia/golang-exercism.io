// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate return the acronim of the phrase
func Abbreviate(s string) string {
	// Clean the input from speacial chars
	input := cleanInput(s)
	output := ""
	// Separate every word in an string array
	words := strings.Fields(input)
	// Iterate over the array
	for i := 0; i < len(words); i++ {
		// Append the first char of every word to output
		output = output + string(words[i][0])
	}
	// Return the uppercase version of the output
	return strings.ToUpper(output)
}

// cleanInput replaces the items not defined in the Regex by blank spaces
func cleanInput(s string) string {
	// Create a regex to search on this chars
	reg, err := regexp.Compile("[^a-zA-Z' ]+")
	if err != nil {
		panic(err)
	}
	// Replace all the values NON inclued in the regex and replace it by blank spaces
	return reg.ReplaceAllString(s, " ")

}
