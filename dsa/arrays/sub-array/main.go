package main

// Find the all possible sub-array whose length is 3
// and sum of the elements is 6 from the given array with positive numbers.
// For ex:
// Input: [0,1,2,3,4,5,6,7]
// Output: [ [0,1,5], [0,2,4], [1,2,3] ]

import "fmt"

func findSubarrays(nums []int) [][]int {
	result := [][]int{}
	n := len(nums)

	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if nums[i]+nums[j]+nums[k] == 6 {
					result = append(result, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}

	return result
}

func main() {
	nums := []int{0, 1, 2, 3, 4, 6, 5, 7, 8}
	result := findSubarrays(nums)
	fmt.Println(result)
}
