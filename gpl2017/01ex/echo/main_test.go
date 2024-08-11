package main

import (
	"fmt"
	"testing"
)

func TestEfficientJoinBasic(t *testing.T) {
	input := []string{"command", "arg1", "arg2", "arg3"}
	output := "command arg1 arg2 arg3"
	result := efficientJoin(input)
	if result != output {
		t.Errorf("Get: %s Want: %s", result, output)
	}
}

func TestEfficientJoinTableDriven(t *testing.T) {
	var tests = []struct {
		input []string
		want  string
	}{
		{[]string{"command", "arg1", "arg2", "arg3"}, "command arg1 arg2 arg3"},
		{[]string{"12", "arg1", "23", "89"}, "12 arg1 23 89"},
		{[]string{"12", "arg3", "23", "45"}, "12 arg3 23 45"},
		{[]string{"12", "arg4", "23", "90"}, "12 arg4 23 90"},
		{[]string{"12", "arg5", "23", "19"}, "12 arg5 23 19"},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s", tt.input)
		t.Run(testName, func(t *testing.T) {
			result := efficientJoin(tt.input)
			if result != tt.want {
				t.Errorf("Get: %s Want: %s", result, tt.want)
			}
		})
	}
}

func BenchmarkEfficientJoin(b *testing.B) {
	args := []string{
		"command", "arg1", "arg2", "arg3", "command", "arg1", "arg2", "arg3",
		"command", "arg1", "arg2", "arg3", "command", "arg1", "arg2", "arg3",
	}

	for i := 0; i < b.N; i++ {
		result := efficientJoin(args)
		_ = result // Prevent compiler optimization
	}
}

func BenchmarkUnEfficientJoin(b *testing.B) {
	args := []string{
		"command", "arg1", "arg2", "arg3", "command", "arg1", "arg2", "arg3",
		"command", "arg1", "arg2", "arg3", "command", "arg1", "arg2", "arg3",
	}
	for i := 0; i < b.N; i++ {
		result := unEfficientJoin(args)
		_ = result
	}
}
