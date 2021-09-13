package impl

import (
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/function"
	"lisp_lambda-go/internal/function/impl/internal"
)

type plusFunction struct {
	numericValueRetriever internal.NumericValueRetriever
}

func NewPlusFunction(
	numericValueRetriever internal.NumericValueRetriever,
) *plusFunction {
	return &plusFunction{
		numericValueRetriever,
	}
}

func (p plusFunction) Evaluate(params datamodels.NodeStack) (
	datamodels.Node,
	error,
) {
	result := 0
	current := 1
	for params.IsNotEmpty() {
		numeric, err := p.numericValueRetriever.RetrieveNumericValue(
			params.Pop(),
			constants.PLUS,
			current,
		)
		if err != nil {
			return nil, err
		}
		result += numeric
		current++
	}
	return datamodels.NewAtomNodeFromInt(result), nil
}

var _ function.Function = &plusFunction{}
