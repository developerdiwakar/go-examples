package main

import "fmt"

// Input: 12,-4,-5,-11,7,10,-9,13,-14
// Output: -4,-5,-9,-11,-14,7,10,12,13
// Negative should be first and positives later. Ordering doesn't matter.
// Don't use extra variables.

func partition(arr []int) {
	left, right := 0, len(arr)-1

	for left <= right {
		if arr[left] < 0 {
			left++
		} else if arr[right] >= 0 {
			right--
		} else {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}
}

func main() {
	arr := []int{12, -4, -5, -11, 7, 10, -9, 13, -14}
	partition(arr)

	fmt.Println(arr)
}
