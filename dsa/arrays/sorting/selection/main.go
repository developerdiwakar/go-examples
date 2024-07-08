package main

import (
	"fmt"
	"time"
)

// Selection sort is a simple and efficient sorting algorithm that works by repeatedly selecting the smallest (or largest) element from the unsorted portion of the list and moving it to the sorted portion of the list.
// The algorithm repeatedly selects the smallest (or largest) element from the unsorted portion of the list and swaps it with the first element of the unsorted part.
//  This process is repeated for the remaining unsorted portion until the entire list is sorted.

// SelectionSort applies the selection sort algorithm to sort an array
func SelectionSort(a []int) {
	fmt.Println(a)
	var minIdx, swap int
	count := len(a)
	for i := 0; i < count-1; i++ {
		minIdx = i
		for j := i + 1; j < count; j++ {
			if a[j] < a[minIdx] {
				minIdx = j
			}
		}

		if minIdx != i {
			swap = a[i]
			a[i] = a[minIdx]
			a[minIdx] = swap

		}
	}
}

func main() {
	var arr = []int{71, 41, 59, 26, 41, 58, 2, 21, 33, 4, 45, 55}

	start := time.Now()
	SelectionSort(arr)
	fmt.Println("SelectionSort takes:", time.Since(start).Seconds(), "Seconds")
	fmt.Println(arr)
	fmt.Println("Finish!")
}
