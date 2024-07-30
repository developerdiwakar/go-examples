package main

import (
	"fmt"
	"sync"
)

// Function to check if a number contains only even digits
func hasOnlyEvenDigits(n int) bool {
	for n > 0 {
		digit := n % 10
		if digit%2 != 0 {
			return false
		}
		n /= 10
	}
	return true
}

// Function to find all even numbers between start and end that contain only even digits
func findEvenDigitEvens(start, end int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := start; i <= end; i++ {
		if i%2 == 0 && hasOnlyEvenDigits(i) {
			results <- i
		}
	}
}

func main() {
	var N int
	fmt.Println("Enter the value of N:")
	fmt.Scan(&N)

	var wg sync.WaitGroup
	results := make(chan int)

	// Number of goroutines to use
	numGoroutines := 4
	rangeSize := N / numGoroutines

	// Start goroutines
	for i := 0; i < numGoroutines; i++ {
		start := i*rangeSize + 1
		end := (i + 1) * rangeSize
		if i == numGoroutines-1 {
			end = N
		}
		wg.Add(1)
		go findEvenDigitEvens(start, end, results, &wg)
	}

	// Collect results in a separate goroutine
	go func() {
		wg.Wait()
		close(results)
	}()

	// Print results
	fmt.Println("Even numbers with only even digits between 1 and", N, ":")
	for num := range results {
		fmt.Printf("%d\t", num)
	}
}
