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
// Note:
// tasks <-chan Task: A receive-only channel of type Task. This channel is used to receive tasks that need to be processed by the worker.
// results chan<- int: A send-only channel of type int. This channel is used to send the results of the processed tasks back to the caller.
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

// Explanation:

// Task Struct: We define a Task struct to represent the work that needs to be done. In this example, a task is simply identified by its ID.

// Worker Function: The worker function represents a worker goroutine. It receives tasks from a tasks channel, processes them, and sends the results (in this case, the task ID) to a results channel.

// Main Function:

// We specify the number of tasks (numTasks) and the number of worker goroutines (numWorkers).
// We create channels for communicating tasks and results (tasks and results channels).
// We spawn worker goroutines, each executing the worker function, passing them the tasks channel to receive tasks and the results channel to send results.
// We enqueue tasks by sending them to the tasks channel.
// We close the tasks channel to signal that all tasks have been enqueued.
// We collect results by receiving them from the results channel and printing them out.

// This worker pool example demonstrates how to limit the number of concurrent tasks being processed by controlling the number of worker goroutines. Each worker is responsible for processing a task and sending the result back.
