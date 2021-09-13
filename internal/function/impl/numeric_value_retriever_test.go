package impl

import (
	"github.com/stretchr/testify/assert"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/printer/fakes"
	"testing"
)

func TestNodeIsNotNumeric(t *testing.T) {
	value := "value"
	node := datamodels.NewAtomNode(value)
	
	functionName := "functionName"
	index := 1

	listNotationPrinter := &fakes.FakeListNotationPrinter{}
	listNotationPrinter.PrintInListNotationReturns(value)
	
	numericValueRetriever := NewNumericValueRetriever(listNotationPrinter)
	
	actual, actualErr := numericValueRetriever.RetrieveNumericValue(node, functionName, index)
	
	assert.Equal(t, 0, actual)
	assert.NotNil(t, actualErr)
	args := listNotationPrinter.PrintInListNotationArgsForCall(0)
	assert.Equal(t, node, args)
}

func TestNodeIsNumeric(t *testing.T) {
	value := 34
	node := datamodels.NewAtomNodeFromInt(value)

	functionName := "functionName"
	index := 1

	listNotationPrinter := &fakes.FakeListNotationPrinter{}
	numericValueRetriever := NewNumericValueRetriever(listNotationPrinter)

	actual, actualErr := numericValueRetriever.RetrieveNumericValue(node, functionName, index)

	assert.Equal(t, value, actual)
	assert.Nil(t, actualErr)
	assert.Equal(t, 0, listNotationPrinter.PrintInListNotationCallCount())
}