package impl

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/core/datamodels/fakes"
	internalFakes "lisp_lambda-go/internal/function/impl/internal/fakes"
	"testing"
)

func TestListValueRetrieverThrowsError(t *testing.T) {
	first := &fakes.FakeNode{}
	params := datamodels.NewNodeStack()
	params.Push(first)

	err := errors.New("error")
	listValueRetriever := &internalFakes.FakeListValueRetriever{}
	listValueRetriever.RetrieveListValueReturns(nil, err)

	carFunction := NewCarFunction(listValueRetriever)

	actual, actualErr := carFunction.Evaluate(params)
	assert.Nil(t, actual)
	assert.Equal(t, err, actualErr)
}

func TestNoErrors(t *testing.T) {
	first := &fakes.FakeNode{}
	params := datamodels.NewNodeStack()
	params.Push(first)


	child := &fakes.FakeNode{}
	firstList := datamodels.NewExpressionListNode([]datamodels.Node{child})
	listValueRetriever := &internalFakes.FakeListValueRetriever{}
	listValueRetriever.RetrieveListValueReturns(firstList, nil)

	carFunction := NewCarFunction(listValueRetriever)

	actual, actualErr := carFunction.Evaluate(params)
	assert.Equal(t, child, actual)
	assert.Nil(t, actualErr)
}