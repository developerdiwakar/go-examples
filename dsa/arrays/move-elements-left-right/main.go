package main

import "fmt"

// Function to move all 0's to the end while maintaining order of other elements
func moveZeroes(arr []int) {
	// Pointer for the position of non-zero elements
	nonZeroPos := 0

	// Iterate through the array
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			// Swap non-zero elements to the front
			arr[nonZeroPos], arr[i] = arr[i], arr[nonZeroPos]
			nonZeroPos++
		}
	}
}

// Function to move all negatives's to the end while maintaining order of other elements
func moveNegatives(arr []int) {
	// Pointer for the position of negative elements
	negPos := 0

	// Iterate through the array
	for i := 0; i < len(arr); i++ {
		if arr[i] >= 0 {
			// Swap negative elements to the front
			arr[negPos], arr[i] = arr[i], arr[negPos]
			negPos++
		}
	}
}

func main() {
	// Example array with zeros
	arr := []int{0, 1, 0, 3, 12}
	fmt.Println("Original array:", arr)
	moveZeroes(arr)
	fmt.Println("Array after moving 0's to the end:\n", arr)

	// Example array with negatives
	arr1 := []int{12, -4, -5, -11, 7, 10, -9, 13, -14}
	fmt.Println("Original array:", arr1)
	moveNegatives(arr1)
	fmt.Println("Array after moving negative's to the end:\n", arr1)

}
