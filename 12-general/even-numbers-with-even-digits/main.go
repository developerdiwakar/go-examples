package main

import "fmt"

// 10 and 1000 - find all evens that contains even digits

func main() {
	count := 1000
	var evens []int

	for i := 1; i < count; i++ {
		if i%2 == 0 {
			evens = append(evens, i)
		}
	}

	var final_evens []int
OUTER_LOOP:
	for _, v := range evens {
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

	fmt.Println(final_evens)
}
