package main

import "fmt"

func main() {

	fmt.Println("a+b=", doMath(3, 4, add))
	fmt.Println("a-b=", doMath(3, 4, sub))
	fmt.Println("a*b=", doMath(3, 4, mul))
	fmt.Println("a/b=", doMath(3, 4, div))
}

func doMath(a, b int, f func(int, int) int) int {

	return f(a, b)
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func mul(x, y int) int {
	return x * y
}

func div(x, y int) int {
	return x / y
}
