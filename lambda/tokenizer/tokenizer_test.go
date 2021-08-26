package tokenizer

import (
	"github.com/stretchr/testify/assert"
	"lisp_lambda-go/lambda/tokenizer/internal"
	"testing"
)

func TestTokenizer(t *testing.T) {
	input := "input"

	tokenizer := New(
		internal.New(),
	)

	actual := tokenizer.Tokenize(input)

	assert.Equal(t, input, actual[0])
}
