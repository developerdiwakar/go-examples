package main

import (
	"fmt"
	"time"
)

//Bubble Sort is the simplest sorting algorithm that works by repeatedly swapping the adjacent elements if they are in the wrong order.
// This algorithm is not suitable for large data sets as its average and worst-case time complexity is quite high.

// Bubble Sort Algorithm
// In Bubble Sort algorithm,

// traverse from left and compare adjacent elements and the higher one is placed at right side.
// In this way, the largest element is moved to the rightmost end at first.
// This process is then continued to find the second largest and place it and so on until the data is sorted.

// BubbleSort function applies the bubble sort algorithm on the given array
func BubbleSort(a []int) {
	fmt.Println(a)
	count := len(a)
	var max = a[0]
	var swapped bool
	for i := 0; i < count-1; i++ {
		swapped = false
		for j := 0; j < count-i-1; j++ {
			if a[j] > a[j+1] {
				max = a[j]
				a[j] = a[j+1]
				a[j+1] = max

				swapped = true
			}

			if !swapped {
				break
			}
		}

	}

}

func main() {
	// var arr = []int{71, 41, 59, 26, 41, 58, 2, 21, 33, 4, 45, 55}
	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 8}

	start := time.Now()
	BubbleSort(arr)
	fmt.Println("BubbleSort takes:", time.Since(start).Seconds(), "Seconds")
	fmt.Println(arr)
	fmt.Println("Finish!")
}
