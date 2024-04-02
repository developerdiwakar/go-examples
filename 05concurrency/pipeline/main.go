package main

import "fmt"

func sliceToChannel(nums []int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(input <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range input {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	// input
	nums := []int{2, 4, 5, 3, 6}
	// stage 1
	sliceOfChannels := sliceToChannel(nums)
	// stage 2
	channelOfSquares := sq(sliceOfChannels)
	// stage 3
	for s := range channelOfSquares {
		fmt.Println(s)
	}
}
