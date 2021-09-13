package impl

import (
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/function"
	"lisp_lambda-go/internal/function/impl/internal"
)

type lessEqFunction struct {
	numericValueRetriever internal.NumericValueRetriever
}

func NewLessEqFunction(numericValueRetriever internal.NumericValueRetriever) *lessEqFunction {
	return &lessEqFunction{numericValueRetriever}
}

func (l lessEqFunction) Evaluate(params datamodels.NodeStack) (
	datamodels.Node,
	error,
) {
	first := params.Pop()
	firstNumeric, err := l.numericValueRetriever.RetrieveNumericValue(
		first,
		constants.LESS_EQ,
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
			constants.LESS_EQ,
			current,
		)
		if err != nil {
			return nil, err
		}
		if result > numeric {
			return datamodels.NewAtomNode(constants.NIL), nil
		}
		result = numeric
		current++
	}
	return datamodels.NewAtomNode(constants.T), nil
}

var _ function.Function = &lessEqFunction{}
