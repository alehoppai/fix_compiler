package parser

import (
    "github.com/fix_compiler/lexer"
)

func ParseModifiers(token *lexer.Token, astNode *AstNode) bool {
    if token.Type == "keyword_mod" {
        switch token.Value {
        case "pub":
            astNode.Exports = true
        case "mut":
            astNode.Mutable = true
        default:
            return false
        }
        return true
    }
    return false
}

