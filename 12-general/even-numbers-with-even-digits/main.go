package main

import "fmt"

// 10 and 1000 - find all evens that contains even digits

func main() {
	count := 1000
	var evens = make(chan int, count)
	var done = make(chan struct{})

	go func() {
		for i := 1; i < count; i++ {
			if i%2 == 0 {
				evens <- i
			}
		}
		close(evens)
	}()

	var final_evens []int
	go func() {
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
			final_evens = append(final_evens, v)
		}
		done <- struct{}{}
	}()

	<-done
	fmt.Println(final_evens)
}
