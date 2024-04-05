package main

import (
	"fmt"
	"sync"
)

// In-Memory Cache with Mutex:

// Imagine a web application that caches frequently accessed data in memory to improve performance.
// Here's a simplified example:

// Cache represents a simple in-memory key-value store
type Cache struct {
	mu   sync.Mutex
	data map[string]string
}

// Get retrieves a value from the cache
func (c *Cache) Get(key string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.data[key]
	return val, ok
}

// Set stores a key-value pair in the cache
func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

func main() {
	var wg sync.WaitGroup
	cache := &Cache{data: make(map[string]string)}

	// Concurrent access from multiple Goroutines (e.g., web requests)
	wg.Add(1)
	go func() {
		wg.Done()
		cache.Set("key1", "value1")
	}()
	wg.Add(1)
	go func() {
		wg.Done()
		val, ok := cache.Get("key1")
		if ok {
			fmt.Println("Retrieved from cache:", val)
		} else {
			fmt.Println("Key not found in cache")
		}
	}()

	// Wait for Goroutines to finish (optional)
	wg.Wait()

	fmt.Println("Cache operations complete")
}

// In this example, the Cache struct uses a mutex (mu) to protect the underlying data map.
// The Get and Set methods acquire the lock before accessing or modifying the map, ensuring that only one Goroutine can modify the cache at a time.
// This prevents race conditions and data corruption when multiple requests access the cache concurrently.
