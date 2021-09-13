package impl

import (
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/function"
)

type notEqFunction struct {}

func NewNotEqFunction() *notEqFunction{
	return &notEqFunction{}
}

func (n notEqFunction) Evaluate(params datamodels.NodeStack) (
	datamodels.Node,
	error,
) {
	first := params.Pop()
	firstValue := first.(datamodels.AtomNode).GetValue()
	values := map[string]bool{firstValue: true}
	for params.IsNotEmpty() {
		node := params.Pop()
		value := node.(datamodels.AtomNode).GetValue()
		isInMap := values[value]
		if isInMap {
			return datamodels.NewAtomNode(constants.NIL), nil
		}
		values[value] = true
	}
	return datamodels.NewAtomNode(constants.T), nil
}

var _ function.Function = &notEqFunction{}
