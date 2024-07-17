package main

import (
	"fmt"
	"time"
)

// GO Select: Example: 2
func main() {
	ch := make(chan string)

	go func() {
		ch <- "Hello from Example 2"
	}()

	timeout := time.After(1 * time.Second)

	select {
	case msg := <-ch:
		fmt.Println("Received from ch:", msg)
	case <-timeout:
		fmt.Println("Timed out waiting for message")
	}
}
