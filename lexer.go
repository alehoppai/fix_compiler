package main

import (
	"fmt"
	"log"
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
	Type  string
	Value string
}

func Lexer(sourceCode string) []Token {
	var tokens []Token
	position := 0
	tokenTypes := getTokenTypes()

	for position < len(sourceCode) {
		var bufToken *Token

		for tokenType, regex := range tokenTypes {
			substring := sourceCode[:position]
			match := regex.FindStringSubmatch(substring)

			fmt.Println(bufToken)
			fmt.Println(tokenType)
			fmt.Println(position)
			fmt.Println(match)
			fmt.Println("========")

			if len(match) > 0 {
				bufToken = &Token{Type: tokenType, Value: match[0]}
			}
			//if regex.MatchString(substring) {
			//    bufToken = &Token{Type: tokeType, Value: }
			//}
		}

		if bufToken != nil {
			if bufToken.Type != "comment" && bufToken.Type != "whitespace" {
				token := Token{Type: bufToken.Type, Value: bufToken.Value}
				tokens = append(tokens, token)
				bufToken = nil
			}
			position += len(bufToken.Value)
		} else {
			log.Fatal("Not a single substring matched token")
			position += 1
			continue
		}
	}

	return tokens
}
