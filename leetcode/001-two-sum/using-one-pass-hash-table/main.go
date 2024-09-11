package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var m = make(map[int]int)
	for i, v := range nums {
		m[v] = i
		complement := target - v
		if j, ok := m[complement]; ok && i != j {
			return []int{j, i}
		}
	}

	return []int{}
}

func main() {
	target := 6
	nums := []int{3, 2, 4}
	result := twoSum(nums, target)

	fmt.Println("Result", result)
}
