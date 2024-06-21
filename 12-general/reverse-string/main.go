package main

import (
	"fmt"
	"log"
)

func reverseString(str string) string {
	count := len(str)
	runes := []rune(str)
	for i, j := 0, count-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)

}

func main() {
	var str string
	fmt.Println("Please enter any string to reverse(without space):")
	_, err := fmt.Scanln(&str)

	if err != nil {
		log.Fatalln("Error while taking input")
	}

	out := reverseString(str)

	fmt.Println("Revers of", str, "is", out)
}
