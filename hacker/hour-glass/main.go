package main

import "fmt"

func hourGlassSumMax(arr [][]int32) int32 {
	// Write your code here
	var sumOfHourGlass []int32
	var r int32
	var c int32
	var i int32
	var j int32
	var num int32 = 6
	var gRow int32 = num / 2
	var glass int32 = 0

	// Loop For NumberOfHourGlasses For Row
	for i = 0; i <= gRow; i++ {
		// Loop for NumberOfHourGlasses For Column
		for j = 0; j <= gRow; j++ {
			glass = glass + 1
			var glassSum int32 = 0
			// fmt.Println("glass: ", glass)
			// Loof for HourGlass Row
			iteration := 0
			for r = 0 + i; r < gRow+i; r++ {
				// Loof for HourGlass Column
				for c = 0 + j; c < gRow+j; c++ {
					iteration = iteration + 1
					value := arr[r][c]
					if iteration == 4 || iteration == 6 {
						continue
					}
					// fmt.Printf("(%d,%d) ", r, c)
					// fmt.Printf("%d ", value)
					glassSum += value
				}
				// fmt.Printf("%d ", glassSum)
				// fmt.Println("")
			}
			// Storing Sum of Glass of Hour
			sumOfHourGlass = append(sumOfHourGlass, glassSum)
		}
	}

	fmt.Println("Sum of Hour Glasses", sumOfHourGlass)
	// Find the Max sum
	maxSum := sumOfHourGlass[0]
	for i := 0; i < len(sumOfHourGlass)-1; i++ {
		if sumOfHourGlass[i+1] > maxSum {
			maxSum = sumOfHourGlass[i+1]
		}
	}

	return maxSum
}

func main() {
	// Change the Matrix according to the need
	var arr = [][]int32{
		{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 0, 2, 4, 4, 0},
		{0, 0, 0, 2, 0, 0},
		{0, 0, 1, 2, 4, 0},
	}
	fmt.Println("Matrix of 6X6: ", arr)
	maxSum := hourGlassSumMax(arr)
	fmt.Println("Answer is: ", maxSum)
}
