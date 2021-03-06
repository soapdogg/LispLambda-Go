package impl

import (
	"errors"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/function/impl/internal"
)

type listValueRetriever struct {}

func NewListValueRetriever() *listValueRetriever {
	return &listValueRetriever{}
}

func (l listValueRetriever) RetrieveListValue(node datamodels.Node, functionName string) (datamodels.ExpressionListNode, error) {
	expressionList, isExpressionList := node.(datamodels.ExpressionListNode)
	if isExpressionList {
		return expressionList, nil
	}
	atomNode := node.(datamodels.AtomNode)
	return nil, errors.New("Error! Parameter of " + functionName + " is not a list.    Actual: " + atomNode.GetValue())
}

var _ internal.ListValueRetriever = &listValueRetriever{}
