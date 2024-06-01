package parser

type DataType int
const (
    Num DataType = iota
    Str
    Bool
    BigInt
    Opt
    Obj
)
func (d DataType) String() string {
    return [...]string{"num", "str", "bool", "bigint", "opt"}[d]
}
func ParseDataType(value string) DataType {
    switch value {
    case "num":
        return Num
    case "str":
        return Str
    case "bool":
        return Bool
    case "bigint":
        return BigInt
    case "obj":
        return Obj
    default:
        return Opt
    }
}

type ExpressionType int
const (
    BinaryExpression ExpressionType = iota
    LiteralExpression
    IdentifierExpression
)

type ExpressionNode struct {
    Type     ExpressionType
    DataType DataType
    Value    string
    Left     *ExpressionNode
    Right    *ExpressionNode
    Operator string
}


type NodeType int
const (
    VariableDeclaration NodeType = iota
    VariableAssignment
    FunctionDeclaration
    FunctionParamDeclaration
    ObjectDeclaration
    ProtoDeclaration
)
func (n NodeType) String() string {
    return [...]string{"VariableDeclaration", "VariableAssignment", "FunctionDeclaration", "ObjectDeclaration", "ProtoDeclaration"}[n]
}

type AstNode struct  {
    Type       NodeType
    Exports    bool 
    Mutable    bool
    DataType   DataType
    Identifier string
    Literal    string
    Expression *ExpressionNode
    Parameters []AstNode
    Body       []AstNode
}
