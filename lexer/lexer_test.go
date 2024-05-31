package lexer

import (
	"testing"
)

func TestLexerVars(t *testing.T) {
	sourceCode := `num my_num = 42
num my_bigint = 1234567890123456789012345678901234567890b
pub str my_string = "Hello, world!"
bool my_bool = false`

	expectedTokens := []Token{
		{Type: "keyword_var",     Value: "num"},
		{Type: "identifier",      Value: "my_num"},
		{Type: "operator",        Value: "="},
		{Type: "literal_number",  Value: "42"},
		{Type: "keyword_var",     Value: "num"},
		{Type: "identifier",      Value: "my_bigint"},
		{Type: "operator",        Value: "="},
		{Type: "literal_number",  Value: "1234567890123456789012345678901234567890b"},
		{Type: "keyword_mod",     Value: "pub"},
		{Type: "keyword_var",     Value: "str"},
		{Type: "identifier",      Value: "my_string"},
		{Type: "operator",        Value: "="},
		{Type: "literal_string",  Value: `"Hello, world!"`},
		{Type: "keyword_var",     Value: "bool"},
		{Type: "identifier",      Value: "my_bool"},
		{Type: "operator",        Value: "="},
		{Type: "literal_boolean", Value: "false"},
	}

	tokens := Tokenize(sourceCode)

	if len(tokens) != len(expectedTokens) {
		t.Fatalf("expected %d tokens, got %d", len(expectedTokens), len(tokens))
	}

	for i, token := range tokens {
		if token != expectedTokens[i] {
			t.Errorf("expected token %v, got %v", expectedTokens[i], token)
		}
	}
}

func TestLexerStructs(t *testing.T) {
	sourceCode := `struct {
	num age
	str name
} DataT

proto {
	# instance method
	pub fn &introduce () {
		print("Hello, I am !{&.name}, and I'm !{&.name} y.o.")
	}
	# static method
	pub fn new (num age, str name) DataT {
		Data(age: age, name: name)
	}
} DataT`

	expectedTokens := []Token{
        {Type: "keyword_obj",    Value: "struct"},
        {Type: "delimiter",      Value: "{"},
        {Type: "keyword_var",    Value: "num"},
        {Type: "identifier",     Value: "age"},
        {Type: "keyword_var",    Value: "str"},
        {Type: "identifier",     Value: "name"},
        {Type: "delimiter",      Value: "}"},
        {Type: "identifier",     Value: "DataT"},
        {Type: "keyword_obj",    Value: "proto"},
        {Type: "delimiter",      Value: "{"},
        {Type: "keyword_mod",    Value: "pub"},
        {Type: "keyword_fun",    Value: "fn"},
        {Type: "identifier",     Value: "introduce"},
        {Type: "delimiter",      Value: "("},
        {Type: "delimiter",      Value: ")"},
        {Type: "delimiter",      Value: "{"},
        {Type: "identifier",     Value: "print"},
        {Type: "delimiter",      Value: "("},
        {Type: "literal_string", Value: `"Hello, I am !{&.name}, and I'm !{&.name} y.o."`},
        {Type: "delimiter",      Value: ")"},
        {Type: "delimiter",      Value: "}"},
        {Type: "keyword_mod",    Value: "pub"},
        {Type: "keyword_fun",    Value: "fn"},
        {Type: "identifier",     Value: "new"},
        {Type: "delimiter",      Value: "("},
        {Type: "keyword_var",    Value: "num"},
        {Type: "identifier",     Value: "age"},
        {Type: "delimiter",      Value: ","},
        {Type: "keyword_var",    Value: "str"},
        {Type: "identifier",     Value: "name"},
        {Type: "delimiter",      Value: ")"},
        {Type: "identifier",     Value: "DataT"},
        {Type: "delimiter",      Value: "{"},
        {Type: "identifier",     Value: "Data"},
        {Type: "delimiter",      Value: "("},
        {Type: "identifier",     Value: "age"},
        {Type: "delimiter",      Value: ":"},
        {Type: "identifier",     Value: "age"},
        {Type: "delimiter",      Value: ","},
        {Type: "identifier",     Value: "name"},
        {Type: "delimiter",      Value: ":"},
        {Type: "identifier",     Value: "name"},
        {Type: "delimiter",      Value: ")"},
        {Type: "delimiter",      Value: "}"},
        {Type: "delimiter",      Value: "}"},
        {Type: "identifier",     Value: "DataT"},
	}

	tokens := Tokenize(sourceCode)

	if len(tokens) != len(expectedTokens) {
		t.Fatalf("expected %d tokens, got %d", len(expectedTokens), len(tokens))
	}

	for i, token := range tokens {
		if token != expectedTokens[i] {
			t.Errorf("expected token %v, got %v", expectedTokens[i], token)
		}
	}
}
