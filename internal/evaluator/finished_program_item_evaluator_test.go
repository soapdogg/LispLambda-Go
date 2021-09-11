package evaluator

import (
	"errors"
	"github.com/stretchr/testify/assert"
	top "lisp_lambda-go/internal"
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/core/datamodels/fakes"
	internalFakes "lisp_lambda-go/internal/evaluator/internal/fakes"
	topFakes "lisp_lambda-go/internal/fakes"
	"testing"
)

func TestFunctionNameIsNil(t *testing.T) {
	head := &fakes.FakeProgramItem{}
	userDefinedFunctions := map[string]datamodels.DefunFunction{}
	evalStack := datamodels.NewNodeStack()
	programStack := datamodels.NewProgramItemStack()

	functionExpressionNode := &fakes.FakeExpressionListNode{}
	head.GetFunctionExpressionNodeReturns(functionExpressionNode)

	child0 := &fakes.FakeNode{}
	child1 := &fakes.FakeNode{}
	functionExpressionNode.GetChildrenReturns([]datamodels.Node{child0, child1})
	evalStack.Push(child0)
	evalStack.Push(child1)

	variableMap := map[string]datamodels.Node{}
	head.GetVariableMapReturns(variableMap)

	functionName := constants.NIL
	head.GetFunctionNameReturns(functionName)

	postEvaluationStackUpdater := &internalFakes.FakePostEvaluationStackUpdater{}
	functionMap := map[string]top.Function{}
	builtInFunctionEvaluator := &internalFakes.FakeBuiltInFunctionEvaluator{}
	userDefinedFunctionEvaluator := &internalFakes.FakeUserDefinedFunctionEvaluator{}

	finishedProgramItemEvaluator := NewFinishedProgramItemEvaluator(
		postEvaluationStackUpdater,
		functionMap,
		builtInFunctionEvaluator,
		userDefinedFunctionEvaluator,
	)

	actual := finishedProgramItemEvaluator.EvaluateFinishedProgramItem(
		head,
		userDefinedFunctions,
		evalStack,
		programStack,
	)

	assert.Nil(t, actual)
}

func TestFunctionNameIsBuiltIn(t *testing.T) {
	head := &fakes.FakeProgramItem{}
	userDefinedFunctions := map[string]datamodels.DefunFunction{}
	evalStack := datamodels.NewNodeStack()
	programStack := datamodels.NewProgramItemStack()

	functionExpressionNode := &fakes.FakeExpressionListNode{}
	head.GetFunctionExpressionNodeReturns(functionExpressionNode)

	child0 := &fakes.FakeNode{}
	child1 := &fakes.FakeNode{}
	functionExpressionNode.GetChildrenReturns([]datamodels.Node{child0, child1})
	evalStack.Push(child0)
	evalStack.Push(child1)

	variableMap := map[string]datamodels.Node{}
	head.GetVariableMapReturns(variableMap)

	functionName := constants.ATOM
	head.GetFunctionNameReturns(functionName)

	postEvaluationStackUpdater := &internalFakes.FakePostEvaluationStackUpdater{}
	function := &topFakes.FakeFunction{}
	functionMap := map[string]top.Function{functionName: function}
	builtInFunctionEvaluator := &internalFakes.FakeBuiltInFunctionEvaluator{}
	err := errors.New("error")
	builtInFunctionEvaluator.EvaluateBuiltInFunctionReturns(err)
	userDefinedFunctionEvaluator := &internalFakes.FakeUserDefinedFunctionEvaluator{}

	finishedProgramItemEvaluator := NewFinishedProgramItemEvaluator(
		postEvaluationStackUpdater,
		functionMap,
		builtInFunctionEvaluator,
		userDefinedFunctionEvaluator,
	)

	actual := finishedProgramItemEvaluator.EvaluateFinishedProgramItem(
		head,
		userDefinedFunctions,
		evalStack,
		programStack,
	)

	assert.Equal(t, err, actual)
}

func TestFunctionNameIsUserDefined(t *testing.T) {
	head := &fakes.FakeProgramItem{}
	functionName := "userDefined"
	head.GetFunctionNameReturns(functionName)
	userDefinedFunction := &fakes.FakeDefunFunction{}
	userDefinedFunctions := map[string]datamodels.DefunFunction{functionName: userDefinedFunction}
	evalStack := datamodels.NewNodeStack()
	programStack := datamodels.NewProgramItemStack()

	functionExpressionNode := &fakes.FakeExpressionListNode{}
	head.GetFunctionExpressionNodeReturns(functionExpressionNode)

	child0 := &fakes.FakeNode{}
	child1 := &fakes.FakeNode{}
	functionExpressionNode.GetChildrenReturns([]datamodels.Node{child0, child1})
	evalStack.Push(child0)
	evalStack.Push(child1)

	variableMap := map[string]datamodels.Node{}
	head.GetVariableMapReturns(variableMap)

	postEvaluationStackUpdater := &internalFakes.FakePostEvaluationStackUpdater{}
	functionMap := map[string]top.Function{}
	builtInFunctionEvaluator := &internalFakes.FakeBuiltInFunctionEvaluator{}
	userDefinedFunctionEvaluator := &internalFakes.FakeUserDefinedFunctionEvaluator{}

	finishedProgramItemEvaluator := NewFinishedProgramItemEvaluator(
		postEvaluationStackUpdater,
		functionMap,
		builtInFunctionEvaluator,
		userDefinedFunctionEvaluator,
	)

	actual := finishedProgramItemEvaluator.EvaluateFinishedProgramItem(
		head,
		userDefinedFunctions,
		evalStack,
		programStack,
	)

	assert.Nil(t, actual)
}

func TestFunctionNameIsInvalid(t *testing.T) {
	head := &fakes.FakeProgramItem{}
	functionName := "invalid"
	head.GetFunctionNameReturns(functionName)
	userDefinedFunctions := map[string]datamodels.DefunFunction{}
	evalStack := datamodels.NewNodeStack()
	programStack := datamodels.NewProgramItemStack()

	functionExpressionNode := &fakes.FakeExpressionListNode{}
	head.GetFunctionExpressionNodeReturns(functionExpressionNode)

	child0 := &fakes.FakeNode{}
	child1 := &fakes.FakeNode{}
	functionExpressionNode.GetChildrenReturns([]datamodels.Node{child0, child1})
	evalStack.Push(child0)
	evalStack.Push(child1)

	variableMap := map[string]datamodels.Node{}
	head.GetVariableMapReturns(variableMap)

	postEvaluationStackUpdater := &internalFakes.FakePostEvaluationStackUpdater{}
	functionMap := map[string]top.Function{}
	builtInFunctionEvaluator := &internalFakes.FakeBuiltInFunctionEvaluator{}
	userDefinedFunctionEvaluator := &internalFakes.FakeUserDefinedFunctionEvaluator{}

	finishedProgramItemEvaluator := NewFinishedProgramItemEvaluator(
		postEvaluationStackUpdater,
		functionMap,
		builtInFunctionEvaluator,
		userDefinedFunctionEvaluator,
	)

	actual := finishedProgramItemEvaluator.EvaluateFinishedProgramItem(
		head,
		userDefinedFunctions,
		evalStack,
		programStack,
	)

	assert.NotNil(t, actual)
}