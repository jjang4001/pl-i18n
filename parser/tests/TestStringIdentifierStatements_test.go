package tests

import (
	"testing"

	"pl-i18n/ast"
	"pl-i18n/lexer"
	"pl-i18n/parser"
)

func TestStringIdentifierStatements(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      interface{}
	}{
		{`string x = "hello world";`, "x", "hello world"},
		// {"let foobar = y;", "foobar", "y"},
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
		if !testStringIdStatement(t, stmt, tt.expectedIdentifier) {
			return
		}

		val := stmt.(*ast.StringIdStatement).Value
		if !testStringLiteralExpression(t, val, tt.expectedValue) {
			return
		}
	}
}

func testStringLiteralExpression(
	t *testing.T,
	exp ast.Expression,
	expected interface{},
) bool {
	switch v := expected.(type) {
	case string:
		return testStringLiteral(t, exp, v)
	}
	t.Errorf("type of exp not handled. got=%T", exp)
	return false
}
