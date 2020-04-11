// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"regexp"
	"unicode"
)

var emptyAnswer string = "Fine. Be that way!"
var questionAnswer string = "Sure."
var nonSenseAnswer string = "Whatever."
var screamAnswer string = "Whoa, chill out!"
var screamQuestionAnswer string = "Calm down, I know what I'm doing!"

// Hey is a function that decides what message nees to be returned from Bob
func Hey(message string) string {
	// Clean the input for escape chars
	input := cleanString(message)

	// If the cleaned input is empty, return the correct answer
	if len(input) == 0 {
		return emptyAnswer
	}
	// Validate if the input string contains letters and all of them are uppercase
	isUpper := isUpperCase(input)
	// Validate if the string is a question
	isQuestion := isQuestion(input)

	// Just a validation if the statement is uppercase and is a question
	isUpperQuestion := isUpper && isQuestion

	// Return the message if is uppercase and a question
	if isUpperQuestion {
		return screamQuestionAnswer
	}

	// Return the message if is uppercase
	if isUpper {
		return screamAnswer
	}

	// Return the message if is a question
	if isQuestion {
		return questionAnswer
	}

	// Return the default response
	return nonSenseAnswer
}

func cleanString(input string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9?!,:)]+")
	if err != nil {
		panic(err)
	}
	return reg.ReplaceAllString(input, "")
}

func isUpperCase(input string) bool {
	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		panic(err)
	}

	inputReplaced := reg.ReplaceAllString(input, "")

	if len(inputReplaced) == 0 {
		return false
	}

	isUpper := true

	for _, r := range inputReplaced {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			isUpper = false
			break
		}
	}

	return isUpper
}

func isQuestion(input string) bool {
	return input[len(input)-1:] == "?"
}
