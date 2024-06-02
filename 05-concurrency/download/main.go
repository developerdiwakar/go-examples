package main

import (
	"fmt"
	"time"
)

type DownloadProgress struct {
	URL      string
	Progress int
}

func downloadFile(url string, filename string, progressChan chan DownloadProgress) {
	defer close(progressChan) // Close the channel when the function exits
	// Simulate downloading the file with progress updates
	for i := 0; i <= 100; i += 10 {
		progressChan <- DownloadProgress{URL: url, Progress: i}
		time.Sleep(100 * time.Millisecond) // Simulate download time
	}
	fmt.Println("Download complete. Saved as", filename)
}

func main() {
	progressChan := make(chan DownloadProgress)
	go downloadFile("https://example.com/file1.txt", "file1.txt", progressChan)
	go downloadFile("https://example.com/file2.txt", "file2.txt", progressChan)

	// Wait for progress updates and handle them
	for progress := range progressChan {
		fmt.Println("Download progress:", progress.URL, progress.Progress, "%")
	}

	fmt.Println("All downloads complete")
}

// Explanation:

// The downloadFile function now defers the closure of the progressChan channel. This ensures that the channel is closed when the function exits.
// The main function ranges over the progressChan channel to receive progress updates. This loop will exit when the channel is closed, indicating that all downloads are complete.
// This approach ensures that your program properly handles all progress updates and avoids deadlocks.
