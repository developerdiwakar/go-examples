package main

import "fmt"

func main() {
	// Start Incrementor function
	nextNumber := startIncrementor()

	fmt.Println(nextNumber()) // out: 1
	fmt.Println(nextNumber()) // out: 2
	fmt.Println(nextNumber()) // out: 3
	fmt.Println(nextNumber()) // out: 4
	fmt.Println(nextNumber()) // out: 5
	fmt.Println(nextNumber()) // out: 6
	fmt.Println(nextNumber()) // out: 7
	fmt.Println(nextNumber()) // out: 8

}

func startIncrementor() func() int {
	var n int
	// Closure function
	return func() int {
		n++
		return n
	}
}
