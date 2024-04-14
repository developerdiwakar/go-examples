// maths package includes some basic mathematical operations
package maths

// Subtract function subtracts any two numbers of type int
func Subtract(a, b int) int {
	return a - b
}

// Add function adds any two numbers of type int
func Add(a, b int) int {
	return a + b
}

// Exportable function (Public Modifier)
// Mul function Call multiply function internally
func Mul(a, b int) int {
	return multipy(a, b)
}

// Non-exportable function (Private Modifier)
// multiply function multiplied two given numbers of type int
func multipy(x, y int) int {
	return x * y
}
