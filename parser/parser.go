package parser

import (
    "fmt"
    "github.com/fix_compiler/lexer"
    "github.com/fix_compiler/utils"
)

func PrepareAST(tokens []lexer.Token) []AstNode {
    AST := []AstNode{}

    for len(tokens) > 0 {
        astNode := AstNode{Exports: false, Mutable: false}
        token := utils.ShiftMut(&tokens)

        for ParseModifiers(&token, &astNode) {
            if len(tokens) > 0 {
                token = utils.ShiftMut(&tokens)
            }
        }

        if ParseFunctionDefinition(&token, &tokens, &astNode) ||
            ParseVariableDeclaration(&token, &tokens, &astNode) ||
            ParseVariableDeclaration(&token, &tokens, &astNode) {
            AST = append(AST, astNode)
            continue
        }

        fmt.Printf("Unexpected token type: %s, value: %s\n", token.Type, token.Value)
    }

    return AST
}
