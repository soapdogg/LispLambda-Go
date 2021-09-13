package impl

import (
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/function"
	"lisp_lambda-go/internal/function/impl/internal"
)

type evenPFunction struct {
	numericValueRetriever internal.NumericValueRetriever
}

func NewEvenPFunction(
	numericValueRetriever internal.NumericValueRetriever,
) *evenPFunction {
	return &evenPFunction{
		numericValueRetriever,
	}
}

func (p evenPFunction) Evaluate(params datamodels.NodeStack) (
	datamodels.Node,
	error,
) {
	numeric, err := p.numericValueRetriever.RetrieveNumericValue(
		params.Pop(),
		constants.EVEN_P,
		1,
	)
	if err != nil {
		return nil, err
	}
	isEven := numeric % 2 == 0
	return datamodels.NewAtomNodeFromBool(isEven), nil
}

var _ function.Function = &evenPFunction{}
