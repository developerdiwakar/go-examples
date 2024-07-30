package main

// Prob: Find all even numbers between start and end that contain only even digits

import (
	"fmt"
	"sync"
)

func isEvenDigits(n int) bool {
	for n > 0 {
		if n&1 == 1 {
			return false
		}
		n /= 10
	}
	return true
}

func findEvenDigitEvens(start, end int, evens chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := start; i <= end; i++ {
		if isEvenDigits(i) {
			evens <- i
		}
	}
}

func main() {
	var count int
	fmt.Println("Enter the end number that you want even number with even digits:")
	fmt.Scan(&count)

	var wg sync.WaitGroup
	var numOfWorkers = 4
	var rangeSize = count / numOfWorkers
	var evens = make(chan int)

	// Started workers
	for i := 0; i < numOfWorkers; i++ {
		start := i*rangeSize + 1
		end := (i + 1) * rangeSize
		if i == numOfWorkers-1 {
			end = count
		}
		wg.Add(1)
		go findEvenDigitEvens(start, end, evens, &wg)
	}

	go func() {
		wg.Wait()
		close(evens)
	}()

	for even := range evens {
		fmt.Println(even)
	}
}
