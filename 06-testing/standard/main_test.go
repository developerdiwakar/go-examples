package main

import (
	"fmt"
	"testing"
)

// The function name must be start with TestXxxx (Xxxx)
func TestExtractUsername(t *testing.T) {
	t.Run("withoutDot", func(t *testing.T) {
		email := "ddiwakar@email.com"
		username := ExtractUsername(email)
		if username != "diwakar" {
			t.Fatalf("Got: %v, Expected: %v\n", username, "ddiwakar")
		}
	})

	t.Run("withDot", func(t *testing.T) {
		email := "diwakar.kumar@email.com"
		username := ExtractUsername(email)
		if username != "diwakar.kumar" {
			t.Fatalf("Got: %v, Expected: %v\n", username, "diwakar.kumar")
		}
	})
}

// ExampleFuncName - This is another way to test the method
func ExampleExtractUsername() {
	email := "ddiwakar@email.com"
	fmt.Println(ExtractUsername(email))
	//Output:
	// diwakar
}
