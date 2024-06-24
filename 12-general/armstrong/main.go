package main

import (
	"fmt"
	"log"
)

// An Armstrong number is the one whose value is equal to the sum of the cubes of its digits.
// 0, 1, 153, 371, 407, 471, etc are Armstrong numbers.
// For example,
// 407 = (4*4*4) + (0*0*0) + (7*7*7)
//         = 64 + 0 + 343
// 407 = 407

func checkArmstrong(num uint64) bool {
	n := num
	var sum uint64 = 0
	for n > 0 {
		a := n % 10
		sum = sum + (a * a * a)
		n = n / 10
		fmt.Print(sum, " ")
	}

	return sum == num
}

func main() {
	var number uint64
	fmt.Println("Enter the number to check for Armstrong number:")
	if _, err := fmt.Scanln(&number); err != nil {
		log.Fatalf("error while taking input: %v\n", err)
	}

	if checkArmstrong(number) {
		fmt.Printf("\n%d is an armstrong number", number)
	} else {
		fmt.Printf("\n%d is not an armstrong number", number)

	}
}
