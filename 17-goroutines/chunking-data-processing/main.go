package main

import (
	"fmt"
	"sync"
	"time"
)

// You are given a list of integers.
// Write a Go program that spawns multiple goroutines to process chunks of the list concurrently.
// Each goroutine should compute the sum of its chunk of integers and send the result back via a channel.
// The main goroutine should collect these results and compute the total sum of all integers.
// Ensure proper synchronization and avoid race conditions.

// input - nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func sum(start, end int, nums []int) int {
	sum := 0
	for start < end {
		sum += nums[start]
		start++
	}
	return sum
}

func main() {
	start := time.Now()
	var nums []int
	dataSize := 100000000
	for i := 1; i <= dataSize; i++ {
		nums = append(nums, i)
	}
	numCores := 100
	chunkSize := dataSize / numCores
	total := 0

	var wg sync.WaitGroup
	var mu sync.RWMutex

	wg.Add(numCores)
	for i := 0; i < numCores; i++ {
		start := i * chunkSize
		end := (i + 1) * chunkSize

		if i == numCores-1 {
			end = dataSize
		}
		go func(start, end int, nums []int) {
			defer wg.Done()
			mu.Lock()
			total += sum(start, end, nums)
			mu.Unlock()
		}(start, end, nums)

	}

	wg.Wait()

	fmt.Println("Total:", total)
	fmt.Println("Time elapsed:", time.Since(start))
}
