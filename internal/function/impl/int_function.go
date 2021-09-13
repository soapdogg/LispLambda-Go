package impl

import (
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/function"
	"strconv"
)

type intFunction struct {}

func NewIntFunction() *intFunction {
	return &intFunction{}
}

func (i intFunction) Evaluate(params datamodels.NodeStack) (
	datamodels.Node,
	error,
) {
	first := params.Pop()
	atomFirst, isAtom := first.(datamodels.AtomNode)
	if isAtom {
		value := atomFirst.GetValue()
		_, convErr := strconv.Atoi(value)
		isNumeric := convErr == nil
		return datamodels.NewAtomNodeFromBool(isNumeric), nil
	}
	return datamodels.NewAtomNode(constants.NIL), nil
}

var _ function.Function = &intFunction{}