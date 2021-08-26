package internal

import "lisp_lambda-go/lambda/core/constants"

type WordTokenizer interface{
	TokenizeWord(word string) []string
}

type wordTokenizer struct {}

func New() *wordTokenizer {
	return &wordTokenizer{}
}

func (w *wordTokenizer) TokenizeWord(word string) []string {
	var tokens []string
	startingPos := 0

	for startingPos < len(word) {
		currentChar := word[startingPos]
		var token string
		if currentChar == constants.OPEN_TOKEN_CHAR {
			token = constants.OPEN_TOKEN_STR
		} else if currentChar == constants.CLOSE_TOKEN_CHAR {
			token = constants.CLOSE_TOKEN_STR
		} else {

		}
		startingPos += len(token)
		tokens = append(tokens, token)
	}
	return tokens
}

var _ WordTokenizer = &wordTokenizer{}
