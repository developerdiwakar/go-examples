package main

import "fmt"

// Zero's should be left, Ordering doesn't matter.
// IN: {7, 10, 0, 4, 3, 0, 20, 15}
// OUT: {7, 10, 15, 4, 3, 20, 0, 0}

func twoPointer(arr []int) {
	left, right := 0, len(arr)-1

	for left < right {
		if arr[left] <= 0 {
			left++
		} else if arr[right] > 0 {
			right--
		} else {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

}

func main() {
	var arr = []int{7, 10, 0, 4, 3, 0, 20, 15}
	twoPointer(arr)
	fmt.Println(arr)

}
