package main

import "fmt"

// write a program that
// ○ launches 10 goroutines
// 	■ each goroutine adds 10 numbers to a channel
// ○ pull the numbers off the channel and print them

func main() {
	nog := 10
	c := make(chan int)

	for i := 0; i < nog; i++ {
		go func() {
			for j := 20; j <= 30; j++ {
				c <- j
			}
			// close(c)
		}()
	}

	defer close(c) // (optional: closing the channel)
	for k := 0; k < 100; k++ {
		fmt.Printf("<-c = %d (%d) \t", <-c, k)
	}
	fmt.Println()
}
