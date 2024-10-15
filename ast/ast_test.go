package ast

import (
    "monkey/token"
    "testing"
)

func TestSTring(t *testing.T) {
    program := &Program{
        statements: []Statement{
            &LetStatement{
                Token: token.Token{Type: token.LET, Literal: "let"},
                name: &Identifier{
                    Token: token.Token{ Type: token.IDENT, Literal: "myVar"},
                    Value: "myVar",
                },
                Value: &Identifier {
                    Token: token.Token{ Type: token.IDENT, Literal: "anotherVar"},
                    Value: "anotherVar",
                },
            },
        },
    }

    if program.String() != "let myVar = anotherVar;" {
        t.Errorf("program.String wrong got=%q", program.String())
    }
}

