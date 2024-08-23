package main

import "fmt"

// [] = 1,2,3,5,6
// insert value 4 at index position 3
func insert(slice []int, index int, value int) []int {
	if index > len(slice) {
		return slice
	}
	slice = append(slice[:index], append([]int{value}, slice[index:]...)...)
	return slice
}

func main() {
	numbers := []int{1, 2, 3, 5, 6}
	numbers = insert(numbers, 3, 4)
	fmt.Println(numbers) // Output: [1 2 3 4 5 6]
}
