package main

// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// A map is a reference to the dat a st ruc ture cre ate d by make.
	counts := make(map[string]int)
	// Args hold the command-line arguments, starting with the program name
	files := os.Args[1:]
	//Stdin, Stdout, and Stderr are open Files pointing to the standard input,
	// standard output, and standard error file descriptors.
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2.main: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}

	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}

// os.File represents an open file descriptor.
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
