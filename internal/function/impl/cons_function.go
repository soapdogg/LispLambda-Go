package impl

import (
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/function"
)

type consFunction struct {}

func NewConsFunction() *consFunction {
	return &consFunction{}
}

func (c consFunction) Evaluate(params datamodels.NodeStack) (
	datamodels.Node,
	error,
) {
	first := params.Pop()
	second := params.Pop()

	secondList, isList := second.(datamodels.ExpressionListNode)
	var children []datamodels.Node
	if isList {
		children = []datamodels.Node{first}
		children = append(children, secondList.GetChildren()...)
	} else {
		children = []datamodels.Node{first, second}
	}
	return datamodels.NewExpressionListNode(children), nil
}

var _ function.Function = &consFunction{}
