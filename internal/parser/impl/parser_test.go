package impl

import (
	"github.com/stretchr/testify/assert"
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/parser/impl/internal/fakes"
	"testing"
)

func TestParser(t *testing.T) {
	headToken := "token"
	tokens := []string{headToken}

	resultingNode := datamodels.NewAtomNode(headToken)
	nodeParser := &fakes.FakeNodeParser{}
	nodeParser.ParseIntoNodeReturns(resultingNode)
	
	parser := NewParser(nodeParser)
	actual, actualErr := parser.Parse(tokens)
	
	assert.Equal(t, resultingNode, actual[0])
	assert.Nil(t, actualErr)
}

func TestTooManyOpen(t *testing.T) {
	headToken := constants.OPEN_TOKEN_STR
	tokens := []string{headToken}

	nodeParser := &fakes.FakeNodeParser{}

	parser := NewParser(nodeParser)
	actual, actualErr := parser.Parse(tokens)

	assert.Nil(t, actual)
	assert.NotNil(t, actualErr)
}

func TestTooManyClose(t *testing.T) {
	headToken := constants.CLOSE_TOKEN_STR
	tokens := []string{headToken}

	nodeParser := &fakes.FakeNodeParser{}

	parser := NewParser(nodeParser)
	actual, actualErr := parser.Parse(tokens)

	assert.Nil(t, actual)
	assert.NotNil(t, actualErr)
}