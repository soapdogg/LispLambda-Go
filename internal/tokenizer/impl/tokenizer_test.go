package impl

import (
	"github.com/stretchr/testify/assert"
	"lisp_lambda-go/internal/tokenizer/impl/internal/fakes"
	"testing"
)

func TestTokenizeLine(t *testing.T) {
	word1 := "word1"
	word2 := "word2"
	line := "\t" + word1 + "    " + word2 + "\n"

	wordTokenizer := &fakes.FakeWordTokenizer{}
	wordTokenizer.TokenizeWordReturnsOnCall(0, []string{word1})
	wordTokenizer.TokenizeWordReturnsOnCall(1, []string{word2})

	tokenizer := NewTokenizer(wordTokenizer)
	actual := tokenizer.Tokenize(line)

	assert.Equal(t, 2, len(actual))
	assert.Equal(t, word1, actual[0])
	assert.Equal(t, word2, actual[1])
}