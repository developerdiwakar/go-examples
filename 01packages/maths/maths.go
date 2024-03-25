package maths

func Substract(a, b int) int {
	return a - b
}

func Add(a, b int) int {
	return a + b
}

// Exportable function (Public Modifier)
func Mul(a, b int) int {
	return multipy(a, b)
}

// Non-exportable function (Private Modifier)
func multipy(x, y int) int {
	return x * y
}
