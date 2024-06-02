package main

import (
	"fmt"
	"packages/maths"
	advanced "packages/maths/advanced"
)

func init() {
	fmt.Println("Golang: Example of packages")
	fmt.Println("Defining and importing package")
}

func main() {
	var a int = 2
	var b int = 3
	fmt.Printf("%d+%d = %d\n", a, b, maths.Add(a, a))
	fmt.Printf("%d-%d = %d\n", a, b, maths.Subtract(a, b))
	fmt.Printf("%d*%d = %d\n", a, b, maths.Mul(a, a))
	fmt.Printf("Area of Circle: %d\n", maths.AreaOfTriangle(a, b))
	fmt.Printf("Square of %d: %d\n", a, advanced.Square(a))
}
