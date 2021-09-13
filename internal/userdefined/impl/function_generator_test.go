package impl

import (
	"errors"
	"github.com/stretchr/testify/assert"
	asserterFakes "lisp_lambda-go/internal/asserter/fakes"
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	fakes2 "lisp_lambda-go/internal/userdefined/impl/internal/fakes"
	"testing"
)

func TestFunctionLengthAsserterThrowsError(t *testing.T) {
	err := errors.New("functionLengthAsserter error")
	functionLengthAsserter := &asserterFakes.FakeFunctionLengthAsserter{}
	functionLengthAsserter.AssertLengthIsAsExpectedReturns(err)

	functionNameAsserter := &fakes2.FakeFunctionNameAsserter{}
	formalParametersAsserter := &fakes2.FakeFormalParametersAsserter{}

	functionGenerator := NewFunctionGenerator(
		functionLengthAsserter,
		functionNameAsserter,
		formalParametersAsserter,
	)

	params := datamodels.NewExpressionListNode([]datamodels.Node{})

	actual, actualErr := functionGenerator.GenerateFunction(params)

	assert.Nil(t, actual)
	assert.Equal(t, err, actualErr)

	actualFunctionName, actualLength, actualParams := functionLengthAsserter.AssertLengthIsAsExpectedArgsForCall(0)
	assert.Equal(t, constants.DEFUN, actualFunctionName)
	assert.Equal(t, constants.FOUR, actualLength)
	assert.Equal(t, params, actualParams)

	assert.Equal(t, 0, functionNameAsserter.AssertFunctionNameIsValidCallCount())
	assert.Equal(t, 0, formalParametersAsserter.AssertFormalParametersCallCount())
}

func TestFunctionNameAsserterThrowsError(t *testing.T) {
	functionLengthAsserter := &asserterFakes.FakeFunctionLengthAsserter{}
	functionLengthAsserter.AssertLengthIsAsExpectedReturns(nil)

	err := errors.New("function name asserter error")
	functionNameAsserter := &fakes2.FakeFunctionNameAsserter{}
	functionNameAsserter.AssertFunctionNameIsValidReturns(err)

	formalParametersAsserter := &fakes2.FakeFormalParametersAsserter{}

	functionGenerator := NewFunctionGenerator(
		functionLengthAsserter,
		functionNameAsserter,
		formalParametersAsserter,
	)

	defunNode := datamodels.NewAtomNode(constants.DEFUN)
	functionName := "functionName"
	functionNameNode := datamodels.NewAtomNode(functionName)
	params := datamodels.NewExpressionListNode([]datamodels.Node{defunNode, functionNameNode})

	actual, actualErr := functionGenerator.GenerateFunction(params)

	assert.Nil(t, actual)
	assert.Equal(t, err, actualErr)

	actualFunctionName, actualLength, actualParams := functionLengthAsserter.AssertLengthIsAsExpectedArgsForCall(0)
	assert.Equal(t, constants.DEFUN, actualFunctionName)
	assert.Equal(t, constants.FOUR, actualLength)
	assert.Equal(t, params, actualParams)

	actualFunctionName = functionNameAsserter.AssertFunctionNameIsValidArgsForCall(0)
	assert.Equal(t, functionName, actualFunctionName)

	assert.Equal(t, 0, formalParametersAsserter.AssertFormalParametersCallCount())
}

func TestFormalParameterAsserterThrowsError(t *testing.T) {
	functionLengthAsserter := &asserterFakes.FakeFunctionLengthAsserter{}
	functionLengthAsserter.AssertLengthIsAsExpectedReturns(nil)

	functionNameAsserter := &fakes2.FakeFunctionNameAsserter{}
	functionNameAsserter.AssertFunctionNameIsValidReturns(nil)

	err := errors.New("formalParameters asserter error")
	formalParametersAsserter := &fakes2.FakeFormalParametersAsserter{}
	formalParametersAsserter.AssertFormalParametersReturns(err)

	functionGenerator := NewFunctionGenerator(
		functionLengthAsserter,
		functionNameAsserter,
		formalParametersAsserter,
	)

	defunNode := datamodels.NewAtomNode(constants.DEFUN)
	functionName := "functionName"
	functionNameNode := datamodels.NewAtomNode(functionName)
	formalParameter := "formalParameter"
	formalParameter1 := datamodels.NewAtomNode(formalParameter)
	formalParameterEnd := datamodels.NewAtomNode(constants.NIL)
	formalParametersNode := datamodels.NewExpressionListNode([]datamodels.Node{formalParameter1, formalParameterEnd})
	params := datamodels.NewExpressionListNode([]datamodels.Node{defunNode, functionNameNode, formalParametersNode})

	actual, actualErr := functionGenerator.GenerateFunction(params)

	assert.Nil(t, actual)
	assert.Equal(t, err, actualErr)

	actualFunctionName, actualLength, actualParams := functionLengthAsserter.AssertLengthIsAsExpectedArgsForCall(0)
	assert.Equal(t, constants.DEFUN, actualFunctionName)
	assert.Equal(t, constants.FOUR, actualLength)
	assert.Equal(t, params, actualParams)

	actualFunctionName = functionNameAsserter.AssertFunctionNameIsValidArgsForCall(0)
	assert.Equal(t, functionName, actualFunctionName)

	actualFormalParameters := formalParametersAsserter.AssertFormalParametersArgsForCall(0)

	assert.Equal(t, 1, len(actualFormalParameters))
	assert.Equal(t, formalParameter, actualFormalParameters[0])
}

func TestNoErrors(t *testing.T) {
	functionLengthAsserter := &asserterFakes.FakeFunctionLengthAsserter{}
	functionLengthAsserter.AssertLengthIsAsExpectedReturns(nil)

	functionNameAsserter := &fakes2.FakeFunctionNameAsserter{}
	functionNameAsserter.AssertFunctionNameIsValidReturns(nil)

	formalParametersAsserter := &fakes2.FakeFormalParametersAsserter{}
	formalParametersAsserter.AssertFormalParametersReturns(nil)

	functionGenerator := NewFunctionGenerator(
		functionLengthAsserter,
		functionNameAsserter,
		formalParametersAsserter,
	)

	defunNode := datamodels.NewAtomNode(constants.DEFUN)
	functionName := "functionName"
	functionNameNode := datamodels.NewAtomNode(functionName)
	formalParameter := "formalParameter"
	formalParameter1 := datamodels.NewAtomNode(formalParameter)
	formalParameterEnd := datamodels.NewAtomNode(constants.NIL)
	formalParametersNode := datamodels.NewExpressionListNode([]datamodels.Node{formalParameter1, formalParameterEnd})
	body := "body"
	bodyNode := datamodels.NewAtomNode(body)
	params := datamodels.NewExpressionListNode([]datamodels.Node{defunNode, functionNameNode, formalParametersNode, bodyNode})

	actual, actualErr := functionGenerator.GenerateFunction(params)

	assert.Equal(t, functionName, actual.GetFunctionName())
	assert.Equal(t, formalParameter, actual.GetFormalParameters()[0])
	assert.Equal(t, bodyNode, actual.GetBody())
	assert.Nil(t, actualErr)

	actualFunctionName, actualLength, actualParams := functionLengthAsserter.AssertLengthIsAsExpectedArgsForCall(0)
	assert.Equal(t, constants.DEFUN, actualFunctionName)
	assert.Equal(t, constants.FOUR, actualLength)
	assert.Equal(t, params, actualParams)

	actualFunctionName = functionNameAsserter.AssertFunctionNameIsValidArgsForCall(0)
	assert.Equal(t, functionName, actualFunctionName)

	actualFormalParameters := formalParametersAsserter.AssertFormalParametersArgsForCall(0)

	assert.Equal(t, 1, len(actualFormalParameters))
	assert.Equal(t, formalParameter, actualFormalParameters[0])
}


