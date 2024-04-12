package main

import "fmt"

// get this code working:
// ○ using func literal, aka, anonymous self-executing func
// ○ a buffered channel
func main() {
	c := make(chan int)
	c <- 42

	fmt.Println(c)
}
