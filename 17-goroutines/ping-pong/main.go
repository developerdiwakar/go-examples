package main

import (
	"fmt"
	"sync"
	"time"
)

var N int

func init() {
	N = 50_000
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
		fmt.Println("Pong", i)
		pingSignChan <- struct{}{}
	}
	close(pingSignChan)
}

func main() {
	start := time.Now()
	var pingSignChan = make(chan struct{}, 1)
	var pongSignChan = make(chan struct{}, 1)

	var wg sync.WaitGroup

	wg.Add(2)
	go ping(&wg, pingSignChan, pongSignChan)

	go pong(&wg, pongSignChan, pingSignChan)

	pingSignChan <- struct{}{}

	wg.Wait()

	fmt.Println("Done!")
	fmt.Println("Time elapsed:", time.Since(start))

}
