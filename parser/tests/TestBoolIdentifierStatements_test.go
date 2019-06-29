package tests

import (
	"testing"

	"pl-i18n/ast"
	"pl-i18n/lexer"
	"pl-i18n/parser"
)

func TestBoolIdentifierStatements(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      bool
	}{
		{"bool foobar = true", "foobar", true},
		{"bool a = false", "a", false},
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
		if !testBoolIdStatement(t, stmt, tt.expectedIdentifier) {
			return
		}

		val := stmt.(*ast.BoolIdStatement).Value
		if !testBoolLiteralExpression(t, val, tt.expectedValue) {
			return
		}
	}
}
