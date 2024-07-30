package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	m     sync.RWMutex
	value int
}

func (c *Counter) updateValue(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	c.m.Lock()
	fmt.Printf("Adding %d to %d\n", n, c.value)
	c.value += n
	c.m.Unlock()
}

func (c *Counter) readValue(wg *sync.WaitGroup) {
	defer wg.Done()
	c.m.RLock()
	fmt.Println("Get Value:", c.value)
	time.Sleep(400 * time.Millisecond)
	c.m.RUnlock()

}

func main() {
	var wg sync.WaitGroup

	c := Counter{}
	wg.Add(13)

	go c.updateValue(12, &wg)
	go c.readValue(&wg)
	go c.readValue(&wg)
	go c.readValue(&wg)
	go c.readValue(&wg)

	go c.updateValue(10, &wg)
	go c.readValue(&wg)
	go c.readValue(&wg)
	go c.readValue(&wg)

	wg.Wait()
}

// Note: sync.RWMutex
// Read and Write Locks: An RWMutex allows multiple readers or a single writer. It has two types of locks:
// RLock(): Acquires a read lock. Multiple goroutines can hold read locks simultaneously.
// Lock(): Acquires a write lock. Only one goroutine can hold a write lock, and it excludes all read locks.
// Use Case: Use RWMutex when reads are more frequent than writes. It allows concurrent reads, which can improve performance.
// Performance: Better performance for read-heavy workloads because multiple reads can proceed concurrently.
