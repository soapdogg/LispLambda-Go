package impl

import (
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/function"
	"lisp_lambda-go/internal/function/impl/internal"
)

type oneMinusFunction struct {
	numericValueRetriever internal.NumericValueRetriever
}

func NewOneMinusFunction(
	numericValueRetriever internal.NumericValueRetriever,
) *oneMinusFunction {
	return &oneMinusFunction{
		numericValueRetriever,
	}
}

func (o oneMinusFunction) Evaluate(params datamodels.NodeStack) (
	datamodels.Node,
	error,
) {
	numeric, err := o.numericValueRetriever.RetrieveNumericValue(
		params.Pop(),
		constants.ONE_MINUS,
		1,
	)
	if err != nil {
		return nil, err
	}
	return datamodels.NewAtomNodeFromInt(numeric - 1), nil
}


var _ function.Function = &oneMinusFunction{}
