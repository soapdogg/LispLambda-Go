package function

import (
	"github.com/stretchr/testify/assert"
	"lisp_lambda-go/internal/core/datamodels"
	"testing"
)

func TestInputIsAList(t *testing.T) {
	node := datamodels.NewExpressionListNode([]datamodels.Node{})
	
	listValueRetriever := NewListValueRetriever()
	actual, actualErr := listValueRetriever.RetrieveListValue(node, "functionName")
	
	assert.Equal(t, node, actual)
	assert.Nil(t, actualErr)
}

func TestInputIsAtom(t *testing.T) {
	node := datamodels.NewAtomNode("string")

	listValueRetriever := NewListValueRetriever()
	actual, actualErr := listValueRetriever.RetrieveListValue(node, "functionName")

	assert.Nil(t, actual)
	assert.NotNil(t, actualErr)
}