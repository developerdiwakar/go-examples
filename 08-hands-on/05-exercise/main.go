package main

// Show the comma ok idiom starting with this code.
import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()
	// Received from Channel before closing channel
	v, ok := <-c
	fmt.Println(v, ok)

	close(c)
	// Received from Channel after closing channel
	v, ok = <-c
	fmt.Println(v, ok)
}
