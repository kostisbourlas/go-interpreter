package ast

import (
    "testing"

    "github.com/kostisbourlas/go-interpreter/token"
)


func TestString(t *testing.T) {
    program := &Program{
        Statements: []Statement{
            &LetStatement{
                Token: token.Token{Type: token.LET, Literal: "let"},
                Name: &Identifier{
                    Token: token.Token{Type: token.IDENT, Literal: "myVar"},
                    Value: "myVar",
                },
                Value: &Identifier{
                    Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
                    Value: "anotherVar",
                },
            },
        },
    }

    testExpression := "let myVar = anotherVar;"
    if program.String() != testExpression {
        t.Errorf("program.String() expected `%q`. got=%q", testExpression, program.String())
    }
}
