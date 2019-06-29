package tests

import (
	"testing"

	"pl-i18n/ast"
	"pl-i18n/lexer"
	"pl-i18n/parser"
)

func TestBoolLiteralExpression(t *testing.T) {
	input := `true`
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	stmt := program.Statements[0].(*ast.ExpressionStatement)
	literal, ok := stmt.Expression.(*ast.BooleanLiteral)
	if !ok {
		t.Fatalf("exp not *ast.StringLiteral. got=%T", stmt.Expression)
	}
	if literal.Value != true {
		t.Errorf("literal.Value not %t. got=%t", true, literal.Value)
	}
}
