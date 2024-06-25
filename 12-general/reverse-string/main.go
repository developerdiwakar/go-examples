package main

import (
	"fmt"
	"log"
)

func reverseString(str string) string {
	count := len(str)
	bytes := []byte(str)
	for i, j := 0, count-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}

	return string(bytes)

}

func main() {
	var str string
	fmt.Println("Please enter any string to reverse(without space):")
	_, err := fmt.Scanln(&str)

	if err != nil {
		log.Fatalf("Error while taking input: %v", err)
	}

	out := reverseString(str)

	fmt.Println("Revers of", str, "is", out)
}
