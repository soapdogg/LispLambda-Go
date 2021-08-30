package asserter

import (
	"lisp_lambda-go/lambda/asserter/internal"
	"lisp_lambda-go/lambda/core/constants"
	"lisp_lambda-go/lambda/core/datamodels"
)

type functionLengthDeterminer struct {}

func NewFunctionLengthDeterminer()*functionLengthDeterminer {
	return &functionLengthDeterminer{}
}

func (f *functionLengthDeterminer) DetermineFunctionLength(node datamodels.Node) int {
	expressionListNode, isExpressionList := node.(*datamodels.ExpressionListNode)
	if !isExpressionList {
		atomNode, _ := node.(*datamodels.AtomNode)
		if atomNode.GetValue() == constants.NIL {
			return 0
		} else {
			return 1
		}
	} else {
		children := expressionListNode.GetChildren()
		lastChild := children[len(children) - 1]
		lastChildAtomNode, ok := lastChild.(*datamodels.AtomNode)
		if ok && lastChildAtomNode.GetValue() == constants.NIL {
			return len(children) - 1
		} else {
			return len(children)
		}
	}
}

var _ internal.FunctionLengthDeterminer = &functionLengthDeterminer{}