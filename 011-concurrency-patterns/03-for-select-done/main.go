package main

// For-Select-Done
// Combines a for loop with a select statement, often with a done channel to signal termination.

// Summary:

// Iterates with a for loop.
// Selects among multiple channels.
// Terminates on a signal.

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	done := make(chan bool)

	go func() {
		for {
			select {
			case msg := <-ch:
				fmt.Println(msg)
			case <-done:
				return
			}
		}
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch <- "hello"
		time.Sleep(1 * time.Second)
		done <- true
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("Bye-bye!")
}
