package function

import (
	"github.com/stretchr/testify/assert"
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	"testing"
)

func TestFirstIsAtom(t *testing.T) {
	node := datamodels.NewAtomNode(constants.NIL)
	params := datamodels.NewNodeStack()
	params.Push(node)

	atomFunction := NewAtomFunction()

	actual, actualErr := atomFunction.Evaluate(params)

	atomNode := actual.(datamodels.AtomNode)
	assert.Equal(t, constants.T, atomNode.GetValue())
	assert.Nil(t, actualErr)
}

func TestFirstIsExpressionList(t *testing.T) {
	node := datamodels.NewExpressionListNode([]datamodels.Node{})
	params := datamodels.NewNodeStack()
	params.Push(node)

	atomFunction := NewAtomFunction()

	actual, actualErr := atomFunction.Evaluate(params)

	atomNode := actual.(datamodels.AtomNode)
	assert.Equal(t, constants.NIL, atomNode.GetValue())
	assert.Nil(t, actualErr)
}

