package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
	"time"
)

func allocateMemory() {
	_ = make([]byte, 1024*1024)
}

func main() {
	// Start the http server fro profiling
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	// profile memory usage
	go func() {
		for {
			// start memory profiling
			memFile, err := os.Create("memory_profile.prof")
			if err != nil {
				log.Fatalf("error creating memory profile file %v\n", err)
			}
			pprof.WriteHeapProfile(memFile)

			// memFile.Close()

			allocateMemory()

			time.Sleep(time.Second * 5)
		}

	}()

	// your application code here
	select {}
}
