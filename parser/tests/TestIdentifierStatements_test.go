package tests

import (
	"testing"

	"pl-i18n/ast"
	"pl-i18n/lexer"
	"pl-i18n/parser"
)

func TestIntIdentifierStatements(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      interface{}
	}{
		{"int x = 5", "x", 5},
		{"int hello = 80", "hello", 80},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := parser.New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statements. got=%d",
				len(program.Statements))
		}

		stmt := program.Statements[0]
		if !testIntIdStatement(t, stmt, tt.expectedIdentifier) {
			return
		}

		val := stmt.(*ast.IntIdStatement).Value
		if !testLiteralExpression(t, val, tt.expectedValue) {
			return
		}
	}
}
