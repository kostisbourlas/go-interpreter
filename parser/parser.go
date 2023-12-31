package parser

import (
    "fmt"

    "github.com/kostisbourlas/go-interpreter/ast"
    "github.com/kostisbourlas/go-interpreter/lexer"
    "github.com/kostisbourlas/go-interpreter/token"
)

type Parser struct {
    lexer *lexer.Lexer

    curToken token.Token
    peekToken token.Token
    errors []string
}

func New(l *lexer.Lexer) *Parser {
    p := &Parser{
        lexer: l,
        errors: []string{},
    }
    p.NextToken()
    p.NextToken()

    return p
}

func (p *Parser) Errors() []string {
    return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
    msg := fmt.Sprintf("expected next token to be %s, got %s", t, p.peekToken.Type)
    p.errors = append(p.errors, msg)
}

func (p *Parser) NextToken() {
    p.curToken = p.peekToken
    p.peekToken = p.lexer.NextToken()
}

func (p *Parser) parseStatement() ast.Statement {
    switch p.curToken.Type {
    case token.LET:
        return p.parseLetStatement()
    case token.RETURN:
        return p.parseReturnStatement()
    default:
        return nil
    }
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
    stmt := &ast.LetStatement{Token: p.curToken} 

    if !p.expectPeek(token.IDENT) {
        return nil
    }

    stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
    
    if !p.expectPeek(token.ASSIGN) {
        return nil
    }

    // TODO: We are skipping the expressions until we
    // encounter a semicolon
    for !p.curTokenIs(token.SEMICOLON) {
        p.NextToken()
    }

    return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
    stmt := &ast.ReturnStatement{Token: p.curToken}
    p.NextToken()

    // TODO: We are skipping the expressions until we
    // encounter a semicolon
    for !p.curTokenIs(token.SEMICOLON) {
        p.NextToken()
    }

    return stmt
}

func (p *Parser) expectPeek(t token.TokenType) bool {
    if p.peekTokenIs(t) {
        p.NextToken()
        return true
    }
    p.peekError(t)
    return false
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
   return p.peekToken.Type == t 
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
   return p.curToken.Type == t 
}

func (p *Parser) ParseProgram() *ast.Program {
    program := &ast.Program{}
    program.Statements = []ast.Statement{}

    for p.curToken.Type != token.EOF {
        stmt := p.parseStatement()
        if stmt != nil {
            program.Statements = append(program.Statements, stmt)
        }
        p.NextToken()
    }
    return program
}
