package impl

import (
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/tokenizer/impl/internal"
	"strings"
)

type wordTokenizer struct {}

func NewWordTokenizer() *wordTokenizer {
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
			subStr := word[startingPos:]
			openPos := strings.Index(subStr, constants.OPEN_TOKEN_STR)
			closePos := strings.Index(subStr, constants.CLOSE_TOKEN_STR)
			var firstPos int
			if openPos == -1 && closePos == -1 {
				firstPos = len(word)
			} else if closePos == -1 {
				firstPos = openPos  + startingPos
			} else if openPos == -1 {
				firstPos = closePos + startingPos
			} else if openPos < closePos {
				firstPos = openPos + startingPos
			} else {
				firstPos = closePos + startingPos
			}

			token = word[startingPos:firstPos]
		}
		startingPos += len(token)
		tokens = append(tokens, token)
	}
	return tokens
}

var _ internal.WordTokenizer = &wordTokenizer{}
