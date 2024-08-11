package main

import (
	"fmt"
	"os"
	"strings"
)

func efficientJoin(args []string) string {
	return strings.Join(args, " ")
}

func unEfficientJoin(args []string) string {
	var s, sep string
	for _, v := range args[1:] {
		s += sep + v
		sep = " "
	}
	return s
}

func main() {
	args := os.Args
	// fmt.Println("1.1\n")
	// fmt.Println(strings.Join(args[1:], " "))

	// fmt.Println("1.2")
	// for i, v := range args[1:] {
	// 	fmt.Println(i, v)
	// }

	fmt.Println(efficientJoin(args))
	fmt.Println(unEfficientJoin(args))

}
