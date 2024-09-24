package main

import "fmt"

// Given an array arr[] of size N-1 with integers in the range of [1, N],
// the task is to find the missing number from the first N integers.
// Note: There are no duplicates in the list.
// Input: arr[] = {1, 2, 4, 6, 3, 7, 8} , N = 8
// Output: 5

func missingNumber(arr []int, N int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	// The sum of the first n natural numbers: N * (N + 1) / 2
	return N*(N+1)/2 - sum
}

func main() {
	N := 8
	arr := []int{1, 2, 3, 4, 6, 7, 8}
	num := missingNumber(arr, N)

	fmt.Println("Missing number:", num)
}
