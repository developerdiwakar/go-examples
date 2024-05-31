package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	srv := &http.Server{Addr: ":8080"}
	// Create a channel to listen for shutdown signals
	done := make(chan bool)
	quit := make(chan os.Signal, 1)

	// Listen for SIGINT (Ctrl + C) and SIGKILL
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-quit // Wait for a shutdown signal
		log.Println("Shutdown signal received")
		// Set time for graceful shutdown
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("Shutdown error: %v\n", err)
		}

		close(done)
	}()

	log.Println("Server started on port 8080")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, there!")
	})

	// Run the server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalln(err)
		}
	}()
	// Block until shutdown is complete
	<-done
	log.Println("Server Stopped")
}
