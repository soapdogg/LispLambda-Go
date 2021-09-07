package interpreter

import (
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/interpreter/internal"
)

type rootNodePartitioner struct {}

func NewRootNodePartitioner() *rootNodePartitioner {
	return &rootNodePartitioner{}
}

func (r rootNodePartitioner) PartitionRootNodes(nodes []datamodels.Node) datamodels.PartitionedRootNodes {
	defun := []datamodels.ExpressionListNode{}
	executables := []datamodels.Node{}

	for _, rootNode := range nodes {
		expressionListNode, isExpressionListNode := rootNode.(datamodels.ExpressionListNode)
		if isExpressionListNode {
			firstChild := expressionListNode.GetChildren()[0]
			atomNode, isAtomNode := firstChild.(datamodels.AtomNode)
			if isAtomNode {
				if atomNode.GetValue() == constants.DEFUN {
					defun = append(defun, expressionListNode)
				} else {
					executables = append(executables, rootNode)
				}
			} else {
				executables = append(executables, rootNode)
			}
		} else {
			executables = append(executables, rootNode)
		}
	}

	return datamodels.NewPartitionedRootNodes(
		defun,
		executables,
	)
}

var _ internal.RootNodePartitioner = &rootNodePartitioner{}


