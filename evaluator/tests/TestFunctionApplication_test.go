package tests

import (
	"testing"
)

func TestFunctionApplication(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"int identity = func(x) { x; }; identity(5);", 5},
		{"int identity = func(x) { return x; }; identity(5);", 5},
		{"int double = func(x) { x * 2; }; double(5);", 10},
		{"int add = func(x, y) { x + y; }; add(5, 5);", 10},
		{"int add = func(x, y) { x + y; }; add(5 + 5, add(5, 5));", 20},
		{"func(x) { x; }(5)", 5},
	}
	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}
