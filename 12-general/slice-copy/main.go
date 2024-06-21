// Go program to illustrate how
// to copy an Slice by value

package main

import "fmt"

func main() {

	// Creating and initializing an slice
	// Using shorthand declaration
	my_slice1 := []string{"Scala", "Perl", "Java", "Python", "Go"}

	// Create a new slice with the same length as the original
	my_slice2 := make([]string, len(my_slice1))
	// Copy elements from the original slice to the new slice
	copy(my_slice2, my_slice1)

	fmt.Println("Slice_1: ", my_slice1)
	fmt.Println("Slice_2:", my_slice2)
	// made changes in the original slice
	my_slice1[0] = "C++"

	fmt.Println("\nSlice_1: ", my_slice1)
	fmt.Println("Slice_2: ", my_slice2)
}
