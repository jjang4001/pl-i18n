package tests

import (
	"pl-i18n/lexer"
	"pl-i18n/token"
	"testing"
)

func TestBasic(t *testing.T) {
	input := `정수 five = 5`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.INTID, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.EOF, ""},
	}

	l := lexer.New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
