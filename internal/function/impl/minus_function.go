package impl

import (
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/function"
	"lisp_lambda-go/internal/function/impl/internal"
)

type minusFunction struct {
	numericValueRetriever internal.NumericValueRetriever
}


func NewMinusFunction(
	numericValueRetriever internal.NumericValueRetriever,
) *minusFunction {
	return &minusFunction{numericValueRetriever}
}

func (m minusFunction) Evaluate(params datamodels.NodeStack) (
	datamodels.Node,
	error,
) {
 	first := params.Pop()
 	firstNumeric, err := m.numericValueRetriever.RetrieveNumericValue(
 		first,
 		constants.MINUS,
 		1,
	)
 	if err != nil {
		return nil, err
	}
	if params.IsEmpty() {
		return datamodels.NewAtomNodeFromInt(-firstNumeric), nil
	}

	result := firstNumeric
	current := 2
	for params.IsNotEmpty() {
		numeric, err := m.numericValueRetriever.RetrieveNumericValue(
			params.Pop(),
			constants.MINUS,
			current,
		)
		if err != nil {
			return nil, err
		}
		result -= numeric
		current++
	}
	return datamodels.NewAtomNodeFromInt(result), nil
}

var _ function.Function = &minusFunction{}