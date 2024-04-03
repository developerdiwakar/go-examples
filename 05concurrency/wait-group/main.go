package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(n int) {
	fmt.Printf("Worker %d - starting\n", n)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d - Done\n", n)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // Increment waitGroup counter by 1 for each goroutine
		// Create a local copy of i (to remove warning: loop variable i captured by func literal)
		// i := i
		// Wrap the worker call in a closure that makes sure to tell the WaitGroup that this worker is done.
		// This way the worker itself does not have to be aware of the concurrency primitives involved in its execution.
		go func() { // started several goroutine worker
			defer wg.Done() //
			worker(i)
		}()
	}
	wg.Wait()
	fmt.Println("All Done")
}
