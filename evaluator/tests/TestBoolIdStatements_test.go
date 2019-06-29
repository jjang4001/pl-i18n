package tests

import (
	"testing"
)

func TestBoolIdStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{`bool hi = true; hi;`, true},
		{"bool bye = false; bye;", false},
	}
	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBoolObject(t, evaluated, tt.expected)
	}
}
