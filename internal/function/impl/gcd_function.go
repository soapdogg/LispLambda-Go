package impl

import (
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/function"
	"lisp_lambda-go/internal/function/impl/internal"
)

type gcdFunction struct {
	numericValueRetriever internal.NumericValueRetriever
	gcdCalculator internal.GCDCalculator
}

func NewGCDFunction(
	numericValueRetriever internal.NumericValueRetriever,
	gcdCalculator internal.GCDCalculator,
) *gcdFunction {
	return &gcdFunction{
		numericValueRetriever,
		gcdCalculator,
	}
}

func (g gcdFunction) Evaluate(params datamodels.NodeStack) (
	datamodels.Node,
	error,
) {
	if params.IsEmpty() {
		return datamodels.NewAtomNodeFromInt(0), nil
	}
	first := params.Pop()
	firstNumeric, err := g.numericValueRetriever.RetrieveNumericValue(
		first,
		constants.GCD,
		1,
	)
	if err != nil {
		return nil, err
	}
	result := firstNumeric
	if result < 0 {
		result = -result
	}
	current := 2
	for params.IsNotEmpty() {
		numeric, err := g.numericValueRetriever.RetrieveNumericValue(
			params.Pop(),
			constants.GCD,
			current,
		)
		if err != nil {
			return nil, err
		}
		if numeric < 0 {
			numeric = -numeric
		}

		result = g.gcdCalculator.CalculateGCD(result, numeric)
		current++
	}
	return datamodels.NewAtomNodeFromInt(result), nil
}

var _ function.Function = &gcdFunction{}
