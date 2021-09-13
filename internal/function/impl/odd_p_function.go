package impl

import (
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/function"
	"lisp_lambda-go/internal/function/impl/internal"
)

type oddPFunction struct {
	numericValueRetriever internal.NumericValueRetriever
}

func NewOddPFunction(
	numericValueRetriever internal.NumericValueRetriever,
) *oddPFunction {
	return &oddPFunction{
		numericValueRetriever,
	}
}

func (p oddPFunction) Evaluate(params datamodels.NodeStack) (
	datamodels.Node,
	error,
) {
	numeric, err := p.numericValueRetriever.RetrieveNumericValue(
		params.Pop(),
		constants.ODD_P,
		1,
	)
	if err != nil {
		return nil, err
	}
	isOdd := numeric % 2 != 0
	return datamodels.NewAtomNodeFromBool(isOdd), nil
}

var _ function.Function = &oddPFunction{}
