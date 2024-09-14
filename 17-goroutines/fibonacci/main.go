package main

import (
	"fmt"
	"sync"
)

// The Fibonacci sequence is defined as follows:
// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89,...N

func fibonacci(N int, fibChan chan<- []int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(fibChan)
	fib := []int{}
	a := 0
	b := 1
	for a <= N {
		fib = append(fib, a)
		temp := a + b
		a = b
		b = temp
	}
	fibChan <- fib
}

func main() {
	var wg sync.WaitGroup
	var fibChan = make(chan []int)
	N := 100
	wg.Add(1)
	go fibonacci(N, fibChan, &wg)
	fmt.Println(<-fibChan)
	wg.Wait()
}
