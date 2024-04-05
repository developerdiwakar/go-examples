package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(t *testing.T) {
	a := 60
	b := 12
	addResult := doMath(a, b, add)
	// Basic test example
	if addResult == 72 {
		fmt.Println("Test Success.")
	} else {
		t.Fatalf("Test Failed. Got %v, expected %v", addResult, 10)
	}

	// Test Example using require package from testify
	require.GreaterOrEqual(t, a, b)
	// Test Example using assert package from testify
	assert.Equal(t, 72, doMath(a, b, add))
	assert.Equal(t, 48, doMath(a, b, sub))
	assert.Equal(t, 720, doMath(a, b, mul))
	assert.Equal(t, 5, doMath(a, b, div))

}
