package impl

import (
	"errors"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/function/impl/internal"
	"lisp_lambda-go/internal/printer"
	"strconv"
)

type numericValueRetriever struct {
	listNotationPrinter printer.ListNotationPrinter
}

func NewNumericValueRetriever(
	listNotationPrinter printer.ListNotationPrinter,
) *numericValueRetriever {
	return &numericValueRetriever{
		listNotationPrinter,
	}
}

func (n numericValueRetriever) RetrieveNumericValue(node datamodels.Node, functionName string, index int) (int, error) {
	atomNode, isAtomNode := node.(datamodels.AtomNode)
	if isAtomNode {
		value := atomNode.GetValue()
		valueNumeric, err := strconv.Atoi(value)
		if err == nil {
			return valueNumeric, nil
		}
	}
	value := n.listNotationPrinter.PrintInListNotation(node)
	return 0, errors.New("Error! Parameter at position: " + strconv.Itoa(index) + " of function " + functionName + " is not numeric!    Actual: " + value)
}

var _ internal.NumericValueRetriever = &numericValueRetriever{}