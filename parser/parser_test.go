package parser

import (
	"github.com/leonardchinonso/leon/ast"
	"github.com/leonardchinonso/leon/lexer"
	"testing"
)

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d error(s)\n", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q\n", msg)
	}

	t.FailNow()
}

func TestVarStatements(t *testing.T) {
	input := `
var x = 5;
var y = 10;
var foobar = 215723;
`

	l := lexer.NewLexer(input)
	p := NewParser(l)

	prog := p.ParseProgram()
	checkParserErrors(t, p)
	if prog == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(prog.Statements) != 3 {
		t.Fatalf("prog.Statements does not contain 3 statements. got=%d", len(prog.Statements))
	}

	type tester struct {
		expectedIdentifier string
	}

	tests := []tester{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tst := range tests {
		stmt := prog.Statements[i]
		if !testVarStatement(t, stmt, tst.expectedIdentifier) {
			return
		}
	}
}

func testVarStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "var" {
		t.Errorf("s.TokenLiteral not 'var', got=%T", s)
		return false
	}

	varStmt, ok := s.(*ast.VarStatement)
	if !ok {
		t.Errorf("s not *ast.VarStatement, got=%T", s)
		return false
	}

	if varStmt.Name.Value != name {
		t.Errorf("varStmt.Name.Value not '%s', got =%s", name, varStmt.Name.Value)
		return false
	}

	if varStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s', got =%s", name, varStmt.Name)
		return false
	}

	return true
}

func TestReturnStatements(t *testing.T) {
	input := `
return 5;
return 10;
return 102434;
`

	l := lexer.NewLexer(input)
	p := NewParser(l)

	prog := p.ParseProgram()
	checkParserErrors(t, p)
	if prog == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(prog.Statements) != 3 {
		t.Fatalf("prog.Statements does not contain 3 statements. got=%d", len(prog.Statements))
	}

	for _, stmt := range prog.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.returnStatement, got=%T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q", returnStmt.TokenLiteral())
		}
	}

}
