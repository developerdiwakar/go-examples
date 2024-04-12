package main

// Show the comma ok idiom starting with this code.
import (
	"fmt"
)

func main() {
	c := make(chan int)

	v, ok :=
		fmt.Println(v, ok)

	close(c)

	v, ok =
		fmt.Println(v, ok)
}
