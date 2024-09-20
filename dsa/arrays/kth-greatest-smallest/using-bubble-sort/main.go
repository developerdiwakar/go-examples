package main

import "fmt"

// IN: 4, 3, 6, 4, 1
// Find Kth largest element

// findKthElement return kth greatest number using bubble sort
// Note this technique also works for kth smallest number with little change
// Instead of returning arr[count-1-kth], arr[kth-1] need to be return
func findKthElement(arr []int, kth int) int {
	count := len(arr)
	swap := false
	for i := 0; i < count; i++ {
		for j := 0; j < count-i-1; j++ {
			swap = true
			if arr[j] > arr[j+1] {
				temp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp
			}
		}
		if !swap {
			break
		}

	}

	fmt.Println("arr=", arr)
	return arr[count-1-kth]
}

func main() {
	var arr = []int{4, 3, 6, 4, 1}
	kth := 3
	out := findKthElement(arr, kth)

	fmt.Printf("kth(%d) greatest element: %d\n", kth, out)
}
