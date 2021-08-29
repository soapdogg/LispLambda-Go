package internal

import (
	"github.com/stretchr/testify/assert"
	"lisp_lambda-go/lambda/core/constants"
	"testing"
)


func TestCloseParentheses(t *testing.T) {
	wordTokenizer := NewWordTokenizer()

	word := constants.CLOSE_TOKEN_STR
	actual := wordTokenizer.TokenizeWord(word)

	assert.Equal(t, 1, len(actual))
	assert.Equal(t, word, actual[0])
}

func TestOpenParentheses(t *testing.T) {
	wordTokenizer := NewWordTokenizer()

	word := constants.OPEN_TOKEN_STR
	actual := wordTokenizer.TokenizeWord(word)

	assert.Equal(t, 1, len(actual))
	assert.Equal(t, word, actual[0])
}

func TestLiteralParentheses(t *testing.T) {
	wordTokenizer := NewWordTokenizer()

	word := "ABC"
	actual := wordTokenizer.TokenizeWord(word)

	assert.Equal(t, 1, len(actual))
	assert.Equal(t, word, actual[0])
}

func TestMixParentheses(t *testing.T) {
	wordTokenizer := NewWordTokenizer()

	literal := "ABC"
	word := constants.OPEN_TOKEN_STR + literal + constants.CLOSE_TOKEN_STR
	actual := wordTokenizer.TokenizeWord(word)

	assert.Equal(t, 3, len(actual))
	assert.Equal(t, constants.OPEN_TOKEN_STR, actual[0])
	assert.Equal(t, literal, actual[1])
	assert.Equal(t, constants.CLOSE_TOKEN_STR, actual[2])
}