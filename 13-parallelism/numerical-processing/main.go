package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func sum(start, end int) int {
	sum := 0
	for i := start; i <= end; i++ {
		sum += i
	}
	return sum
}

func main() {
	var wg sync.WaitGroup
	var mu sync.RWMutex // TO handle data race condition for shared variable `total`
	total := 0
	numCores := 4       // Adjust based on your system
	dataSize := 1000000 // Divide the work into chunks for parallel processing
	chunkSize := dataSize / numCores
	start := time.Now()
	for i := 0; i < numCores; i++ {
		start := i * chunkSize
		end := (i+1)*chunkSize - 1
		if i == numCores-1 {
			end = dataSize // Adjust for the last chunk
		}

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			mu.Lock()
			log.Println("Total:", total)
			total += sum(start, end)
			mu.Unlock()
		}(start, end)
	}
	wg.Wait()
	fmt.Println("Total:", total)
	log.Println("time elapsed: ", time.Since(start))
}
