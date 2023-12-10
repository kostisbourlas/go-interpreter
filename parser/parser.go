package parser

import (
	"github.com/kostisbourlas/go-interpreter/ast"
	"github.com/kostisbourlas/go-interpreter/lexer"
	"github.com/kostisbourlas/go-interpreter/token"
)

type Parser struct {
    lexer *lexer.Lexer

    curToken token.Token
    peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
    p := &Parser{lexer: l}
    p.NextToken()
    p.NextToken()

    return p
}

func (p *Parser) NextToken() {
    p.curToken = p.peekToken
    p.peekToken = p.lexer.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
    return nil
}
