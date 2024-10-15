package main

import (
	"fmt"
	"sync"
)

func performSum(i, a, b int, sum []int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	mu.Lock()
	sum[i] = a + b
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	numWorkers := 10
	var sum = make([]int, numWorkers)

	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go performSum(i, 2+i, 3+i, sum, &wg, &mu)
	}
	wg.Wait()

	for _, result := range sum {
		fmt.Println(result)
	}

}
