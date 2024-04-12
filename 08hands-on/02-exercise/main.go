package main

// Got Error: invalid operation: cannot receive from send-only channel cs (variable of type chan<- int)
// Get this code running:
// ○ https://play.golang.org/p/oB-p3KMiH6
// ○ https://play.golang.org/p/_DBRueImEq

import (
	"fmt"
)

func main() {
	cs := make(<-chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
