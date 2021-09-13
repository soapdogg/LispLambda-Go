package impl

import (
	tokenizer2 "lisp_lambda-go/internal/tokenizer"
	"lisp_lambda-go/internal/tokenizer/impl/internal"
	"strings"
)

type tokenizer struct {
	wordTokenizer internal.WordTokenizer
}

func NewTokenizer(
	wordTokenizer internal.WordTokenizer,
) *tokenizer {
	return &tokenizer{
		wordTokenizer,
	}
}

func (t *tokenizer) Tokenize(input string) []string {
	trimmedInput := strings.TrimSpace(input)
	words := strings.Fields(trimmedInput)

	var result []string
	for _, word := range words {
		tokens := t.wordTokenizer.TokenizeWord(word)
		result = append(result, tokens...)
	}
	return result
}

var _ tokenizer2.Tokenizer = &tokenizer{}


