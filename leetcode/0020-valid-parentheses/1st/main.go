package main

import (
	"fmt"
	"time"
)

func isValid(exp string) bool {
	var stack []rune
	bracs := map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	}

	for _, c := range exp {
		if c == '{' || c == '(' || c == '[' {
			stack = append(stack, c)
		} else if c == '}' || c == ')' || c == ']' {
			if len(stack) == 0 {
				return false
			} else if stack[len(stack)-1] == bracs[c] {
				// deleting last element
				stack = stack[:len(stack)-1]
				continue
			} else {
				return false
			}

		}
	}
	// fmt.Println(stack)
	return len(stack) == 0
}

func main() {
	start := time.Now()

	expr1 := "a + ({ b - c} + d ( e * e)) - [d*d]"

	expr2 := "a + (({ b - c} + d ( e * e)) - [d*d]"

	expr3 := "(a + )({ b - c} + d ( e * e)) - [d*d]"

	expr4 := "a + ({ b - c)} + d ( e * e) - [d*d]"

	list := []string{expr1, expr2, expr3, expr4}

	for _, l := range list {
		if isValid(l) {
			fmt.Println("Correct")
		} else {
			fmt.Println("Not Correct")
		}
	}

	fmt.Println("Time elapsed:", time.Since(start))
}
