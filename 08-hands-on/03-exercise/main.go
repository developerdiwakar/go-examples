package main

// Starting with this code, pull the values off the channel
// using a for range loop

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Goroutine: ", runtime.NumGoroutine())
	c := gen()
	fmt.Println("Goroutine: ", runtime.NumGoroutine())
	receive(c)
	fmt.Println("Goroutine: ", runtime.NumGoroutine())
	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
