package main

import (
	"fmt"
	"strings"
)

func ExtractUsername(email string) string {
	at := strings.Index(email, "@")
	return email[:at]
}

func main() {
	fmt.Println(ExtractUsername("diwakar@emal.com"))
	fmt.Println(ExtractUsername("diwakar.kumar@emal.com"))
}
