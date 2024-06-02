package main

// Semaphore
// Controls access to a common resource with a counting semaphore.

// Summary:
// Controls resource access.
// Uses a counting semaphore.

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, sem chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	sem <- struct{}{}
	fmt.Printf("Worker %d started\n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Worker %d done\n", id)
	// func() { <-sem }()
	<-sem
}

func main() {
	const maxConcurrent = 2
	sem := make(chan struct{}, maxConcurrent)
	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go worker(i, sem, &wg)
	}

	wg.Wait()
}
