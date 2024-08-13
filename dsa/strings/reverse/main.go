package main

import (
	"fmt"
)

func reverseString(s string) string {
	// Convert the string to a slice of runes
	runes := []rune(s)
	// Iterate over the first half of the runes slice
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		// Swap the runes
		runes[i], runes[j] = runes[j], runes[i]
	}
	// Convert the slice of runes back to a string
	return string(runes)
}

func main() {
	str := "Hello, 世界!"
	fmt.Println("Original string:", str)

	reversed := reverseString(str)
	fmt.Println("Reversed string:", reversed)
}
