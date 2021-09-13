package impl

import (
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/function"
	"lisp_lambda-go/internal/function/impl/internal"
)

type carFunction struct{
	listValueRetriever internal.ListValueRetriever
}

func NewCarFunction(
	listValueRetriever internal.ListValueRetriever,
) *carFunction {
	return &carFunction{
		listValueRetriever,
	}
}

func (c carFunction) Evaluate(params datamodels.NodeStack) (
	datamodels.Node,
	error,
) {
	first := params.Pop()
	firstExpressionListNode, err := c.listValueRetriever.RetrieveListValue(first, constants.CAR)
	if err != nil {
		return nil, err
	}

	return firstExpressionListNode.GetChildren()[0], nil
}

var _ function.Function = &carFunction{}