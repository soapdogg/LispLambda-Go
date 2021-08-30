package function

import (
	top "lisp_lambda-go/internal"
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
)

type nullFunction struct{}

func NewNullFunction() *nullFunction {
	return &nullFunction{}
}

func (n nullFunction) Evaluate(params *datamodels.NodeStack) datamodels.Node {
	first := params.Pop()
	atomNode, isAtomNode := first.(*datamodels.AtomNode)
	if isAtomNode {
		value := atomNode.GetValue()
		isNil := value == constants.NIL
		return datamodels.NewAtomNodeFromBool(isNil)
	}
	return datamodels.NewAtomNode(constants.NIL)
}

var _ top.Function = &nullFunction{}