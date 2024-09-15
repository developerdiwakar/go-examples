package main

import "fmt"

// Prime Numbers: 2,3,5,7,11,13,17,19,23,29,..N
func main() {
	N := 100
	var primes []int
	for i := 2; i < N; i++ {
		isPrime := true
		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, i)
		}
	}

	fmt.Println("Primes:", primes)
}
