package main

import "fmt"

func powerSeries(a int) (int, int) {
	square := a * a
	cube := a * a * a
	// return as a tuples
	return square, cube
}

func main() {
	square, cube := powerSeries(3)

	fmt.Printf("Square: %d, Cube: %d\n", square, cube)
}
