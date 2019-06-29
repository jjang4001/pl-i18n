package tests

import (
	"testing"
)

func TestStringIdStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`string hi = "hello"; hi;`, "hello"},
	}
	for _, tt := range tests {
		testStringObject(t, testEval(tt.input), tt.expected)
	}
}
