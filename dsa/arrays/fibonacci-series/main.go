package main

import "fmt"

// The Fibonacci sequence is defined as follows:
// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89,...N

func main() {
	a := 0
	b := 1
	N := 100
	fib := []int{}
	for a <= N {
		fib = append(fib, a)
		temp := a + b
		a = b
		b = temp
	}

	fmt.Println(fib)
}
