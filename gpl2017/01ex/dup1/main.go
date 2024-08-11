package main

// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	// fmt.Println("input", input)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	// fmt.Println("counts", counts)
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}

// Note: To see the output first run the code, enter some input as a string,
// then press "Ctl+d"
