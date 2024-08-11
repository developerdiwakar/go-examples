package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)

	for _, filename := range os.Args[1:] {
		// ReadFile reads the named file and returns the contents.
		// A successful call returns err == nil, not err == EOF.
		// Because ReadFile reads the whole file,
		// it does not treat an EOF from Read as an error to be reported.
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v \n", err)
			continue
		}
		// Split slices s into all substrings separated by sep and returns
		// a slice of the substrings between those separators.
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}

	}

	for text, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, text)
		}
	}
}
