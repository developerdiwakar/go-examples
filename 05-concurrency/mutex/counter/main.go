package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	m     sync.Mutex
	value int
}

func (c *Counter) Update(n int, wg *sync.WaitGroup) {
	c.m.Lock()
	defer wg.Done()
	fmt.Printf("Adding %d to %d\n", n, c.value)
	c.value += n
	c.m.Unlock()
}

func main() {
	var wg sync.WaitGroup

	c := Counter{}
	wg.Add(4)

	go c.Update(10, &wg)
	go c.Update(-5, &wg)
	go c.Update(25, &wg)
	go c.Update(19, &wg)

	wg.Wait()

	fmt.Printf("Result is %d\n", c.value)

}

// Note: sync.Mutex
// Exclusive Lock: A Mutex provides mutual exclusion, allowing only one goroutine to access the critical section at a time.
// Use Case: Use Mutex when you need exclusive access to a resource and when reads and writes are not significantly imbalanced.
// Performance: Since it only allows one goroutine to proceed, it can be a bottleneck if many goroutines need read access.
