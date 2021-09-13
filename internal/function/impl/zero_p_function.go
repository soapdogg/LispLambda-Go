package impl

import (
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/function"
	"lisp_lambda-go/internal/function/impl/internal"
)

type zeroPFunction struct {
	numericValueRetriever internal.NumericValueRetriever
}

func NewZeroPFunction(
	numericValueRetriever internal.NumericValueRetriever,
) *zeroPFunction {
	return &zeroPFunction{
		numericValueRetriever,
	}
}

func (z zeroPFunction) Evaluate(params datamodels.NodeStack) (
	datamodels.Node,
	error,
) {
	numeric, err := z.numericValueRetriever.RetrieveNumericValue(
		params.Pop(),
		constants.ZERO_P,
		1,
	)
	if err != nil {
		return nil, err
	}
	isZero := numeric == 0
	return datamodels.NewAtomNodeFromBool(isZero), nil
}

var _ function.Function = &zeroPFunction{}
