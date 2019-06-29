package tests

import (
	"pl-i18n/lexer"
	"pl-i18n/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `정수 five = 5
	부동소수 ten = 10.0
	문자 hi = "hello"
	불 a = true
	if (5 < 10) {
		산출 진실;
	} else {
		산출 거짓;
	}
	while (1 < 2) {
		정수 i = 1
	}
	for (;;) {
		정수 i = 1
	}
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.INTID, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		// {token.SEMICOLON, ";"},
		{token.FLOATID, "ten"},
		{token.ASSIGN, "="},
		{token.FLOAT, "10.0"},
		// {token.SEMICOLON, ";"},
		{token.STRINGID, "hi"},
		{token.ASSIGN, "="},
		{token.STRING, "hello"},
		// {token.SEMICOLON, ";"},
		{token.BOOLID, "a"},
		{token.ASSIGN, "="},
		{token.TRUE, "true"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.WHILE, "while"},
		{token.LPAREN, "("},
		{token.INT, "1"},
		{token.LT, "<"},
		{token.INT, "2"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.INTID, "i"},
		{token.ASSIGN, "="},
		{token.INT, "1"},
		{token.RBRACE, "}"},
		{token.FOR, "for"},
		{token.LPAREN, "("},
		{token.SEMICOLON, ";"},
		{token.SEMICOLON, ";"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.INTID, "i"},
		{token.ASSIGN, "="},
		{token.INT, "1"},
		{token.RBRACE, "}"},
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
