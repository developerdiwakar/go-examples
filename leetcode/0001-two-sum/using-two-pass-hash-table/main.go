package main

import "fmt"

var twoSum = func(nums []int, target int) []int {

	if len(nums) == 2 {
		return []int{0, 1}
	}
	var m = make(map[int]int)
	for i, v := range nums {
		m[v] = i
	}

	for i, v := range nums {
		complement := target - v
		if j, ok := m[complement]; ok && i != j {
			return []int{i, j}
		}
	}

	return nil

}

func main() {
	target := 9
	nums := []int{2, 7, 11, 15}
	result := twoSum(nums, target)

	fmt.Println("Result", result)
}
