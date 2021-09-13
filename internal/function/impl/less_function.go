package impl

import (
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/function"
	"lisp_lambda-go/internal/function/impl/internal"
)

type lessFunction struct {
	numericValueRetriever internal.NumericValueRetriever
}

func NewLessFunction(numericValueRetriever internal.NumericValueRetriever) *lessFunction {
	return &lessFunction{numericValueRetriever}
}

func (l lessFunction) Evaluate(params datamodels.NodeStack) (
	datamodels.Node,
	error,
) {
	first := params.Pop()
	firstNumeric, err := l.numericValueRetriever.RetrieveNumericValue(
		first,
		constants.LESS,
		1,
	)
	if err != nil {
		return nil, err
	}

	result := firstNumeric
	current := 2
	for params.IsNotEmpty() {
		numeric, err := l.numericValueRetriever.RetrieveNumericValue(
			params.Pop(),
			constants.LESS,
			current,
		)
		if err != nil {
			return nil, err
		}
		if result >= numeric {
			return datamodels.NewAtomNode(constants.NIL), nil
		}
		result = numeric
		current++
	}
	return datamodels.NewAtomNode(constants.T), nil
}

var _ function.Function = &lessFunction{}
