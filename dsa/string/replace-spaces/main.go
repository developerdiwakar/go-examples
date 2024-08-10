package main

import "fmt"

// Given a string, write a function to replace all spaces in the string with '%20'.
// Assume that the string has enough space to hold the additional characters.

// Ex: hello, how are you? should become hello,%20how%20are%20you?

// replaceSpaces: 1st approach
func replaceSpaces(str string) string {
	// Calculate the number of spaces
	spaceCount := 0
	for _, char := range str {
		if char == ' ' {
			spaceCount++
		}
	}

	// Calculate the new lenght of the string
	newLength := len(str) + spaceCount*2

	// Create a new array byte to hold the result
	result := make([]byte, newLength)
	index := newLength - 1

	// Iterate over the original string in reverse
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == ' ' {
			result[index] = '0'
			result[index-1] = '2'
			result[index-2] = '%'

			index -= 3
		} else {
			result[index] = str[i]
			index--
		}
	}

	return string(result)
}

func main() {
	hello := "hello, how are you?"

	result := replaceSpaces(hello)
	fmt.Println(result)
}
