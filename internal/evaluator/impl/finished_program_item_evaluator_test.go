package impl

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/core/datamodels/fakes"
	fakes2 "lisp_lambda-go/internal/evaluator/impl/internal/fakes"
	function2 "lisp_lambda-go/internal/function"
	functionFakes "lisp_lambda-go/internal/function/fakes"
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

	postEvaluationStackUpdater := &fakes2.FakePostEvaluationStackUpdater{}
	functionMap := map[string]function2.Function{}
	builtInFunctionEvaluator := &fakes2.FakeBuiltInFunctionEvaluator{}
	userDefinedFunctionEvaluator := &fakes2.FakeUserDefinedFunctionEvaluator{}

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

	postEvaluationStackUpdater := &fakes2.FakePostEvaluationStackUpdater{}
	function := &functionFakes.FakeFunction{}
	functionMap := map[string]function2.Function{functionName: function}
	builtInFunctionEvaluator := &fakes2.FakeBuiltInFunctionEvaluator{}
	err := errors.New("error")
	builtInFunctionEvaluator.EvaluateBuiltInFunctionReturns(err)
	userDefinedFunctionEvaluator := &fakes2.FakeUserDefinedFunctionEvaluator{}

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

	postEvaluationStackUpdater := &fakes2.FakePostEvaluationStackUpdater{}
	functionMap := map[string]function2.Function{}
	builtInFunctionEvaluator := &fakes2.FakeBuiltInFunctionEvaluator{}
	userDefinedFunctionEvaluator := &fakes2.FakeUserDefinedFunctionEvaluator{}

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

	postEvaluationStackUpdater := &fakes2.FakePostEvaluationStackUpdater{}
	functionMap := map[string]function2.Function{}
	builtInFunctionEvaluator := &fakes2.FakeBuiltInFunctionEvaluator{}
	userDefinedFunctionEvaluator := &fakes2.FakeUserDefinedFunctionEvaluator{}

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