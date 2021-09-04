package function

import (
	top "lisp_lambda-go/internal"
	"lisp_lambda-go/internal/core/datamodels"
)

type atomFunction struct{}

func NewAtomFunction() *atomFunction {
	return &atomFunction{}
}

func (a *atomFunction) Evaluate(params *datamodels.NodeStack) (datamodels.Node, error) {
	first := params.Pop()
	_, isAtomNode := first.(*datamodels.AtomNode)
	return datamodels.NewAtomNodeFromBool(isAtomNode), nil
}

var _ top.Function = &atomFunction{}