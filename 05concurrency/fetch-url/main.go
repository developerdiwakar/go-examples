package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

// Task represents a URL to be fetched
type Task struct {
	URL string
}

// Fetcher fetches content from a URL
func Fetcher(tasks <-chan Task, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		client := http.Client{
			Timeout: 2 * time.Second, // Timeout for HTTP requests
		}
		resp, err := client.Get(task.URL)
		if err != nil {
			log.Println("Error fetching URL:", err)
			continue
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("%v", string(body))
		// Simulate processing the response
		fmt.Println("Fetched content from:", task.URL)
		results <- fmt.Sprintf("Content from %s", task.URL)
	}
}

func main() {
	// Channel for tasks (URLs)
	tasks := make(chan Task)

	// Channel for results
	results := make(chan string)

	// WaitGroup to track worker Goroutines
	var wg sync.WaitGroup

	urls := []string{
		"https://go.abc/play/",
		"https://url3.dev",
		"https://jsonplaceholder.typicode.com/posts/1",
		"https://jsonplaceholder.typicode.com/posts/2",
		"https://jsonplaceholder.typicode.com/posts/3",
		"https://jsonplaceholder.typicode.com/posts/4",
		"https://jsonplaceholder.typicode.com/posts/5",
		"https://jsonplaceholder.typicode.com/posts/6",
		"https://jsonplaceholder.typicode.com/posts/7",
		"https://jsonplaceholder.typicode.com/posts/8",
		"https://jsonplaceholder.typicode.com/posts/9",
		"https://jsonplaceholder.typicode.com/posts/10",
		"https://jsonplaceholder.typicode.com/posts/11",
		"https://jsonplaceholder.typicode.com/posts/12",
		"https://jsonplaceholder.typicode.com/posts/13",
		"https://jsonplaceholder.typicode.com/posts/14",
		"https://jsonplaceholder.typicode.com/posts/15",
		"https://jsonplaceholder.typicode.com/posts/16",
		"https://jsonplaceholder.typicode.com/posts/17",
		"https://jsonplaceholder.typicode.com/posts/18",
		"https://jsonplaceholder.typicode.com/posts/19",
		"https://jsonplaceholder.typicode.com/posts/20",
		"https://jsonplaceholder.typicode.com/posts/21",
		"https://jsonplaceholder.typicode.com/posts/22",
	}
	// Number of worker Goroutines
	numWorkers := len(urls)

	// Launch worker Goroutines
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		// wg.Add(1)
		go Fetcher(tasks, results, &wg)
	}

	// Define tasks (URLs) with a mutex to protect concurrent access (optional)
	tasksMutex := &sync.Mutex{}
	for _, url := range urls {
		tasksMutex.Lock()
		tasks <- Task{URL: url}
		tasksMutex.Unlock()
	}

	// Close the tasks channel after adding all tasks
	close(tasks)

	// Wait for all worker Goroutines to finish
	go func() {
		wg.Wait()
		close(results) // Close results channel after all tasks are processed
	}()

	// Print results received from worker Goroutines
	for result := range results {
		fmt.Println("Result:", result)
	}

	fmt.Println("All tasks completed!")
}
