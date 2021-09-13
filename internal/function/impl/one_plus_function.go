package impl

import (
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/function"
	"lisp_lambda-go/internal/function/impl/internal"
)

type onePlusFunction struct {
	numericValueRetriever internal.NumericValueRetriever
}

func NewOnePlusFunction(
	numericValueRetriever internal.NumericValueRetriever,
) *onePlusFunction {
	return &onePlusFunction{
		numericValueRetriever,
	}
}

func (o onePlusFunction) Evaluate(params datamodels.NodeStack) (
	datamodels.Node,
	error,
) {
	numeric, err := o.numericValueRetriever.RetrieveNumericValue(
		params.Pop(),
		constants.ONE_PLUS,
		1,
	)
	if err != nil {
		return nil, err
	}
	return datamodels.NewAtomNodeFromInt(numeric + 1), nil
}


var _ function.Function = &onePlusFunction{}
