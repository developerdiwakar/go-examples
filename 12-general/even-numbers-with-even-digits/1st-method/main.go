package main

import "fmt"

// 10 and 1000 - find all evens that contains even digits
func getEvens(count int, evens chan<- int) {
	for i := 1; i < count; i++ {
		if i%2 == 0 {
			evens <- i
		}
	}
	close(evens)
}

func getFinalEvens(evens <-chan int, final_evens chan<- int, done chan<- struct{}) chan<- int {
	defer close(final_evens)
OUTER_LOOP:
	for v := range evens {
		if v >= 10 {
			n := v
			for n > 0 {
				d := n % 10
				n = n / 10
				if d%2 != 0 {
					continue OUTER_LOOP
				}
			}
		}
		final_evens <- v
	}
	done <- struct{}{}
	return final_evens
}

func main() {
	count := 1000
	var evens = make(chan int)
	var done = make(chan struct{})
	var final_evens = make(chan int)

	go getEvens(count, evens)

	go getFinalEvens(evens, final_evens, done)

	go func() {
		<-done
	}()
	for v := range final_evens {
		fmt.Printf("%d\t", v)
	}
	fmt.Println("\nFinished!")
}
