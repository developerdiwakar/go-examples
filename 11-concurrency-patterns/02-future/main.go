package main

// Future
// A future represents a value that will be available at some point. It’s used to handle results of asynchronous operations.

// Summary:

// Represents a value available in the future.
// Used for handling async results.
import (
	"fmt"
	"time"
)

func asyncFunction() chan int {
	future := make(chan int)
	go func() {
		time.Sleep(1 * time.Second)
		future <- 42
	}()
	return future
}

func main() {
	future := asyncFunction()
	fmt.Println("Waiting for result...")
	result := <-future
	fmt.Println("Result:", result)
}
