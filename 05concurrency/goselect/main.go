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
		time.Sleep(1000 * time.Millisecond)
		ch1 <- "Hello"
	}()

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch2 <- "World"
	}()

	// time.Sleep(100 * time.Millisecond) // after 100 milli seconds only channel ch1 will be received
	time.Sleep(1000 * time.Millisecond) // after 1000 milli seconds both channels(ch1,ch2) will be received

	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println("Received from ch1:", msg)
		case msg := <-ch2: // This case might not be executed if ch1 is ready first
			fmt.Println("Received from ch2:", msg)
		default:
			fmt.Println("No messages ready yet")
		}
	}

	// #### GO Select: Example: 2 ####
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
