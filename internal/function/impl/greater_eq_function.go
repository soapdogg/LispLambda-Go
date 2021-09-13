package impl

import (
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/function"
	"lisp_lambda-go/internal/function/impl/internal"
)

type greaterEqFunction struct {
	numericValueRetriever internal.NumericValueRetriever
}

func NewGreaterEqFunction(
	numericValueRetriever internal.NumericValueRetriever,
) *greaterEqFunction {
	return &greaterEqFunction{numericValueRetriever}
}

func (g greaterEqFunction) Evaluate(params datamodels.NodeStack) (
	datamodels.Node,
	error,
) {
	first := params.Pop()
	firstNumeric, err := g.numericValueRetriever.RetrieveNumericValue(
		first,
		constants.GREATER_EQ,
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
			constants.GREATER_EQ,
			current,
		)
		if err != nil {
			return nil, err
		}
		if result < numeric {
			return datamodels.NewAtomNode(constants.NIL), nil
		}
		result = numeric
		current++
	}
	return datamodels.NewAtomNode(constants.T), nil
}

var _ function.Function = &greaterEqFunction{}
