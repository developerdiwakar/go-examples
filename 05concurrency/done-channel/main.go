package main

import (
	"fmt"
	"time"
)

// This doWork function only accept the receive-only channel with type bool.
// It means the "done" channel is used to receive the data.
func doWork(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("Doing work")
		}
	}
}

func main() {
	done := make(chan bool)

	go doWork(done)

	time.Sleep(3 * time.Second)

	fmt.Println("I'm done")
}
