package tokenizer

import (
	"lisp_lambda-go/lambda/tokenizer/internal"
	"strings"
)

type Tokenizer interface {
	Tokenize(input string) []string
}

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

var _ Tokenizer = &tokenizer{}


