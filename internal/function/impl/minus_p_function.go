package impl

import (
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/function"
	"lisp_lambda-go/internal/function/impl/internal"
)

type minusPFunction struct {
	numericValueRetriever internal.NumericValueRetriever
}

func NewMinusPFunction(
	numericValueRetriever internal.NumericValueRetriever,
) *minusPFunction {
	return &minusPFunction{
		numericValueRetriever,
	}
}

func (m minusPFunction) Evaluate(params datamodels.NodeStack) (
	datamodels.Node,
	error,
) {
	numeric, err := m.numericValueRetriever.RetrieveNumericValue(
		params.Pop(),
		constants.MINUS_P,
		1,
	)
	if err != nil {
		return nil, err
	}
	isNegative := numeric < 0
	return datamodels.NewAtomNodeFromBool(isNegative), nil
}

var _ function.Function = &minusPFunction{}
