package impl

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/evaluator/impl/internal/fakes"
	function2 "lisp_lambda-go/internal/function"
	functionFakes "lisp_lambda-go/internal/function/fakes"
	"testing"
)

func TestEvaluateBuiltInFunctionNoErrors(t *testing.T) {
	functionStack := datamodels.NewNodeStack()
	functionExpressionNode := datamodels.NewExpressionListNode([]datamodels.Node{})
	variableMap := map[string]datamodels.Node{}
	head := datamodels.NewProgramItem(
		functionExpressionNode,
		1,
		variableMap,
		"functionName",
	)
	evalStack := datamodels.NewNodeStack()
	programStack := datamodels.NewProgramItemStack()

	functionName := "functionName"
	function := &functionFakes.FakeFunction{}
	functionMap := map[string]function2.Function{functionName: function}
	evaluatedNode := datamodels.NewAtomNode("evaluated")
	function.EvaluateReturns(evaluatedNode, nil)

	postEvaluationStackUpdater := &fakes.FakePostEvaluationStackUpdater{}

	builtInFunctionEvaluator := NewBuiltInFunctionEvaluator(
		functionMap,
		postEvaluationStackUpdater,
	)

	err := builtInFunctionEvaluator.EvaluateBuiltInFunction(
		functionName,
		functionStack,
		head,
		evalStack,
		programStack,
	)

	assert.Nil(t, err)

	actualEvaluatedFunctionResult, actualVariableMap, actualEvalStack, actualProgramStack := postEvaluationStackUpdater.UpdateStacksAfterEvaluationArgsForCall(0)
	assert.Equal(t, evaluatedNode, actualEvaluatedFunctionResult)
	assert.Equal(t, variableMap, actualVariableMap)
	assert.Equal(t, evalStack, actualEvalStack)
	assert.Equal(t, programStack, actualProgramStack)
}

func TestEvaluateBuiltInFunctionErrorThrown(t *testing.T) {
	functionStack := datamodels.NewNodeStack()
	functionExpressionNode := datamodels.NewExpressionListNode([]datamodels.Node{})
	variableMap := map[string]datamodels.Node{}
	head := datamodels.NewProgramItem(
		functionExpressionNode,
		1,
		variableMap,
		"functionName",
	)
	evalStack := datamodels.NewNodeStack()
	programStack := datamodels.NewProgramItemStack()

	functionName := "functionName"
	function := &functionFakes.FakeFunction{}
	functionMap := map[string]function2.Function{functionName: function}
	err := errors.New("function error")
	function.EvaluateReturns(nil, err)

	postEvaluationStackUpdater := &fakes.FakePostEvaluationStackUpdater{}

	builtInFunctionEvaluator := NewBuiltInFunctionEvaluator(
		functionMap,
		postEvaluationStackUpdater,
	)

	actual := builtInFunctionEvaluator.EvaluateBuiltInFunction(
		functionName,
		functionStack,
		head,
		evalStack,
		programStack,
	)

	assert.Equal(t, err, actual)

	assert.Equal(t, 0, postEvaluationStackUpdater.UpdateStacksAfterEvaluationCallCount())
}