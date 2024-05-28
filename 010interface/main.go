package main

import "log"

type Shape interface {
	area() float32
}

type Square struct {
	side float32
}

type Triangle struct {
	base   float32
	height float32
}

type Rectangle struct {
	length float32
	width  float32
}

// To implement an interface in Go, we just need to implement all the methods in the interface.
//
//	Here we implement shape interface method(area) on square.
func (s Square) area() float32 {
	return s.side * s.side
}

// Here we implement shape interface method(area) on triangle.
func (t Triangle) area() float32 {
	return t.base * t.height / float32(2)
}

// Here we implement shape interface method(area) on rectangle.
func (r Rectangle) area() float32 {
	return r.length * r.width
}

func calculateArea(s Shape) float32 {
	return s.area()
}

func main() {
	log.Println("Interface Example")
	r := Rectangle{7.4, 3.2}
	t := Triangle{3.1, 4.3}
	s := Square{3.5}

	log.Printf("Area of Rectangle: %f", calculateArea(r))
	log.Printf("Area of Triangle: %f", calculateArea(t))
	log.Printf("Area of Square: %f", calculateArea(s))
}
