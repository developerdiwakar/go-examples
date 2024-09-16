package main

import "fmt"

func romanToInt(s string) int {
	var roman = make(map[byte]int, 7)
	roman = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	sum := 0
	for i := 0; i < len(s); i++ {

		if i == len(s)-1 {
			sum += roman[s[i]]
			break
		}

		if roman[s[i]] < roman[s[i+1]] {
			sum -= roman[s[i]]
		} else {
			sum += roman[s[i]]
		}
	}

	return sum
}

func main() {
	fmt.Println(romanToInt("LVIII"))
	// Output: 58
	// Explanation: L = 50, V= 5, III = 3
	fmt.Println(romanToInt("MCMXCIV"))
	// Output: 1994
	// Explanation: M = 1000, CM = 900, XC = 90 and IV = 4
}
