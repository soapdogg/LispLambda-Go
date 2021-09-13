package impl

import (
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/function"
	"lisp_lambda-go/internal/function/impl/internal"
)

type lcmFunction struct {
	numericValueRetriever internal.NumericValueRetriever
	gcdCalculator internal.GCDCalculator
}

func NewLCMFunction(
	numericValueRetriever internal.NumericValueRetriever,
	gcdCalculator internal.GCDCalculator,
) *lcmFunction {
	return &lcmFunction{
		numericValueRetriever,
		gcdCalculator,
	}
}

func (g lcmFunction) Evaluate(params datamodels.NodeStack) (
	datamodels.Node,
	error,
) {
	first := params.Pop()
	firstNumeric, err := g.numericValueRetriever.RetrieveNumericValue(
		first,
		constants.LCM,
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
			constants.LCM,
			current,
		)
		if err != nil {
			return nil, err
		}
		if numeric < 0 {
			numeric = -numeric
		}

		gcd := g.gcdCalculator.CalculateGCD(result, numeric)
		result = (result * numeric) / gcd
		current++
	}
	return datamodels.NewAtomNodeFromInt(result), nil
}

var _ function.Function = &lcmFunction{}
