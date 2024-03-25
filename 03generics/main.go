package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type MyList[T any] struct {
	data []T
}

// Constructor function to create MyList instances with type inference
func NewMyList[T any](values ...T) MyList[T] {
	return MyList[T]{data: values}
}

// *** Type Parameters ***
func maximum[T constraints.Ordered](x T, y T) T {
	if x > y {
		return x
	}
	return y
}

// *** Type Sets ***
func swap[T comparable](x *T, y *T) {
	temp := *x
	*x = *y
	*y = temp
}

func main() {

	// intMax := maximum[int](5, 10)
	intMax := maximum(5, 10) // T is inferred as int
	stringMax := maximum[string]("hello", "world")
	// stringMax := maximum("hello", "world") // T is inferred as string

	println("intMax: ", intMax)
	println("stringMax: ", stringMax)

	// Allowed: int, float(32|64), string (implement comparable)
	// var a int = 5
	// var b int = 8
	// var a float32 = 2.43
	// var b float32 = 3.14
	var a string = "hello"
	var b string = "world"
	swap(&a, &b)
	// Not allowed: MyCustomType (not comparable)
	fmt.Println("a =", a, "b =", b)

	// Type inference deduces the type argument (int) from the provided values
	myList := NewMyList(1, 2, 3)

	// Explicit type arguments can still be used for clarity
	anotherList := NewMyList[string]("apple", "banana")

	fmt.Println("myList:", myList)
	fmt.Println("anotherList:", anotherList)
}
