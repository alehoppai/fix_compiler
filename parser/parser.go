package parser

import (
    "fmt"
    "github.com/fix_compiler/lexer"
    "github.com/fix_compiler/utils"
)

type AstNode struct  {
    Type       string
    Exports    bool 
    Mutable    bool
    DataType   string
    Identifier string
    Literal    string
}

func checkForModifiers(token *lexer.Token, tokens *[]lexer.Token, astNode *AstNode) bool {
    if token.Type == "keyword_mod" {
        // case with "mut" only
        if token.Value == "mut" {
            astNode.Mutable = true
            return true
        }

        // case with "pub mut"
        if token.Value == "pub" {
            astNode.Exports = true
            nextToken := utils.ShiftMut(tokens)

            if nextToken.Type == "keyword_mod" && nextToken.Value == "mut" {
                astNode.Mutable = true
            }

            return true
        }
    }

    return false
}

/**
 * token - current processing token - should be of type "keyword_var"
 * tokens - list of all remaning tokens - this should be sliced with approximation
 * e.g. num my_num = 43
 * [num - current] [my_num, =, 3] should be sliced from tokens
 */
func checkForVariable(
    token *lexer.Token,
    tokens *[]lexer.Token,
    astNode *AstNode,
) bool {
        if token.Type == "keyword_var" {
            astNode.Type = "VariableDefinition"
            astNode.DataType = token.Value

            varTokens := (*tokens)[:3]
            *tokens = (*tokens)[3:]

            if varTokens[0].Type == "identifier" {
                astNode.Identifier = varTokens[0].Value
            } else {
                fmt.Printf(
                    "Expected identifier after var type declaration, got %s: %s",
                    token.Type,
                    token.Value,
                )
                panic("Failed to parse tokens")
            }

            if varTokens[1].Type != "operator" {
                fmt.Printf(
                    "Expected operator after var identifier, got %s: %s",
                    token.Type,
                    token.Value,
                )
                panic("Failed to parse tokens")
            }

            if varTokens[2].Type == "literal_number" || varTokens[2].Type == "literal_string" || varTokens[2].Type == "literal_boolean" {
                astNode.Literal = varTokens[2].Value
            } else {
                fmt.Printf(
                    "Expected literal after var operator, got %s: %s",
                    token.Type,
                    token.Value,
                )
                panic("Failed to parse tokens")
            }
            return true
        }
        return false
}

func PrepareAST(tokens []lexer.Token) {
    AST := []AstNode{}
    astNode := AstNode{Exports: false, Mutable: false}

    for len(tokens) > 0 {
        token := utils.ShiftMut(&tokens)
        
        if checkForModifiers(&token, &tokens, &astNode) {
            continue
        }
        if checkForVariable(&token, &tokens, &astNode) {
            AST = append(AST, astNode)
            astNode = AstNode{Exports: false, Mutable: false}
            continue
        }

    }

    fmt.Println(AST)
}
