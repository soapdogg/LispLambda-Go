package function

import (
	"github.com/stretchr/testify/assert"
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	"testing"
)

func TestAtomIsNil(t *testing.T) {
	nullFunction := NewNullFunction()

	node := datamodels.NewAtomNode(constants.NIL)
	params := datamodels.NewNodeStack()
	params.Push(node)

	actual := nullFunction.Evaluate(params)

	atomNode := actual.(*datamodels.AtomNode)
	assert.Equal(t, constants.T, atomNode.GetValue())
}

func TestAtomIsExpressionList(t *testing.T) {
	nullFunction := NewNullFunction()

	node := datamodels.NewExpressionListNode([]datamodels.Node{})
	params := datamodels.NewNodeStack()
	params.Push(node)

	actual := nullFunction.Evaluate(params)

	atomNode := actual.(*datamodels.AtomNode)
	assert.Equal(t, constants.NIL, atomNode.GetValue())
}
