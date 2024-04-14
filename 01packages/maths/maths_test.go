package maths

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	want := 4
	if res := Add(2, 2); res != want {
		t.Errorf("Got: %v, Want: %v", res, want)
	}
}

func ExampleAdd() {
	fmt.Println(Add(2, 4))
	//Output:
	// 6
}

// BenchmarkAdd test the Add function upto b.N time
// To run this test use go test -bench BenchmarkAdd or go test -bench .
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(2, 3)
	}
}
