package main

import (
	"fmt"
	"time"
)

func main() {
	// #### GO Select: Example: 1 ####
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Hello"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "World"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println("Received from ch1:", msg)
		case msg := <-ch2: // This case might not be executed if ch1 is ready first
			fmt.Println("Received from ch2:", msg)
			// default:
			// 	fmt.Println("No messages ready yet")
		}
	}

	// #### GO Select: Example: 2 ####
	// ch := make(chan string)

	// go func() {
	// 	ch <- "Hello"
	// }()

	// timeout := time.After(1 * time.Second)

	// select {
	// case msg := <-ch:
	// 	fmt.Println("Received from ch:", msg)
	// case <-timeout:
	// 	fmt.Println("Timed out waiting for message")
	// }
}
