package main

import (
	"fmt"
	"runtime"
)

// write a program that
// ○ launches 10 goroutines
// 	■ each goroutine adds 10 numbers to a channel
// ○ pull the numbers off the channel and print them

func main() {
	nog := 10
	c := make(chan int)
	fmt.Println("Rotines: ", runtime.NumGoroutine())
	for i := 0; i < nog; i++ {
		go func() {
			for j := 20; j <= 30; j++ {
				c <- j
			}
			// close(c)
		}()
	}
	fmt.Println("Rotines: ", runtime.NumGoroutine())

	for k := 0; k < 100; k++ {
		fmt.Printf("<-c = %d (%d) \t", <-c, k)
	}
	// close(c) // (optional: closing the channel)
	fmt.Println()
	fmt.Println("Rotines: ", runtime.NumGoroutine())
}
