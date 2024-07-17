package main

import (
	"fmt"
	"time"
)

func speak1(msg string, ch chan<- string) {
	time.Sleep(time.Second * 1)
	ch <- msg
}

func speak2(msg string, ch chan<- string) {
	time.Sleep(time.Second * 1)
	ch <- msg
}

func main() {
	var chan1 = make(chan string)
	var chan2 = make(chan string)
	var count = 10
	for i := 0; i < count; i++ {

		go speak1("Hello from Chan one", chan1)
		go speak2("Hello from chan two", chan2)
	}

	for i := 0; i < count; i++ {
		select {
		case result := <-chan1:
			fmt.Println("Received:", result)
		case result := <-chan2:
			fmt.Println("Received:", result)
		default:
			fmt.Println("Message not received yet")
			time.Sleep(200 * time.Millisecond)
		}
	}

	close(chan1)
	close(chan2)

}
