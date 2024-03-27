package main

import (
	"fmt"
	"time"
)

// Define the task struct, which represents the work to be done.
type Task struct {
	ID int
}

// Define the worker function, which processes tasks from a channel.
func worker(id int, tasks <-chan Task, results chan<- int) {
	for task := range tasks {
		// Simulate processing time for the task.
		time.Sleep(time.Second)
		fmt.Printf("Worker %d processed task %d\n", id, task.ID)
		results <- task.ID // Send the task ID to the results channel.
	}
}

func main() {
	numTasks := 10
	numWorkers := 3

	// Create channels for tasks and results.
	tasks := make(chan Task, numTasks)
	results := make(chan int, numTasks)

	// Create worker goroutines.
	for i := 1; i <= numWorkers; i++ {
		go worker(i, tasks, results)
	}

	// Enqueue tasks.
	for i := 1; i <= numTasks; i++ {
		tasks <- Task{ID: i}
	}
	close(tasks) // Close the tasks channel to signal that all tasks have been enqueued.

	// Collect results.
	for i := 0; i < numTasks; i++ {
		result := <-results
		fmt.Println("Task completed:", result)
	}
}
