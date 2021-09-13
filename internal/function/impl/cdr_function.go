package impl

import (
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/function"
	"lisp_lambda-go/internal/function/impl/internal"
)

type cdrFunction struct {
	listValueRetriever internal.ListValueRetriever
}

func NewCdrFunction(
	listValueRetriever internal.ListValueRetriever,
) *cdrFunction {
	return &cdrFunction{
		listValueRetriever,
	}
}

func (c cdrFunction) Evaluate(params datamodels.NodeStack) (
	datamodels.Node,
	error,
) {
	first := params.Pop()
	firstExpressionListNode, err := c.listValueRetriever.RetrieveListValue(
		first,
		constants.CDR,
	)
	if err != nil {
		return nil, err
	}
	childrenLength := len(firstExpressionListNode.GetChildren())
	if childrenLength == 1 {
		return firstExpressionListNode.GetChildren()[0], nil
	} else if childrenLength == 2 {
		return firstExpressionListNode.GetChildren()[1], nil
	} else {
		return datamodels.NewExpressionListNode(
			firstExpressionListNode.GetChildren()[1:len(firstExpressionListNode.GetChildren())],
		), nil
	}
}

var _ function.Function = &cdrFunction{}