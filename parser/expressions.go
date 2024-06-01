package parser

import (
    "fmt"
    "github.com/fix_compiler/lexer"
    "github.com/fix_compiler/utils"
)

func ParseExpression(tokens *[]lexer.Token) *ExpressionNode {
    if len(*tokens) == 0 {
        panic("No tokens to parse for expression")
    }

    // Parse the left operand
    leftToken := utils.ShiftMut(tokens)
    leftNode := &ExpressionNode{
        Type:     IdentifierExpression,
        DataType: Opt,
        Value:    leftToken.Value,
    }

    if leftToken.Type == "identifier" {
        if len(*tokens) > 0 {
            nextToken := (*tokens)[0]
            if nextToken.Type == "delimiter" && nextToken.Value == "(" {
                fmt.Println("PARSING FUNCTION CALL")
                // Function call
                utils.ShiftMut(tokens) // Consume '('
                leftNode.Type = IdentifierExpression
                leftNode.Left = parseFunctionCall(leftToken.Value, tokens)
                return leftNode
            }
        }
    }

    if len(*tokens) == 0 {
        return leftNode
    }

    // Check if there's an operator next
    operatorToken := utils.ShiftMut(tokens)
    if operatorToken.Type != "operator" {
        // If no operator, the expression is just a single literal or identifier
        *tokens = append([]lexer.Token{operatorToken}, *tokens...)
        return leftNode
    }

    // Parse the right operand
    rightNode := ParseExpression(tokens)

    // Return the combined expression node
    return &ExpressionNode{
        Type:     BinaryExpression,
        DataType: Opt,
        Left:     leftNode,
        Right:    rightNode,
        Operator: operatorToken.Value,
    }
}

func parseFunctionCall(identifier string, tokens *[]lexer.Token) *ExpressionNode {
    args := []ExpressionNode{}
    for len(*tokens) > 0 {
        token := utils.ShiftMut(tokens)
        if token.Type == "delimiter" && token.Value == ")" {
            break
        }
        if token.Type == "delimiter" && token.Value == "," {
            continue
        }
        exprNode := ParseExpression(&[]lexer.Token{token})
        args = append(args, *exprNode)
    }

    // Create a function call expression node
    return &ExpressionNode{
        Type:     IdentifierExpression,
        DataType: Opt,
        Value:    identifier,
        Left:     nil,
        Right:    nil,
        Operator: "",
    }
}

