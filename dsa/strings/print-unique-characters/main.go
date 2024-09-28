package main

import (
	"fmt"
	"strings"
)

// I am an Indian
// Output : I am n d

func main() {
	s := "I am an Indian"
	// Converted to lower
	s = strings.ToLower(s)
	// Defined a map to count to characters occurence
	var smap = make(map[byte]int, len(s))

	for i := 0; i < len(s); i++ {
		if smap[s[i]] < 1 {
			fmt.Printf("%s", string(s[i]))
		}
		if s[i] != ' ' {
			smap[s[i]]++
		}
	}
	fmt.Println()
}
