package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID int
}

func worker(i int, tasks <-chan Task, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // Signal that this worker is done when it finishes
	for task := range tasks {
		time.Sleep(time.Millisecond * 100)
		fmt.Printf("Worker %d processed Taks %d\n", i, task.ID)
		results <- task.ID // Send the result back to the results channel
	}
}

func main() {
	var wg sync.WaitGroup
	numWorkers := 3
	numTasks := 10

	var tasks = make(chan Task, numTasks)
	var results = make(chan int, numTasks)

	// Start Workers
	wg.Add(numWorkers)
	for i := 1; i <= numWorkers; i++ {
		go worker(i, tasks, results, &wg)
	}

	// Enqueue Tasks
	for i := 1; i <= numTasks; i++ {
		tasks <- Task{ID: i}
	}
	// Close the tasks channel after enqueuing all tasks
	close(tasks)

	// Wait for all workers to finish
	wg.Wait()
	// Close the results channel after all workers are done
	close(results)

	// Collect results
	for res := range results {
		// for i := 0; i < numTasks; i++ {
		// res := <-results
		fmt.Printf("Task %d completed.\n", res)
	}

}
