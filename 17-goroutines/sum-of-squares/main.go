package main

import (
	"fmt"
)

// /*Sum of all squares between 1 and d.

// Use go routine and channel to compute the square and sum.

// chan1: Use to read/write the squares. Not allowed to send the n or index (i) from main() to SquaresSum() through channel.
// quit: Use it to quit the computations.*/

func SquaresSum(ch1 chan<- int, quit <-chan struct{}) {
	go func() {
		i := 1
		for {
			select {
			case <-quit:
				close(ch1)
				return
			default:
				ch1 <- i * i
				i++
			}
		}
	}()

}

func main() {

	chan1 := make(chan int)
	quit := make(chan struct{}) // Done channel

	sum := 0
	n := 4

	go func() {
		for i := 0; i < n; i++ {
			sum += <-chan1
		}
		quit <- struct{}{}

	}()

	SquaresSum(chan1, quit)
	<-quit // Wait for the goroutine to signal completion
	fmt.Println("Sum", sum)
}

// N = 4
// INPUT = 1 + 4 + 9 + 16
// Output = 30
