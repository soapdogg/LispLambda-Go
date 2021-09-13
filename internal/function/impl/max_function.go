package impl

import (
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/function"
	"lisp_lambda-go/internal/function/impl/internal"
)

type maxFunction struct {
	numericValueRetriever internal.NumericValueRetriever
}

func NewMaxFunction(numericValueRetriever internal.NumericValueRetriever) *maxFunction {
	return &maxFunction{numericValueRetriever}
}

func (l maxFunction) Evaluate(params datamodels.NodeStack) (
	datamodels.Node,
	error,
) {
	first := params.Pop()
	result, err := l.numericValueRetriever.RetrieveNumericValue(
		first,
		constants.MAX,
		1,
	)
	if err != nil {
		return nil, err
	}

	current := 2
	for params.IsNotEmpty() {
		numeric, err := l.numericValueRetriever.RetrieveNumericValue(
			params.Pop(),
			constants.MAX,
			current,
		)
		if err != nil {
			return nil, err
		}
		if result < numeric {
		 	result = numeric
		}
		current++
	}
	return datamodels.NewAtomNodeFromInt(result), nil
}

var _ function.Function = &maxFunction{}
