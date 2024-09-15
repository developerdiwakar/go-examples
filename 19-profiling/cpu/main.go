package main

import (
	"fmt"
	"log"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"

	"net/http"

	"time"
)

func expensiveOperation() {

	// Simulate a time-consuming operation

	for i := 0; i < 10_00_00_000; i++ {

		_ = i

	}

}

func main() {

	// Start the HTTP server for profiling

	go func() {

		http.ListenAndServe("localhost:6060", nil)

	}()

	// Profile the CPU usage

	go func() {

		for {

			startTime := time.Now()
			// Start CPU profiling
			f, err := os.Create("cpu_profile.prof")
			if err != nil {
				log.Fatalf("Error creating cpu profile file %v\n", err)
				return
			}
			pprof.StartCPUProfile(f)
			f.Close()
			// Run your application code

			expensiveOperation()

			// Stop CPU profiling

			pprof.StopCPUProfile()

			elapsedTime := time.Since(startTime)

			fmt.Printf("CPU profiling took %s\n", elapsedTime)

			time.Sleep(5 * time.Second) // Run profiling every 5 seconds (adjust as needed)

		}

	}()

	// Your application code here

	select {}

}
