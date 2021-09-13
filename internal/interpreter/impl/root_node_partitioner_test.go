package impl

import (
	"github.com/stretchr/testify/assert"
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	"testing"
)

func TestAtomRootNode(t *testing.T){
	rootNodePartitioner := NewRootNodePartitioner()
	atomNode := datamodels.NewAtomNode("value")
	rootNodes := []datamodels.Node{atomNode}

	actual := rootNodePartitioner.PartitionRootNodes(rootNodes)
	assert.Equal(t, 0, len(actual.GetDefunNodes()))
	assert.Equal(t, 1, len(actual.GetEvaluatableNodes()))
	assert.Equal(t, atomNode, actual.GetEvaluatableNodes()[0])
}

func TestExpressionRootNodeWithExpressionAddress(t *testing.T) {
	rootNodePartitioner := NewRootNodePartitioner()
	child := datamodels.NewExpressionListNode([]datamodels.Node{})
	rootNode := datamodels.NewExpressionListNode([]datamodels.Node{child})
	rootNodes := []datamodels.Node{rootNode}

	actual := rootNodePartitioner.PartitionRootNodes(rootNodes)
	assert.Equal(t, 0, len(actual.GetDefunNodes()))
	assert.Equal(t, 1, len(actual.GetEvaluatableNodes()))
	assert.Equal(t, rootNode, actual.GetEvaluatableNodes()[0])
}

func TestExpressionRootNodeWithNonDefun(t *testing.T) {
	rootNodePartitioner := NewRootNodePartitioner()
	child := datamodels.NewAtomNode("new")
	rootNode := datamodels.NewExpressionListNode([]datamodels.Node{child})
	rootNodes := []datamodels.Node{rootNode}

	actual := rootNodePartitioner.PartitionRootNodes(rootNodes)
	assert.Equal(t, 0, len(actual.GetDefunNodes()))
	assert.Equal(t, 1, len(actual.GetEvaluatableNodes()))
	assert.Equal(t, rootNode, actual.GetEvaluatableNodes()[0])
}

func TestExpressionRootNodeWithDefun(t *testing.T) {
	rootNodePartitioner := NewRootNodePartitioner()
	child := datamodels.NewAtomNode(constants.DEFUN)
	rootNode := datamodels.NewExpressionListNode([]datamodels.Node{child})
	rootNodes := []datamodels.Node{rootNode}

	actual := rootNodePartitioner.PartitionRootNodes(rootNodes)
	assert.Equal(t, 0, len(actual.GetEvaluatableNodes()))
	assert.Equal(t, 1, len(actual.GetDefunNodes()))
	assert.Equal(t, rootNode, actual.GetDefunNodes()[0])
}
