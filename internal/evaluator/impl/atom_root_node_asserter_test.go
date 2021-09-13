package impl

import (
	"github.com/stretchr/testify/assert"
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	"testing"
)

func TestValueIsNumeric(t *testing.T) {
	atomNode := datamodels.NewAtomNodeFromInt(34)
	atomRootNodeAsserter := NewAtomRootNodeAsserter()
	actual := atomRootNodeAsserter.AssertAtomRootNode(atomNode)
	assert.Nil(t, actual)
}

func TestValueIsT(t *testing.T) {
	atomNode := datamodels.NewAtomNode(constants.T)
	atomRootNodeAsserter := NewAtomRootNodeAsserter()
	actual := atomRootNodeAsserter.AssertAtomRootNode(atomNode)
	assert.Nil(t, actual)
}

func TestValueIsNil(t *testing.T) {
	atomNode := datamodels.NewAtomNode(constants.NIL)
	atomRootNodeAsserter := NewAtomRootNodeAsserter()
	actual := atomRootNodeAsserter.AssertAtomRootNode(atomNode)
	assert.Nil(t, actual)
}

func TestValueIsInvalid(t *testing.T) {
	atomNode := datamodels.NewAtomNode("invalid")
	atomRootNodeAsserter := NewAtomRootNodeAsserter()
	actual := atomRootNodeAsserter.AssertAtomRootNode(atomNode)
	assert.NotNil(t, actual)
}
