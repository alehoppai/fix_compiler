package parser

import (
    "fmt"
    "github.com/fix_compiler/lexer"
    "github.com/fix_compiler/utils"
)

type AstNode struct  {
    Exports    bool 
    Mutable    bool
    Type       string
    Identifier string
    Literal    string
}

func parseVariable() {

}

func PrepareAST(tokens []lexer.Token) {
    for len(tokens) > 0 {
        token := utils.ShiftMut(&tokens)
        fmt.Println(token)
    }
}
