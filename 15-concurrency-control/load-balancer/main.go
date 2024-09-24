package main

import (
	"fmt"
	"time"
)

func worker(i int, reqChan <-chan Request, done chan<- struct{}) {
	for req := range reqChan {
		fmt.Printf("Worker %d is handling Request %d\n", i+1, req.ID)
		time.Sleep(100 * time.Millisecond)
	}
	done <- struct{}{}
}

type Request struct {
	ID int
}

func main() {

	numWorkers := 3
	numRequests := 100
	var reqChan = make([]chan Request, numWorkers)
	var done = make(chan struct{})

	// Create channels for each worker
	for i := 0; i < numWorkers; i++ {
		reqChan[i] = make(chan Request)
		go worker(i, reqChan[i], done)
	}

	// Distribute requests to workers in round-robin fashion
	for i := 0; i < numRequests; i++ {
		workerId := i % numWorkers
		reqChan[workerId] <- Request{ID: i + 1}
	}

	// Close all worker channels once all requests are sent
	for i := 0; i < numWorkers; i++ {
		close(reqChan[i])
	}

	// Wait for all workers to finish
	for i := 0; i < numWorkers; i++ {
		<-done
	}

	fmt.Println("All requests handled.")
}
