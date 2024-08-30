package main

import "fmt"

func removeDuplicateAndSort(s []int) []int {
	arrMap := make(map[int]struct{})
	var result []int

	for _, e := range s {
		if _, exists := arrMap[e]; !exists {
			arrMap[e] = struct{}{} // remove the element from a map
			result = append(result, e)
		}
	}
	// using bubble sort
	count := len(result)
	var swapped bool
	for i := 0; i < count; i++ {
		swapped = false
		for j := 0; j < count-i-1; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}

	return result
}

func main() {
	arr := []int{3, 4, 1, 2, 4, 5, 6, 7, 1, 1}
	fmt.Println("Before Removal: ", arr)
	result := removeDuplicateAndSort(arr)

	fmt.Println("After Removal: ", result)
}
