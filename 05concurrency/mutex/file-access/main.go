package main

import (
	"fmt"
	"os"
	"sync"
)

// Consider a scenario where multiple processes need to read or write to a file concurrently.
// Here's a basic example:

// FileHandler provides safe concurrent access to a file
type FileHandler struct {
	mu   sync.Mutex
	path string
}

// ReadFile reads the content of the file
func (fh *FileHandler) ReadFile() ([]byte, error) {
	fh.mu.Lock()
	defer fh.mu.Unlock()
	return os.ReadFile(fh.path)
}

// WriteFile writes data to the file
func (fh *FileHandler) WriteFile(data []byte) error {
	fh.mu.Lock()
	defer fh.mu.Unlock()
	return os.WriteFile(fh.path, data, 0644)
}

func main() {
	var wg sync.WaitGroup
	fileHandler := &FileHandler{path: "data.txt"}

	// Concurrent read/write operations from multiple processes
	wg.Add(1)
	go func() {
		defer wg.Done()
		data, err := fileHandler.ReadFile()
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		fmt.Println("Read data:", string(data))
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := fileHandler.WriteFile([]byte("New data"))
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
		fmt.Println("Data written successfully")
	}()

	// Wait for Goroutines to finish
	wg.Wait()

	fmt.Println("File operations complete")
}

// In this example, the FileHandler struct uses a mutex to control access to the file specified by the path field.
// The ReadFile and WriteFile methods acquire the lock before performing file operations, ensuring that only one process can read from or write to the file at a time.
// This prevents data corruption issues that might arise from concurrent access.

// These are just two illustrative examples.
// Mutexes are a versatile tool for synchronization in various scenarios where concurrent access to shared resources needs to be controlled in Go applications.
// Remember to use them strategically to address potential race conditions and maintain data integrity in your concurrent programs.
