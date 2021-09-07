package function

import (
	"errors"
	top "lisp_lambda-go/internal"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/function/internal"
	"strconv"
)

type numericValueRetriever struct {
	listNotationPrinter top.ListNotationPrinter
}

func NewNumericValueRetriever(
	listNotationPrinter top.ListNotationPrinter,
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