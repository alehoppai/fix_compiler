package lexer

type Token struct {
	Type  string
	Value string
}

func checkForToken(substring string) (bool, string, string)  {
    for _, tokenType := range TokenTypesPriority {
		regex := TokenTypesMap[tokenType]
		if loc := regex.FindStringIndex(substring); loc != nil && loc[0] == 0 {
			match := substring[:loc[1]]
			return true, tokenType, match
		}
	}
	return false, "", ""
}

func Tokenize(sourceCode string) []Token {
	var tokens []Token
	position := 0

	for position < len(sourceCode) {
        substring := sourceCode[position:]
        match, tokenType, tokenValue := checkForToken(substring)

		if match {
			if tokenType != "comment" && tokenType != "whitespace" {
				token := Token{Type: tokenType, Value: tokenValue}
				tokens = append(tokens, token)
			}
			position += len(tokenValue)
		} else {
			position++
		}
	}

	return tokens
}
