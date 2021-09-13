package impl

import (
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/function"
	"lisp_lambda-go/internal/function/impl/internal"
)

type timesFunction struct {
	numericValueRetriever internal.NumericValueRetriever
}

func NewTimesFunction(
	numericValueRetriever internal.NumericValueRetriever,
) *timesFunction {
	return &timesFunction{
		numericValueRetriever,
	}
}

func (t timesFunction) Evaluate(params datamodels.NodeStack) (
	datamodels.Node,
	error,
) {
	result := 1
	current := 1
	for params.IsNotEmpty() {
		numeric, err := t.numericValueRetriever.RetrieveNumericValue(
			params.Pop(),
			constants.TIMES,
			current,
		)
		if err != nil {
			return nil, err
		}
		result *= numeric
		current++
	}
	return datamodels.NewAtomNodeFromInt(result), nil
}

var _ function.Function = &timesFunction{}
