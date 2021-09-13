package impl

import (
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/function"
)

type atomFunction struct{}

func NewAtomFunction() *atomFunction {
	return &atomFunction{}
}

func (a *atomFunction) Evaluate(params datamodels.NodeStack) (datamodels.Node, error) {
	first := params.Pop()
	_, isAtomNode := first.(datamodels.AtomNode)
	return datamodels.NewAtomNodeFromBool(isAtomNode), nil
}

var _ function.Function = &atomFunction{}