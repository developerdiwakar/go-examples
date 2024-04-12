package main

// Starting with this code, pull the values off the channel
// using a select statement

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()

	return c
}

func receive(c1, q1 <-chan int) {
	for {
		select {
		case i := <-c1:
			fmt.Println(i)
		case <-q1:
			return
		}
	}
}
