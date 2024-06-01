package parser

import (
    "fmt"
	"github.com/fix_compiler/lexer"
	"github.com/fix_compiler/utils"
)

func parseFunctionParameters(tokens *[]lexer.Token) []AstNode {
    parameters := []AstNode{}

    for len(*tokens) > 0 {
        paramToken := utils.ShiftMut(tokens)
        if paramToken.Type == "delimiter" && paramToken.Value == ")" {
            break
        }

        paramNode := AstNode{}

        if paramToken.Type == "keyword_var" {
            paramNode.DataType = ParseDataType(paramToken.Value)
        } else {
            panic(
                fmt.Sprintf(
                    "Expected data type for parameter, got %s: %s",
                    paramToken.Type,
                    paramToken.Value,
                ),
            )
        }

        paramToken = utils.ShiftMut(tokens)
        if paramToken.Type == "identifier" {
            paramNode.Identifier = paramToken.Value
        } else {
            panic(
                fmt.Sprintf(
                    "Expected identifier for parameter, got %s: %s",
                    paramToken.Type,
                    paramToken.Value,
                ),
            )
        }

        paramNode.Type = FunctionParamDeclaration
        parameters = append(parameters, paramNode)

        if len(*tokens) > 0 {
            nextToken := (*tokens)[0]
            if nextToken.Type == "delimiter" && nextToken.Value == "," {
                utils.ShiftMut(tokens)
            } else if nextToken.Type == "delimiter" && nextToken.Value == ")" {
                break
            }
        }
    }

    return parameters
}

func parseFunctionBody(tokens *[]lexer.Token) []AstNode {
    body := []AstNode{}

    for len(*tokens) > 0 {
        token := utils.ShiftMut(tokens)
        if token.Type == "delimiter" && token.Value == "}" {
            break
        }

        bodyNode := AstNode{}
        if ParseVariableDeclaration(&token, tokens, &bodyNode) || ParseVariableAssignment(&token, tokens, &bodyNode) {
            body = append(body, bodyNode)
            continue
        }

        panic(
            fmt.Sprintf(
                "Unexpected token in function body: %s: %s",
                token.Type,
                token.Value,
            ),
        )
    }

    return body
}

func ParseFunctionDefinition(token *lexer.Token, tokens *[]lexer.Token, astNode *AstNode) bool {
    if token.Type != "keyword_fun" {
        return false
    }

    if len(*tokens) < 3 {
        panic("Not enough tokens to parse function definition")
    }

    astNode.Type = FunctionDeclaration

    nameToken := utils.ShiftMut(tokens)
    if nameToken.Type != "identifier" {
        panic(
            fmt.Sprintf(
                "Expected function name, got %s: %s",
                nameToken.Type,
                nameToken.Value,
            ),
        )
    }
    astNode.Identifier = nameToken.Value

    paramsToken := utils.ShiftMut(tokens)
    if paramsToken.Type != "delimiter" || paramsToken.Value != "(" {
        panic(
            fmt.Sprintf(
                "Expected '(', got %s: %s",
                paramsToken.Type,
                paramsToken.Value,
            ),
        )
    }

    astNode.Parameters = parseFunctionParameters(tokens)
    astNode.Body = parseFunctionBody(tokens)

    return true
}
