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
    checkParserErrors(t, parser)

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
            break
        }
    }
}

func checkParserErrors(t *testing.T, p *Parser) {
    errors := p.Errors()
    if len(errors) == 0 {
        return
    }
    t.Errorf("parser has %d errors", len(errors))
    for _, msg := range errors {
        t.Errorf("parser error: %q", msg)
    }
    t.FailNow()
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

func TestReturnStatements(t *testing.T) {
    input := `
        return 5;
        return 10;
        return 993322;
    `
    l := lexer.New(input)
    p := New(l)
    program := p.ParseProgram()
    checkParserErrors(t, p)

    if len(program.Statements) != 3 {
        t.Fatalf("program.Statements expected 3 statements. got=%d",
            len(program.Statements))
    }

    for _, stmt := range program.Statements {
        returnStmt, ok := stmt.(*ast.ReturnStatement)
        if !ok {
            t.Errorf("stmt expected *ast.ReturnStatement. got=%T", stmt)
            continue
        }
        if returnStmt.TokenLiteral() != "return" {
            t.Errorf("returnStmt.TokenLiteral expected 'return', got %q",
                returnStmt.TokenLiteral())
        }
    }
}
