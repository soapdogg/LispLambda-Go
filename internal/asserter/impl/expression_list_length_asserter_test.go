package impl

import (
	"errors"
	"github.com/stretchr/testify/assert"
	asserterFakes "lisp_lambda-go/internal/asserter/fakes"
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/core/datamodels/fakes"
	"testing"
)

func TestFunctionLengthThrowsError(t *testing.T) {
	node := &fakes.FakeExpressionListNode{}
	nodes := []datamodels.Node{node}

	address := &fakes.FakeAtomNode{}
	node.GetChildrenReturns([]datamodels.Node{address})

	addressValue := "functionName"
	address.GetValueReturns(addressValue)

	userDefinedFunctions := map[string] datamodels.DefunFunction{}

	functionLengthAsserter := &asserterFakes.FakeFunctionLengthAsserter{}
	err := errors.New("error")
	functionLengthAsserter.AssertLengthIsAsExpectedReturns(err)

	functionLengthMap := map[string]int {addressValue: constants.TWO}

	expressionListLengthAsserter := NewExpressionListLengthAsserter(
		functionLengthAsserter,
		functionLengthMap,
		functionLengthMap,
	)

	actual := expressionListLengthAsserter.AssertLengthIsAsExpected(
		nodes,
		userDefinedFunctions,
	)

	assert.Equal(t, err, actual)

	actualValue, actualFunctionLength, actualNode := functionLengthAsserter.AssertLengthIsAsExpectedArgsForCall(0)
	assert.Equal(t, addressValue, actualValue)
	assert.Equal(t, constants.TWO, actualFunctionLength)
	assert.Equal(t, node, actualNode)
}

func TestMinimumFunctionLengthThrowsError(t *testing.T) {
	node := &fakes.FakeExpressionListNode{}
	nodes := []datamodels.Node{node}

	address := &fakes.FakeAtomNode{}
	node.GetChildrenReturns([]datamodels.Node{address})

	addressValue := "functionName"
	address.GetValueReturns(addressValue)

	userDefinedFunctions := map[string] datamodels.DefunFunction{}

	functionLengthAsserter := &asserterFakes.FakeFunctionLengthAsserter{}
	err := errors.New("error")
	functionLengthAsserter.AssertLengthIsAtLeastMinimumReturns(err)

	functionLengthMap := map[string]int {}
	minimumFunctionLengthMap := map[string]int {addressValue: constants.TWO}

	expressionListLengthAsserter := NewExpressionListLengthAsserter(
		functionLengthAsserter,
		functionLengthMap,
		minimumFunctionLengthMap,
	)

	actual := expressionListLengthAsserter.AssertLengthIsAsExpected(
		nodes,
		userDefinedFunctions,
	)

	assert.Equal(t, err, actual)

	actualValue, actualFunctionLength, actualNode := functionLengthAsserter.AssertLengthIsAtLeastMinimumArgsForCall(0)
	assert.Equal(t, addressValue, actualValue)
	assert.Equal(t, constants.TWO, actualFunctionLength)
	assert.Equal(t, node, actualNode)
}

func TestUserDefinedFunctionThrowsError(t *testing.T) {
	node := &fakes.FakeExpressionListNode{}
	nodes := []datamodels.Node{node}

	address := &fakes.FakeAtomNode{}
	node.GetChildrenReturns([]datamodels.Node{address})

	addressValue := "functionName"
	address.GetValueReturns(addressValue)

	userDefinedFunction := &fakes.FakeDefunFunction{}
	userDefinedFunctions := map[string] datamodels.DefunFunction{addressValue: userDefinedFunction}

	formalParameters := []string{}
	userDefinedFunction.GetFormalParametersReturns(formalParameters)


	functionLengthAsserter := &asserterFakes.FakeFunctionLengthAsserter{}
	err := errors.New("error")
	functionLengthAsserter.AssertLengthIsAsExpectedReturns(err)

	functionLengthMap := map[string]int {}
	minimumFunctionLengthMap := map[string]int {}

	expressionListLengthAsserter := NewExpressionListLengthAsserter(
		functionLengthAsserter,
		functionLengthMap,
		minimumFunctionLengthMap,
	)

	actual := expressionListLengthAsserter.AssertLengthIsAsExpected(
		nodes,
		userDefinedFunctions,
	)

	assert.Equal(t, err, actual)

	actualValue, actualFunctionLength, actualNode := functionLengthAsserter.AssertLengthIsAsExpectedArgsForCall(0)
	assert.Equal(t, addressValue, actualValue)
	assert.Equal(t, len(formalParameters) + 1, actualFunctionLength)
	assert.Equal(t, node, actualNode)
}

func TestCondFunctionExpressionListChildrenThrowsError(t *testing.T) {
	node := &fakes.FakeExpressionListNode{}
	nodes := []datamodels.Node{node}

	address := &fakes.FakeAtomNode{}

	addressValue := constants.COND
	address.GetValueReturns(addressValue)

	condChild := &fakes.FakeExpressionListNode{}
	condGrandChild := &fakes.FakeAtomNode{}
	condChild.GetChildrenReturns([]datamodels.Node{condGrandChild})
	nilChild := &fakes.FakeAtomNode{}
	node.GetChildrenReturns([]datamodels.Node{address, condChild, nilChild})

	userDefinedFunctions := map[string] datamodels.DefunFunction{}


	functionLengthAsserter := &asserterFakes.FakeFunctionLengthAsserter{}
	err := errors.New("error")
	functionLengthAsserter.AssertLengthIsAsExpectedReturns(err)

	functionLengthMap := map[string]int {}
	minimumFunctionLengthMap := map[string]int {}

	expressionListLengthAsserter := NewExpressionListLengthAsserter(
		functionLengthAsserter,
		functionLengthMap,
		minimumFunctionLengthMap,
	)

	actual := expressionListLengthAsserter.AssertLengthIsAsExpected(
		nodes,
		userDefinedFunctions,
	)

	assert.Equal(t, err, actual)

	actualValue, actualFunctionLength, actualNode := functionLengthAsserter.AssertLengthIsAsExpectedArgsForCall(0)
	assert.Equal(t, constants.COND, actualValue)
	assert.Equal(t, constants.TWO, actualFunctionLength)
	assert.Equal(t, condChild, actualNode)
}

func TestCondFunctionChildrenThrowsLength(t *testing.T) {
	node := &fakes.FakeExpressionListNode{}
	nodes := []datamodels.Node{node}

	address := &fakes.FakeAtomNode{}

	addressValue := constants.COND
	address.GetValueReturns(addressValue)

	condChild := &fakes.FakeAtomNode{}
	nilChild := &fakes.FakeAtomNode{}
	node.GetChildrenReturns([]datamodels.Node{address, condChild, nilChild})

	userDefinedFunctions := map[string] datamodels.DefunFunction{}


	functionLengthAsserter := &asserterFakes.FakeFunctionLengthAsserter{}

	functionLengthMap := map[string]int {}
	minimumFunctionLengthMap := map[string]int {}

	expressionListLengthAsserter := NewExpressionListLengthAsserter(
		functionLengthAsserter,
		functionLengthMap,
		minimumFunctionLengthMap,
	)

	actual := expressionListLengthAsserter.AssertLengthIsAsExpected(
		nodes,
		userDefinedFunctions,
	)

	assert.NotNil(t, actual)
}


func TestNoErrorThrown(t *testing.T) {
	node := &fakes.FakeExpressionListNode{}
	nodes := []datamodels.Node{node}

	address := &fakes.FakeAtomNode{}
	data := &fakes.FakeExpressionListNode{}
	dataAddressValue := "dataAddressValue"
	dataAddress := &fakes.FakeAtomNode{}
	dataAddress.GetValueReturns(dataAddressValue)
	data.GetChildrenReturns([]datamodels.Node{dataAddress})

	addressValue := "functionName"
	address.GetValueReturns(addressValue)

	node.GetChildrenReturns([]datamodels.Node{address, data})

	userDefinedFunctions := map[string] datamodels.DefunFunction{}

	functionLengthAsserter := &asserterFakes.FakeFunctionLengthAsserter{}
	functionLengthMap := map[string]int {}
	minimumFunctionLengthMap := map[string]int {}

	expressionListLengthAsserter := NewExpressionListLengthAsserter(
		functionLengthAsserter,
		functionLengthMap,
		minimumFunctionLengthMap,
	)

	actual := expressionListLengthAsserter.AssertLengthIsAsExpected(
		nodes,
		userDefinedFunctions,
	)

	assert.Nil(t, actual)
}
