package main

import "fmt"

// Revers a string preserving the space positions
// in <- "i am Diwakar"
// out -> "r ak awiDmai"

func main() {
	str := "i am Diwakar"
	// out -> "r ak awiDmai"

	runes := []rune(str)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {

		if runes[i] == 32 {
			i++
		}
		if runes[j] == 32 {
			j--
		}

		runes[i], runes[j] = runes[j], runes[i]
	}

	fmt.Println(string(runes))

}
