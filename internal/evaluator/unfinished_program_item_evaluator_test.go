package evaluator

import (
	"github.com/stretchr/testify/assert"
	"lisp_lambda-go/internal/core/datamodels"
	datamodelsFakes "lisp_lambda-go/internal/core/datamodels/fakes"
	"lisp_lambda-go/internal/evaluator/internal/fakes"
	"testing"
)

func TestEvaluateUnfinishedProgramItem(t *testing.T) {
	top := &datamodelsFakes.FakeProgramItem{}
	evalStack := datamodels.NewNodeStack()
	programStack := datamodels.NewProgramItemStack()

	functionExpressionNode := &datamodelsFakes.FakeExpressionListNode{}
	top.GetFunctionExpressionNodeReturns(functionExpressionNode)

	currentParameterIndex := 0
	top.GetCurrentParameterIndexReturns(currentParameterIndex)

	child0 := &datamodelsFakes.FakeNode{}
	functionExpressionNode.GetChildrenReturns([]datamodels.Node{child0})

	variableMap := map[string]datamodels.Node{}
	top.GetVariableMapReturns(variableMap)

	stackUpdateDeterminer := &fakes.FakeStackUpdateDeterminer{}
	unfinishedProgramItemEvaluator := NewUnfinishedProgramItemEvaluator(
		stackUpdateDeterminer,
	)
	unfinishedProgramItemEvaluator.EvaluateUnfinishedProgramItem(
		top,
		evalStack,
		programStack,
	)

	actualChild0, actualVariableMap, actualEvalStack, actualProgramStack := stackUpdateDeterminer.DetermineHowToUpdateStacksArgsForCall(0)
	assert.Equal(t, child0, actualChild0)
	assert.Equal(t, variableMap, actualVariableMap)
	assert.Equal(t, evalStack, actualEvalStack)
	assert.Equal(t, programStack, actualProgramStack)
}