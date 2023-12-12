package parser

import (
	"testing"

	"github.com/kostisbourlas/go-interpreter/ast"
	"github.com/kostisbourlas/go-interpreter/lexer"
) 

func TestLetStatements(t *testing.T) {
    input := `
    let x = 5;
    let y = 10;
    let foobar = 838383;
    `
    lexer := lexer.New(input)
    parser := New(lexer)

    program := parser.ParseProgram()
    if program == nil {
        t.Fatalf("ParseProgram() returned nil")
    }
    if len(program.Statements) != 3 {
        t.Fatalf("program.Statements expected 3 statements. got=%d",
            len(program.Statements))
    }
    tests := []struct {
        expectedIdentifier string
    }{
        {"x"},
        {"y"},
        {"foobar"},
    }
    
    for i, tt := range tests {
        stmt := program.Statements[i]
        if !assertLetStatement(t, stmt, tt.expectedIdentifier) {
            return
        }
    }
}

func assertLetStatement(t *testing.T, s ast.Statement, name string) bool {
    if s.TokenLiteral() != "let" {
        t.Errorf("s.TokenLiteral expected 'let' literal. got=%s", 
            s.TokenLiteral())
        return false
    }

    letStmt, ok := s.(*ast.LetStatement)
    if !ok {
        t.Errorf("s expected type *ast.LetStatement. got=%T", s)
        return false
    }

    if letStmt.Name.Value != name {
        t.Errorf("letStmt.Name.Value expected '%s'. got=%s",
            name, letStmt.Name.Value)
        return false
    }
    
    if letStmt.Name.TokenLiteral() != name {
        t.Errorf("s.Name expected '%s'. got=%s", name, letStmt.Name)
        return false
    }
    return true
}
