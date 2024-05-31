package lexer

import (
	"regexp"
)

var TokenTypesMap map[string]*regexp.Regexp = map[string]*regexp.Regexp{
    "keyword_mod":     regexp.MustCompile(ModKeywordsRgx),
    "keyword_var":     regexp.MustCompile(VarKeywordsRgx),
    "keyword_obj":     regexp.MustCompile(ObjKeywordsRgx),
    "keyword_fun":     regexp.MustCompile(FunKeywordsRgx),
    "literal_number":  regexp.MustCompile(`\b\d+(\.\d+)?b?\b`),
    "literal_string":  regexp.MustCompile(`"([^"\\]|\\.)*"`),
    "literal_boolean": regexp.MustCompile(`\b(true|false)\b`),
    "identifier":      regexp.MustCompile(`\b&?[a-zA-Z_][a-zA-Z0-9_]*\b`),
    "operator":        regexp.MustCompile(`[+\-*/=]=?`),
    "delimiter":       regexp.MustCompile(`::|[(){};,:.]`),
    "whitespace":      regexp.MustCompile(`\s+`),
    "comment":         regexp.MustCompile(`#.*`),
}

var TokenTypesPriority = []string{
	"keyword_mod",
	"keyword_var",
	"keyword_obj",
	"keyword_fun",
	"literal_number",
	"literal_string",
	"literal_boolean",
	"identifier",
	"operator",
	"delimiter",
	"whitespace",
	"comment",
}

