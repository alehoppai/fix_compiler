package parser

import (
    "fmt"
    "github.com/fix_compiler/lexer"
    "github.com/fix_compiler/utils"
)

func ParseVariableDeclaration(
    token *lexer.Token,
    tokens *[]lexer.Token,
    astNode *AstNode,
) bool {
    if token.Type != "keyword_var" {
        return false
    }

    astNode.Type = VariableDeclaration
    astNode.DataType = ParseDataType(token.Value)

    if len(*tokens) < 2 {
        panic("Not enough tokens to parse variable declaration")
    }

    identifierToken := utils.ShiftMut(tokens)
    if identifierToken.Type !=  "identifier" {
        panic(
            fmt.Sprintf(
                "Expected identifier after var type declaration, got %s: %s",
                identifierToken.Type,
                identifierToken.Value,
            ),
        )
    }
    astNode.Identifier = identifierToken.Value

    operatorToken := utils.ShiftMut(tokens) 
    if operatorToken.Type !=  "operator" {
        panic(
            fmt.Sprintf(
                "Expected operator after var identifier, got %s: %s",
                operatorToken.Type,
                operatorToken.Value,
            ),
        )
    }

    astNode.Expression = ParseExpression(tokens)
    return true
}

func ParseVariableAssignment(token *lexer.Token, tokens *[]lexer.Token, astNode *AstNode) bool {
    if token.Type != "identifier" {
        return false
    }

    astNode.Type = VariableAssignment
    astNode.Identifier = token.Value

    operatorToken := utils.ShiftMut(tokens)
    if operatorToken.Type != "operator" || operatorToken.Value != "=" {
        return false
    }

    astNode.Expression = ParseExpression(tokens)
    return true
}
