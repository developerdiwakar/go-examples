package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Println("Server listening on port: 8080")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Create context with Cancellation
	ctx, cancel := context.WithCancel(r.Context())
	// Ensure cancellation on function exit
	defer cancel()

	// launch long running task on separate go routine
	go longRunningTask(ctx)

	// provide a way for the user to cancel the task (e.g. a button)
	select {
	case <-ctx.Done():
		log.Println("Request cancel by the user!")
	case <-r.Context().Done(): // handle content cancellation from the Http request
		log.Println("Request timed out!")
	}
}

func longRunningTask(ctx context.Context) {
	// Simulate some long running task
	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done():
			log.Println("Long running task cancelled!")
			return
		case <-time.After(time.Second * 1):
			log.Printf("Task progress: %d", i+1)
		}
	}
	log.Println("Long running task completed!")
}
