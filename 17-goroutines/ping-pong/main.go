package main

import (
	"fmt"
	"sync"
)

var N int

func init() {
	N = 50
}

func ping(wg *sync.WaitGroup, pingSignChan <-chan struct{}, pongSignChan chan<- struct{}) {
	defer wg.Done()
	for i := 1; i <= N; i++ {
		<-pingSignChan
		fmt.Println("Ping", i)
		pongSignChan <- struct{}{}
	}
	close(pongSignChan)
}

func pong(wg *sync.WaitGroup, pongSignChan <-chan struct{}, pingSignChan chan<- struct{}) {
	defer wg.Done()
	for i := 1; i <= N; i++ {
		<-pongSignChan
		fmt.Println("Ping", i)
		pingSignChan <- struct{}{}
	}
	close(pingSignChan)
}

func main() {
	var pingSignChan = make(chan struct{}, 1)
	var pongSignChan = make(chan struct{}, 1)

	var wg sync.WaitGroup

	wg.Add(2)
	go ping(&wg, pingSignChan, pongSignChan)

	go pong(&wg, pongSignChan, pingSignChan)

	pingSignChan <- struct{}{}

	wg.Wait()

	fmt.Println("Done!")

}
