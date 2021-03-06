package impl

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/core/datamodels/fakes"
	fakes2 "lisp_lambda-go/internal/evaluator/impl/internal/fakes"
	"testing"
)

func TestEvaluateCondFunction(t *testing.T) {
	rootNode := &fakes.FakeExpressionListNode{}
	userDefinedFunctions := map[string]datamodels.DefunFunction{}
	programStack := datamodels.NewProgramItemStack()
	evalStack := datamodels.NewNodeStack()

	rootProgramStackItem := &fakes.FakeProgramItem{}
	programStack.Push(rootProgramStackItem)

	functionName := constants.COND
	rootProgramStackItem.GetFunctionNameReturns(functionName)

	err := errors.New("cond error")
	condFunctionEvaluator := &fakes2.FakeCondFunctionEvaluator{}
	condFunctionEvaluator.EvaluateCondProgramItemReturns(err)

	topProgramItemCreator := &fakes2.FakeTopProgramItemCreator{}
	condChildFunctionEvaluator := &fakes2.FakeCondChildFunctionEvaluator{}
	quoteFunctionEvaluator := &fakes2.FakeQuoteFunctionEvaluator{}
	unfinishedProgramItemEvaluator := &fakes2.FakeUnfinishedProgramStackItemEvaluator{}
	finishedProgramItemEvaluator := &fakes2.FakeFinishedProgramItemEvaluator{}

	rootNodeEvaluator := NewRootNodeEvaluator(
		topProgramItemCreator,
		condFunctionEvaluator,
		condChildFunctionEvaluator,
		quoteFunctionEvaluator,
		unfinishedProgramItemEvaluator,
		finishedProgramItemEvaluator,
	)

	actual, actualErr := rootNodeEvaluator.Evaluate(
		rootNode,
		userDefinedFunctions,
		programStack,
		evalStack,
	)

	assert.Nil(t, actual)
	assert.Equal(t, err, actualErr)
}

func TestEvaluateCondChildFunction(t *testing.T) {
	rootNode := &fakes.FakeExpressionListNode{}
	userDefinedFunctions := map[string]datamodels.DefunFunction{}
	programStack := datamodels.NewProgramItemStack()
	evalStack := datamodels.NewNodeStack()

	rootProgramStackItem := &fakes.FakeProgramItem{}
	programStack.Push(rootProgramStackItem)

	functionName := constants.COND_CHILD
	rootProgramStackItem.GetFunctionNameReturns(functionName)

	condFunctionEvaluator := &fakes2.FakeCondFunctionEvaluator{}
	topProgramItemCreator := &fakes2.FakeTopProgramItemCreator{}
	condChildFunctionEvaluator := &fakes2.FakeCondChildFunctionEvaluator{}
	result := &fakes.FakeNode{}
	evalStack.Push(result)
	quoteFunctionEvaluator := &fakes2.FakeQuoteFunctionEvaluator{}
	unfinishedProgramItemEvaluator := &fakes2.FakeUnfinishedProgramStackItemEvaluator{}
	finishedProgramItemEvaluator := &fakes2.FakeFinishedProgramItemEvaluator{}

	rootNodeEvaluator := NewRootNodeEvaluator(
		topProgramItemCreator,
		condFunctionEvaluator,
		condChildFunctionEvaluator,
		quoteFunctionEvaluator,
		unfinishedProgramItemEvaluator,
		finishedProgramItemEvaluator,
	)

	actual, actualErr := rootNodeEvaluator.Evaluate(
		rootNode,
		userDefinedFunctions,
		programStack,
		evalStack,
	)

	assert.Equal(t, result, actual)
	assert.Nil(t, actualErr)
}

func TestEvaluateQuoteFunctionName(t *testing.T) {
	rootNode := &fakes.FakeExpressionListNode{}
	userDefinedFunctions := map[string]datamodels.DefunFunction{}
	programStack := datamodels.NewProgramItemStack()
	evalStack := datamodels.NewNodeStack()

	rootProgramStackItem := &fakes.FakeProgramItem{}
	programStack.Push(rootProgramStackItem)

	functionName := constants.QUOTE
	rootProgramStackItem.GetFunctionNameReturns(functionName)

	condFunctionEvaluator := &fakes2.FakeCondFunctionEvaluator{}
	topProgramItemCreator := &fakes2.FakeTopProgramItemCreator{}
	condChildFunctionEvaluator := &fakes2.FakeCondChildFunctionEvaluator{}

	quoteFunctionEvaluator := &fakes2.FakeQuoteFunctionEvaluator{}
	result := &fakes.FakeNode{}
	evalStack.Push(result)
	unfinishedProgramItemEvaluator := &fakes2.FakeUnfinishedProgramStackItemEvaluator{}
	finishedProgramItemEvaluator := &fakes2.FakeFinishedProgramItemEvaluator{}

	rootNodeEvaluator := NewRootNodeEvaluator(
		topProgramItemCreator,
		condFunctionEvaluator,
		condChildFunctionEvaluator,
		quoteFunctionEvaluator,
		unfinishedProgramItemEvaluator,
		finishedProgramItemEvaluator,
	)

	actual, actualErr := rootNodeEvaluator.Evaluate(
		rootNode,
		userDefinedFunctions,
		programStack,
		evalStack,
	)

	assert.Equal(t, result, actual)
	assert.Nil(t, actualErr)
}

func TestEvaluateUnfinishedProgram(t *testing.T) {
	rootNode := &fakes.FakeExpressionListNode{}
	userDefinedFunctions := map[string]datamodels.DefunFunction{}
	programStack := datamodels.NewProgramItemStack()
	evalStack := datamodels.NewNodeStack()

	rootProgramStackItem := &fakes.FakeProgramItem{}
	programStack.Push(rootProgramStackItem)

	functionName := constants.ATOM
	rootProgramStackItem.GetFunctionNameReturns(functionName)

	currentParameterIndex := 0
	rootProgramStackItem.GetCurrentParameterIndexReturns(currentParameterIndex)

	functionExpressionListNode := &fakes.FakeExpressionListNode{}
	rootProgramStackItem.GetFunctionExpressionNodeReturns(functionExpressionListNode)

	child0 := &fakes.FakeNode{}
	child1 := &fakes.FakeNode{}
	children := []datamodels.Node{child0, child1}
	functionExpressionListNode.GetChildrenReturns(children)

	condFunctionEvaluator := &fakes2.FakeCondFunctionEvaluator{}
	topProgramItemCreator := &fakes2.FakeTopProgramItemCreator{}
	condChildFunctionEvaluator := &fakes2.FakeCondChildFunctionEvaluator{}
	quoteFunctionEvaluator := &fakes2.FakeQuoteFunctionEvaluator{}
	unfinishedProgramItemEvaluator := &fakes2.FakeUnfinishedProgramStackItemEvaluator{}
	result := &fakes.FakeNode{}
	evalStack.Push(result)
	finishedProgramItemEvaluator := &fakes2.FakeFinishedProgramItemEvaluator{}

	rootNodeEvaluator := NewRootNodeEvaluator(
		topProgramItemCreator,
		condFunctionEvaluator,
		condChildFunctionEvaluator,
		quoteFunctionEvaluator,
		unfinishedProgramItemEvaluator,
		finishedProgramItemEvaluator,
	)

	actual, actualErr := rootNodeEvaluator.Evaluate(
		rootNode,
		userDefinedFunctions,
		programStack,
		evalStack,
	)

	assert.Equal(t, result, actual)
	assert.Nil(t, actualErr)
}

func TestEvaluateFinishedProgram(t *testing.T) {
	rootNode := &fakes.FakeExpressionListNode{}
	userDefinedFunctions := map[string]datamodels.DefunFunction{}
	programStack := datamodels.NewProgramItemStack()
	evalStack := datamodels.NewNodeStack()

	rootProgramStackItem := &fakes.FakeProgramItem{}
	programStack.Push(rootProgramStackItem)

	functionName := constants.ATOM
	rootProgramStackItem.GetFunctionNameReturns(functionName)

	currentParameterIndex := 1
	rootProgramStackItem.GetCurrentParameterIndexReturns(currentParameterIndex)

	functionExpressionListNode := &fakes.FakeExpressionListNode{}
	rootProgramStackItem.GetFunctionExpressionNodeReturns(functionExpressionListNode)

	child0 := &fakes.FakeNode{}
	child1 := &fakes.FakeNode{}
	children := []datamodels.Node{child0, child1}
	functionExpressionListNode.GetChildrenReturns(children)

	condFunctionEvaluator := &fakes2.FakeCondFunctionEvaluator{}
	topProgramItemCreator := &fakes2.FakeTopProgramItemCreator{}
	condChildFunctionEvaluator := &fakes2.FakeCondChildFunctionEvaluator{}
	quoteFunctionEvaluator := &fakes2.FakeQuoteFunctionEvaluator{}
	unfinishedProgramItemEvaluator := &fakes2.FakeUnfinishedProgramStackItemEvaluator{}
	finishedProgramItemEvaluator := &fakes2.FakeFinishedProgramItemEvaluator{}
	err := errors.New("error")
	finishedProgramItemEvaluator.EvaluateFinishedProgramItemReturns(err)

	rootNodeEvaluator := NewRootNodeEvaluator(
		topProgramItemCreator,
		condFunctionEvaluator,
		condChildFunctionEvaluator,
		quoteFunctionEvaluator,
		unfinishedProgramItemEvaluator,
		finishedProgramItemEvaluator,
	)

	actual, actualErr := rootNodeEvaluator.Evaluate(
		rootNode,
		userDefinedFunctions,
		programStack,
		evalStack,
	)

	assert.Nil(t, actual)
	assert.Equal(t, err,  actualErr)
}