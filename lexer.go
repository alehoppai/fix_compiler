package main

import (
    "regexp"
    "strings"
)

func getKeywords() []string {
    varKeywords := []string{"num", "str", "bool", "opt"}
    baseKeywords := []string{"struct", "proto", "pub", "mut", "fn", "mod", "return"}

    return PrependArray(varKeywords, baseKeywords)
}

func getTokenTypes() map[string]*regexp.Regexp {
    keywords := getKeywords()
    tokenTypes := make(map[string]*regexp.Regexp)

    keywordPattern := `\b(` + strings.Join(keywords, "|") + `)\b`
    tokenTypes["keyword"] = regexp.MustCompile(keywordPattern)
    tokenTypes["literal"] = regexp.MustCompile(`\b\d+(\.\d+)?b?\b`)
    tokenTypes["literalString"] = regexp.MustCompile(`"([^"\\]|\\.)*"`)
    tokenTypes["boolean"] = regexp.MustCompile(`\b(true|false)\b`)
    tokenTypes["identifier"] = regexp.MustCompile(`&?[a-zA-Z_][a-zA-Z0-9_]*`)
    tokenTypes["operator"] = regexp.MustCompile(`[+\-*/=]=?`)
    tokenTypes["delimiter"] = regexp.MustCompile(`[(){};,:.]`)
    tokenTypes["whitespace"] = regexp.MustCompile(`\s+`)
    tokenTypes["comment"] = regexp.MustCompile(`#.*`)

    return tokenTypes
}

type Token struct {
    Type string
    Value string
}

func Lexer(sourceCode string) []Token {
    var tokens []Token
    position := 0
    tokenTypes := getTokenTypes()
    

    for position < len(sourceCode) {
        var bufToken *Token

        for tokeType, regex := range tokenTypes {
            substring := sourceCode[:position]

            if regex.MatchString(substring) {
                bufToken = &Token{Type: tokeType, Value: }
            }
        } 
    }

    return tokens
}

