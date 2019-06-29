package tests

import (
	"testing"
)

func TestIntIdStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"int a = 5 \n a", 5},
	}
	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}
