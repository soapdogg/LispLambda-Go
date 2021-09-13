package impl

import (
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/function"
	"lisp_lambda-go/internal/printer"
)

type eqFunction struct {
	listNotationPrinter printer.ListNotationPrinter
}

func NewEqFunction(
	listNotationPrinter printer.ListNotationPrinter,
) *eqFunction {
	return &eqFunction{
		listNotationPrinter,
	}
}

func (e eqFunction) Evaluate(params datamodels.NodeStack) (
	datamodels.Node,
	error,
) {
 	first := params.Pop()
 	value := e.listNotationPrinter.PrintInListNotation(first)

 	for params.IsNotEmpty() {
 		node := params.Pop()
 		nodeValue := e.listNotationPrinter.PrintInListNotation(node)
 		if nodeValue != value {
 			return datamodels.NewAtomNodeFromBool(false), nil
		}
	}
	return datamodels.NewAtomNodeFromBool(true), nil
}

var _ function.Function = &eqFunction{}
