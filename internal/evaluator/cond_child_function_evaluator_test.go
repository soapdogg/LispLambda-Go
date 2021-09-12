package evaluator

import (
	"github.com/stretchr/testify/assert"
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/core/datamodels/fakes"
	internalFakes "lisp_lambda-go/internal/evaluator/internal/fakes"
	"testing"
)

func TestEvaluateCondChildsCondition(t *testing.T) {
	top := &fakes.FakeProgramItem{}
	evalStack := datamodels.NewNodeStack()
	programStack := datamodels.NewProgramItemStack()

	functionExpressionNode := &fakes.FakeExpressionListNode{}
	top.GetFunctionExpressionNodeReturns(functionExpressionNode)

	grandChild0 := &fakes.FakeNode{}
	grandChild1 := &fakes.FakeNode{}
	grandChild2 := &fakes.FakeNode{}
	grandChildren := []datamodels.Node{grandChild0, grandChild1, grandChild2}
	functionExpressionNode.GetChildrenReturns(grandChildren)

	currentParameterIndex := 0
	top.GetCurrentParameterIndexReturns(currentParameterIndex)

	variableMap := map[string]datamodels.Node{}
	top.GetVariableMapReturns(variableMap)

	stackUpdateDeterminer := &internalFakes.FakeStackUpdateDeterminer{}
	condChildFunctionEvaluator := NewCondChildFunctionEvaluator(
		stackUpdateDeterminer,
	)

	condChildFunctionEvaluator.EvaluateCondChildFunction(
		top,
		evalStack,
		programStack,
	)

	actualGrandChild, actualVariableMap, actualEvalStack, actualProgramStack := stackUpdateDeterminer.DetermineHowToUpdateStacksArgsForCall(0)
	assert.Equal(t, grandChild1, actualGrandChild)
	assert.Equal(t, variableMap, actualVariableMap)
	assert.Equal(t, evalStack, actualEvalStack)
	assert.Equal(t, programStack, actualProgramStack)
}

func TestEvaluateCondChildsValue(t *testing.T) {
	top := &fakes.FakeProgramItem{}
	evalStack := datamodels.NewNodeStack()
	programStack := datamodels.NewProgramItemStack()

	functionExpressionNode := &fakes.FakeExpressionListNode{}
	top.GetFunctionExpressionNodeReturns(functionExpressionNode)

	grandChild0 := &fakes.FakeNode{}
	grandChild1 := &fakes.FakeNode{}
	grandChild2 := &fakes.FakeNode{}
	grandChildren := []datamodels.Node{grandChild0, grandChild1, grandChild2}
	functionExpressionNode.GetChildrenReturns(grandChildren)

	currentParameterIndex := 1
	top.GetCurrentParameterIndexReturns(currentParameterIndex)

	evaluatedCondChild := &fakes.FakeAtomNode{}
	evaluatedCondChild.GetValueReturns(constants.T)
	evalStack.Push(evaluatedCondChild)

	programStackItem1 := &fakes.FakeProgramItem{}
	programStackItem1.GetFunctionNameReturns(constants.COND_CHILD)

	programStackItem0 := &fakes.FakeProgramItem{}
	programStackItem0.GetFunctionNameReturns(constants.COND)

	programStack.Push(programStackItem0)
	programStack.Push(programStackItem1)

	variableMap := map[string]datamodels.Node{}
	top.GetVariableMapReturns(variableMap)

	stackUpdateDeterminer := &internalFakes.FakeStackUpdateDeterminer{}
	condChildFunctionEvaluator := NewCondChildFunctionEvaluator(
		stackUpdateDeterminer,
	)

	condChildFunctionEvaluator.EvaluateCondChildFunction(
		top,
		evalStack,
		programStack,
	)

	actualGrandChild, actualVariableMap, actualEvalStack, actualProgramStack := stackUpdateDeterminer.DetermineHowToUpdateStacksArgsForCall(0)
	assert.Equal(t, grandChild2, actualGrandChild)
	assert.Equal(t, variableMap, actualVariableMap)
	assert.Equal(t, evalStack, actualEvalStack)
	assert.Equal(t, programStack, actualProgramStack)
}