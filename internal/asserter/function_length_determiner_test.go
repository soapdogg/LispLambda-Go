package asserter

import (
	"github.com/stretchr/testify/assert"
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	"testing"
)

func TestLengthOfNilNode(t *testing.T) {
	functionLengthDeterminer := NewFunctionLengthDeterminer()
	nilNode := datamodels.NewAtomNode(constants.NIL)
	actual := functionLengthDeterminer.DetermineFunctionLength(nilNode)
	assert.Equal(t, 0, actual)
}

func TestLengthOfNonNilNode(t *testing.T) {
	functionLengthDeterminer := NewFunctionLengthDeterminer()
	notNilNode := datamodels.NewAtomNode("notnil")
	actual := functionLengthDeterminer.DetermineFunctionLength(notNilNode)
	assert.Equal(t, 1, actual)
}

func TestLengthOfExpressionListNode(t *testing.T) {
	functionLengthDeterminer := NewFunctionLengthDeterminer()
	notNilNode := datamodels.NewAtomNode("notnil")
	children := []datamodels.Node{notNilNode}
	expressionList := datamodels.NewExpressionListNode(children)
	actual := functionLengthDeterminer.DetermineFunctionLength(expressionList)
	assert.Equal(t, 1, actual)
}

func TestLengthOfExpressionListNodeLastElementNil(t *testing.T) {
	functionLengthDeterminer := NewFunctionLengthDeterminer()
	notNilNode := datamodels.NewAtomNode(constants.NIL)
	children := []datamodels.Node{notNilNode}
	expressionList := datamodels.NewExpressionListNode(children)
	actual := functionLengthDeterminer.DetermineFunctionLength(expressionList)
	assert.Equal(t, 0, actual)
}