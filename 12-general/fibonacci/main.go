package main

import (
	"fmt"
	"log"
)

// Fibonacci series is the one in which you will get your next term by adding previous two numbers.
// For example ,
// Here, 0 + 1 = 1
//
//	1 + 1 = 2
//	3 + 2 = 5
//
// and so on.

// fibonacciSeriesUptoNth method will print the fibonacci series upto nth number.
// for example, [0 1 1 2 3 5 8 13 21 34 nth]
func fibonacciSeriesUptoNth(nth uint) []uint {
	var first uint = 0
	var second uint = 1
	var fibonacci []uint
	var temp uint
	for second <= nth {
		temp = first + second
		first = second
		fibonacci = append(fibonacci, first)
		second = temp
	}

	return fibonacci
}

// fibonacciSeriesForNth method will print the fibonacci series for nth number.
// example to print the first 12 numbers of a Fibonacci series.
// [0 1 1 2 3 5 8 13 21 34 55 89]
// func fibonacciSeriesForNth(nth uint) []uint {

// }

func main() {
	var nth uint

	fmt.Println("Please enter the nth number of Fibonacci series:")
	if _, err := fmt.Scanln(&nth); err != nil {
		log.Fatalf("Error while taking input: %v\n", err)
	}
	series := fibonacciSeriesUptoNth(nth)

	fmt.Println(series)
}
