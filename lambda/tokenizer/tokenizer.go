package tokenizer

import "lisp_lambda-go/lambda/tokenizer/internal"

type Tokenizer interface {
	Tokenize(input string) []string
}

type tokenizer struct {
	wordTokenizer internal.WordTokenizer
}

func New(
	wordTokenizer internal.WordTokenizer,
) *tokenizer {
	return &tokenizer{
		wordTokenizer,
	}
}

func (t tokenizer) Tokenize(input string) []string {
	return []string{input}

	//trimmedInput := input

	//return t.wordTokenizer.TokenizeWord(trimmedInput)
}

var _ Tokenizer = tokenizer{}


