package lexer

import (
    "strings"
)

var ModKeywordsArr []string = []string{"pub", "mut", "mod"}
var VarKeywordsArr []string = []string{"num", "str", "bool", "opt"}
var ObjKeywordsArr []string = []string{"obj", "proto"}
var FunKeywordsArr []string = []string{"fn", "return"}

var ModKeywordsRgx = `\b(` + strings.Join(ModKeywordsArr, "|") + `)\b` 
var VarKeywordsRgx = `\b(` + strings.Join(VarKeywordsArr, "|") + `)\b` 
var ObjKeywordsRgx = `\b(` + strings.Join(ObjKeywordsArr, "|") + `)\b` 
var FunKeywordsRgx = `\b(` + strings.Join(FunKeywordsArr, "|") + `)\b` 
