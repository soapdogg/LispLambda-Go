package impl

import tokenizer2 "lisp_lambda-go/internal/tokenizer"

type singleton struct {
	tokenizer tokenizer2.Tokenizer
}

func NewSingleton() *singleton {
	wordTokenizer := NewWordTokenizer()
	tokenizer := NewTokenizer(
		wordTokenizer,
	)
	return &singleton{
		tokenizer,
	}
}

func (s *singleton) GetTokenizer() tokenizer2.Tokenizer {
	return s.tokenizer
}