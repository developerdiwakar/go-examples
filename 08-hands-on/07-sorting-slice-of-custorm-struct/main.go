package main

import (
	"log"
	"sort"
)

// We can sort slices of custom structs by using sort.Sort and sort.Stable functions.
// These methods sort any collection that implements sort.
// Interface interface that has Len(), Less() and Swap() methods as shown in the code below:

type Interface interface {
	// find number of elements in collection
	Len() int
	// Less method is used for identifying which elements among index i and j are lesser and is used for sorting
	Less(i, j int) bool
	// Swap method is used for swapping elements with indexes i and j
	Swap(i, j int)
}

type Human struct {
	name string
	age  int
}

// AgeFactor implements sort.Interface that sorts the slice based on age field.
type AgeFactor []Human

func (a AgeFactor) Len() int {
	return len(a)
}

func (a AgeFactor) Less(i, j int) bool {
	return a[i].age < a[j].age
}

func (a AgeFactor) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	audience := []Human{
		{"Alice", 35},
		{"Bob", 32},
		{"James", 34},
	}

	sort.Sort(AgeFactor(audience))
	log.Println(audience)
}
