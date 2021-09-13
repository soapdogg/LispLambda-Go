package impl

import (
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/function"
	"lisp_lambda-go/internal/function/impl/internal"
)

type plusPFunction struct {
	numericValueRetriever internal.NumericValueRetriever
}

func NewPlusPFunction(
	numericValueRetriever internal.NumericValueRetriever,
) *plusPFunction {
	return &plusPFunction{
		numericValueRetriever,
	}
}

func (p plusPFunction) Evaluate(params datamodels.NodeStack) (
	datamodels.Node,
	error,
) {
	numeric, err := p.numericValueRetriever.RetrieveNumericValue(
		params.Pop(),
		constants.PLUS_P,
		1,
	)
	if err != nil {
		return nil, err
	}
	isPositive := numeric > 0
	return datamodels.NewAtomNodeFromBool(isPositive), nil
}

var _ function.Function = &plusPFunction{}
