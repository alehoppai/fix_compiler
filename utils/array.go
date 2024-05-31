package utils

import (
    "github.com/fix_compiler/lexer"
)

func Shift(tokens []lexer.Token) ([]lexer.Token, lexer.Token) {
    if len(tokens) == 0 {
        return tokens, lexer.Token{}
    }
    
    return tokens[1:], tokens[0]
}

func ShiftMut(tokens *[]lexer.Token) lexer.Token {
    if len(*tokens) == 0 {
        return lexer.Token{}
    }

    shiftedToken := (*tokens)[0]
    *tokens = (*tokens)[1:]

    return shiftedToken
}

func Pop(tokens []lexer.Token) ([]lexer.Token, lexer.Token) {
    if len(tokens) == 0 {
        return tokens, lexer.Token{}
    }

    return tokens[:len(tokens) - 1], tokens[len(tokens) - 1]
}

func PopMut(tokens *[]lexer.Token) lexer.Token {
    if len(*tokens) == 0 {
        return lexer.Token{}
    }

    poppedToken := (*tokens)[len(*tokens) - 1]
    *tokens = (*tokens)[:len(*tokens) - 1]

    return poppedToken
}
