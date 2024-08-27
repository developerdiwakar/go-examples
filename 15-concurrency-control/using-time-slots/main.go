package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	// Simulate some work
	fmt.Printf("Goroutine %d is running...\n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d is done.\n", id)
}

func main() {
	// Configuration for the number of goroutines and time slot
	maxGoroutines := 5
	minGoroutines := 2
	tickerInterval := 10 * time.Second

	sem := make(chan struct{}, maxGoroutines) // Semaphore to limit goroutines
	ticker := time.NewTicker(tickerInterval)
	// defer ticker.Stop()

	var wg sync.WaitGroup

	for range ticker.C {
		fmt.Printf("Starting new time slot at: %v\n\n", time.Now())

		// Launch a minimum number of goroutines
		for i := 0; i < minGoroutines; i++ {
			sem <- struct{}{} // Occupy a slot
			wg.Add(1)
			go func(id int) {
				defer func() { <-sem }()
				defer wg.Done()
				worker(id)
			}(i)
		}

		// Launch additional goroutines up to the max limit
		for i := minGoroutines; i < maxGoroutines; i++ {
			sem <- struct{}{} // Occupy a slot
			wg.Add(1)
			go func(id int) {
				defer func() { <-sem }()
				defer wg.Done()
				worker(id)
			}(i)
		}

		// Wait for all goroutines in this time slot to finish
		wg.Wait()
		fmt.Printf("\nAll goroutines for this time slot are done at: %v\n\n", time.Now())
	}
}
