package main

import (
	"fmt"
	"log"
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
	var swapped bool
	for i := 0; i < count; i++ {
		swapped = false
		for j := 0; j < count-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}

	}

}

func main() {
	var arr = []int{71, 41, 59, 26, 41, 58, 2, 21, 33, 4, 45, 55}

	start := time.Now()
	BubbleSort(arr)
	log.Println("BubbleSort takes:", time.Since(start).Seconds(), "Seconds")
	fmt.Println(arr)
	fmt.Println("Finish!")
}
