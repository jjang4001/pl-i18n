package tests

// import (
// 	"testing"
// )

// func TestEvalFloatExpression(t *testing.T) {
// 	tests := []struct {
// 		input    string
// 		expected float64
// 	}{
// 		{"5.1", float64(5.1)},
// 		{"10.1", float64(10.1)},
// 		{"-5.3", float64(-5.3)},
// 		{"-10.01", float64(-10.01)},
// 		{"5.1 + 5.1 + 5.1 + 5.1 - 10.1", float64(10.3)},
// 		{"2.0 * 2.0 * 2.0 * 2.0 * 2.0", float64(32.0)},
// 		{"-50.1 + 100.2 + -50.1", float64(0)},
// 		{"5.5 * 2.0 + 10.0", float64(21.0)},
// 		{"5.0 + 2.0 * 10.1", float64(25.2)},
// 		{"20.0 + 2.0 * -10.1", float64(-0.2)},
// 		{"50.0 / 2.0 * 2.0 + 10.1", float64(60.1)},
// 		{"2.0 * (5.1 + 10.2)", float64(30.6)},
// 		{"3.0 * 3.0 * 3.1 + 10.0", float64(37.9)},
// 		{"3.1 * (3.0 * 3.0) + 10.0", float64(37.3)},
// 		{"(5.0 + 10.0 * 2.0 + 15.5 / 3.1) * 2.0 + -10.0", float64(50)},
// 	}

// 	for _, tt := range tests {
// 		evaluated := testEval(tt.input)
// 		testFloatObject(t, evaluated, tt.expected)
// 	}
// }
