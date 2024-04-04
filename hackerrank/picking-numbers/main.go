package main

import (
	"fmt"
)

/*
 * Complete the 'pickingNumbers' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY a as parameter.
 */

func pickingNumbers(a []int32) int32 {
	// Write your code here
	var max int = 0
	var d int32
	var temp []int32
	var subA [][]int32
	for i := range a {
		temp = nil
		for j := range a {
			d = a[i] - a[j]
			if d <= 1 && d >= 0 {
				temp = append(temp, a[j])
			}
		}
		// fmt.Println(temp)
		if len(temp) >= 2 {
			subA = append(subA, temp)
		}
	}
	// result := removeDuplicates(subA)
	// fmt.Println(subA)
	if len(subA) > 0 {
		max = len(subA[0])
		for _, s := range subA {
			if len(s) > max {
				max = len(s)
			}
		}
		fmt.Println(max)
	}

	return int32(max)
}

// func removeDuplicates(slice [][]int) [][]int {
// 	encountered := map[string]bool{}
// 	result := [][]int{}

// 	for _, v := range slice {
// 		key := fmt.Sprintf("%v", v)
// 		if !encountered[key] {
// 			encountered[key] = true
// 			result = append(result, v)
// 		}
// 	}

// 	return result
// }

func main() {
	arr := []int32{1, 1, 2, 2, 4, 4, 5, 5, 5} // Output:5
	// arr := []int32{4, 6, 5, 3, 3, 1} // Output: 3
	// arr := []int32{4, 2, 3, 4, 4, 9, 98, 98, 3, 3, 3, 4, 2, 98, 1, 98, 98, 1, 1, 4, 98, 2, 98, 3, 9, 9, 3, 1, 4, 1, 98, 9, 9, 2, 9, 4, 2, 2, 9, 98, 4, 98, 1, 3, 4, 9, 1, 98, 98, 4, 2, 3, 98, 98, 1, 99, 9, 98, 98, 3, 98, 98, 4, 98, 2, 98, 4, 2, 1, 1, 9, 2, 4} // Output: 22

	pickingNumbers(arr)
}

// Write your code here// Write your code here// Write your code here
