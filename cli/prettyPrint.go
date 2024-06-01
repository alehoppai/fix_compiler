package cli

import (
	"fmt"

	"github.com/fix_compiler/lexer"
	"github.com/fix_compiler/parser"
)

func PrettyPrintAST(node parser.AstNode, indent int) {
	fmt.Println("============")
	indentString := ""
	for i := 0; i < indent; i++ {
		indentString += "|  "
	}

	fmt.Printf("%sNode Type: %s\n", indentString, node.Type.String())
	fmt.Printf("%sExports: %v\n", indentString, node.Exports)
	fmt.Printf("%sMutable: %v\n", indentString, node.Mutable)
	fmt.Printf("%sDataType: %s\n", indentString, node.DataType.String())
	fmt.Printf("%sIdentifier: %s\n", indentString, node.Identifier)
	fmt.Printf("%sLiteral: %s\n", indentString, node.Literal)

	if node.Expression != nil {
		fmt.Printf("%sExpression: \n", indentString)
		prettyPrintExpression(*node.Expression, indent+1)
	}

	if len(node.Parameters) > 0 {
		fmt.Printf("%sParameters:\n", indentString)
		for _, param := range node.Parameters {
			PrettyPrintAST(param, indent+1)
		}
	}

	if len(node.Body) > 0 {
		fmt.Printf("%sBody:\n", indentString)
		for _, bodyNode := range node.Body {
			PrettyPrintAST(bodyNode, indent+1)
		}
	}
}

func prettyPrintExpression(expr parser.ExpressionNode, indent int) {
	indentString := ""
	for i := 0; i < indent; i++ {
		indentString += "|  "
	}

	fmt.Printf("%sExpression Type: %d\n", indentString, expr.Type)
	fmt.Printf("%sDataType: %s\n", indentString, expr.DataType.String())
	fmt.Printf("%sValue: %s\n", indentString, expr.Value)
	if expr.Left != nil {
		fmt.Printf("%sLeft:\n", indentString)
		prettyPrintExpression(*expr.Left, indent+1)
	}
	if expr.Right != nil {
		fmt.Printf("%sRight:\n", indentString)
		prettyPrintExpression(*expr.Right, indent+1)
	}
	fmt.Printf("%sOperator: %s\n", indentString, expr.Operator)
}

func PrettyPrintTokens(tokens *[]lexer.Token) {
	for i, token := range *tokens {
		fmt.Printf("%d) type: %s, value: %s\n", i, token.Type, token.Value)
	}
}
