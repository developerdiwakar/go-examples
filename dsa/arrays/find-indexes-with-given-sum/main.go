package main

import (
	"fmt"
)

// findIndexWithGivenSum finds the index where the sum of the elements up to that index equals the given sum
func findIndexWithGivenSum(arr []int, targetSum int) int {
	runningSum := 0

	for i, num := range arr {
		runningSum += num
		if runningSum == targetSum {
			return i
		}
	}

	return -1 // return -1 if no such index is found
}

func main() {
	arr := []int{1, 2, 3, -4, 4, 2, -2, 5, -5}
	targetSum := 6

	index := findIndexWithGivenSum(arr, targetSum)
	if index != -1 {
		fmt.Printf("The sum of elements up to index %d equals the target sum %d\n", index, targetSum)
	} else {
		fmt.Println("No index found where the sum of elements equals the target sum.")
	}
}
