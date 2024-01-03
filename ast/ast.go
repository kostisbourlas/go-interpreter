package ast

import (
    "bytes"

    "github.com/kostisbourlas/go-interpreter/token"
)

type Node interface {
    TokenLiteral() string
    String() string
} 

type Statement interface {
    Node
    statementNode()
}

type Expression interface {
    Node
    expressionNode()
}

type Identifier struct {
    Token token.Token // the token.IDENT token
    Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
    return i.Token.Literal
}

func (i *Identifier) String() string {
   return i.Value 
}

type LetStatement struct {
    Token token.Token // the token.LET token
    Name *Identifier
    Value Expression
}

func (ls *LetStatement) statementNode() {}

func (ls *LetStatement) TokenLiteral() string {
    return ls.Token.Literal
}

func (ls *LetStatement) String() string {
    var buf bytes.Buffer

    buf.WriteString(ls.TokenLiteral() + " ")
    buf.WriteString(ls.Name.String())
    buf.WriteString(" = ")
    if ls.Value != nil {
        buf.WriteString(ls.Value.String())
    }
    buf.WriteString(";")

    return buf.String()
}

type ReturnStatement struct {
    Token token.Token // the 'return' token
    ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) TokenLiteral() string {
    return rs.Token.Literal
}

func (rs *ReturnStatement) String() string {
    var buf bytes.Buffer

    buf.WriteString(rs.TokenLiteral() + " ")
    if rs.ReturnValue != nil {
        buf.WriteString(rs.ReturnValue.String())
    }
    buf.WriteString(";")

    return buf.String()
}

type ExpressionStatement struct {
    Token token.Token
    Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

func (es *ExpressionStatement) TokenLiteral() string {
    return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
    if es.Expression != nil {
        return es.Expression.String()
    }
    return ""
}

type Program struct {
    Statements []Statement
}

func (p *Program) TokenLiteral() string {
    if len(p.Statements) <= 0 {
        return ""
    }
    return p.Statements[0].TokenLiteral()
}

func (p *Program) String() string {
    var buf bytes.Buffer

    for _, s := range p.Statements {
        buf.WriteString(s.String())
    }

    return buf.String()
}
