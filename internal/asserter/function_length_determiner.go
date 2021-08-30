package asserter

import (
	"lisp_lambda-go/internal/asserter/internal"
	"lisp_lambda-go/internal/core/constants"
	datamodels2 "lisp_lambda-go/internal/core/datamodels"
)

type functionLengthDeterminer struct {}

func NewFunctionLengthDeterminer()*functionLengthDeterminer {
	return &functionLengthDeterminer{}
}

func (f *functionLengthDeterminer) DetermineFunctionLength(node datamodels2.Node) int {
	expressionListNode, isExpressionList := node.(*datamodels2.ExpressionListNode)
	if !isExpressionList {
		atomNode, _ := node.(*datamodels2.AtomNode)
		if atomNode.GetValue() == constants.NIL {
			return 0
		} else {
			return 1
		}
	} else {
		children := expressionListNode.GetChildren()
		lastChild := children[len(children) - 1]
		lastChildAtomNode, ok := lastChild.(*datamodels2.AtomNode)
		if ok && lastChildAtomNode.GetValue() == constants.NIL {
			return len(children) - 1
		} else {
			return len(children)
		}
	}
}

var _ internal.FunctionLengthDeterminer = &functionLengthDeterminer{}