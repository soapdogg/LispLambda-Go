package impl

import (
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/function"
	"lisp_lambda-go/internal/function/impl/internal"
)

type greaterFunction struct {
	numericValueRetriever internal.NumericValueRetriever
}

func NewGreaterFunction(
	numericValueRetriever internal.NumericValueRetriever,
) *greaterFunction {
	return &greaterFunction{numericValueRetriever}
}

func (g greaterFunction) Evaluate(params datamodels.NodeStack) (
	datamodels.Node,
	error,
) {
	first := params.Pop()
	firstNumeric, err := g.numericValueRetriever.RetrieveNumericValue(
		first,
		constants.GREATER,
		1,
	)
	if err != nil {
		return nil, err
	}

	result := firstNumeric
	current := 2
	for params.IsNotEmpty() {
		numeric, err := g.numericValueRetriever.RetrieveNumericValue(
			params.Pop(),
			constants.GREATER,
			current,
		)
		if err != nil {
			return nil, err
		}
		if result <= numeric {
			return datamodels.NewAtomNode(constants.NIL), nil
		}
		result = numeric
		current++
	}
	return datamodels.NewAtomNode(constants.T), nil
}

var _ function.Function = &greaterFunction{}
