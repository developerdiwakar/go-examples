package main

import (
	"fmt"
	"sync"
	"time"
)

func distributor(N int, oddChan, evenChan chan<- int) {
	defer close(oddChan)
	defer close(evenChan)
	for i := 1; i <= N; i++ {
		// Bit manipulation(using bitwise AND (&) operator) to find odd and even
		if i&1 == 0 {
			evenChan <- i
		} else {
			oddChan <- i
		}
	}
}

func processOdds(odds []int, oddChan <-chan int, wg *sync.WaitGroup) []int {
	defer wg.Done()
	for num := range oddChan {
		odds = append(odds, num)
	}
	return odds
}

func processEvens(evens []int, evenChan <-chan int, wg *sync.WaitGroup) []int {
	defer wg.Done()
	for num := range evenChan {
		evens = append(evens, num)
	}
	return evens
}

func main() {
	start := time.Now()

	// Set N number
	N := 50_00

	// Create channels
	oddChan := make(chan int)
	evenChan := make(chan int)

	// WaitGroup to synchronize goroutines
	var wg sync.WaitGroup

	// Distributor goroutine to send numbers to oddChan and evenChan
	go distributor(N, oddChan, evenChan)

	// Slices to store results
	var odds []int
	var evens []int

	// Goroutine to process odd numbers
	wg.Add(1)
	go func() {
		odds = processOdds(odds, oddChan, &wg)
	}()

	// Goroutine to process even numbers
	wg.Add(1)
	go func() {
		evens = processEvens(evens, evenChan, &wg)
	}()

	// Wait for processing to complete
	wg.Wait()

	// Print results
	fmt.Printf("Odds: %d\n\n", odds)
	fmt.Printf("Evens: %d\n", evens)

	fmt.Printf("Time elapse: %v\n", time.Since(start))
}
