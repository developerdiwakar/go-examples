package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

// Example 1: Golang Program code to Check whether a string is palindrome or not
func checkStringPalindrome(str string) bool {
	count := len(str)
	for i := 0; i < count/2; i++ {
		if str[i] != str[count-1-i] {
			return false
		}
	}

	return true
}

// Example 2: Golang Program Code to Check whether a Number is palindrome or not
func checkNumberPalindrome(number uint64) bool {
	num := number
	var reverseNum uint64 = 0
	var a uint64
	for {
		a = num % 10
		reverseNum = reverseNum*10 + a
		num /= 10

		if num == 0 {
			break
		}
	}

	if number == reverseNum {
		return true
	} else {
		return false
	}

}

// removeSpecialCharacters removes non-alphanumeric characters from a string
func removeSpecialCharacters(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r) {
			return r
		}
		return -1
	}, str)
}

func main() {
	// Taking string input from user for palindrome check
	var str string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a string:")

	// Read the input until the newline character
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Trim the newline character from the end
	str = strings.TrimSpace(input)
	// Replace spaces with empty string
	str = strings.ReplaceAll(str, " ", "")
	// Convert to lower case
	str = strings.ToLower(str)
	// Remove special characters
	str = removeSpecialCharacters(str)

	if checkStringPalindrome(str) {
		fmt.Println(str, "is Palindrome")
	} else {
		fmt.Println(str, "is not Palindrome")
	}

	var number uint64
	fmt.Println("Enter a number(without space): ")
	_, err = fmt.Scanln(&number)
	if err != nil {
		log.Fatalf("Error reading input: %v\n", err)
		return
	}

	if checkNumberPalindrome(number) {
		fmt.Println(number, "is Palindrome")
	} else {
		fmt.Println(number, "is not Palindrome")
	}

}
