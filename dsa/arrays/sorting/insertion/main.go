package main

import (
	"fmt"
	"time"
)

// Insertion sort is a simple sorting algorithm that works by iteratively inserting each element of an unsorted list into its correct position in a sorted portion of the list.
// It is a stable sorting algorithm, meaning that elements with equal values maintain their relative order in the sorted output.
// Insertion sort is like sorting playing cards in your hands. You split the cards into two groups: the sorted cards and the unsorted cards. Then, you pick a card from the unsorted group and put it in the right place in the sorted group.
// Insertion Sort Algorithm:
// Insertion sort is a simple sorting algorithm that works by building a sorted array one element at a time. It is considered an ” in-place ” sorting algorithm, meaning it doesn’t require any additional memory space beyond the original array.
// Algorithm:
// To achieve insertion sort, follow these steps:
// We have to start with second element of the array as first element in the array is assumed to be sorted.
// Compare second element with the first element and check if the second element is smaller then swap them.
// Move to the third element and compare it with the second element, then the first element and swap as necessary to put it in the correct position among the first three elements.
// Continue this process, comparing each element with the ones before it and swapping as needed to place it in the correct position among the sorted elements.
// Repeat until the entire array is sorted.

// InsertionSort method applies the insertion sort algorithm to sort an array
func InsertionSort(b []int) {
	fmt.Println(b)
	count := len(b)
	var temp int
	for j := 1; j < count; j++ {

		for i := j; i > 0 && b[i] < b[i-1]; i-- {
			temp = b[i]
			b[i] = b[i-1]
			b[i-1] = temp
		}
	}
}

func main() {
	var arr = []int{71, 41, 59, 26, 41, 58, 2, 21, 33, 4, 45, 55}

	start := time.Now()
	InsertionSort(arr)
	fmt.Println("InsertionSort takes:", time.Since(start).Seconds(), "Seconds")
	fmt.Println(arr)
	fmt.Println("Finish!")
}

// Note: Slice always passed as a reference(as a &arg) while passing in function arguments by default.
