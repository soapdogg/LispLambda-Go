package evaluator

import (
	"github.com/stretchr/testify/assert"
	"lisp_lambda-go/internal/core/datamodels"
	"testing"
)

func TestUpdateTopProgramItem(t *testing.T) {
	functionExpressionNode := datamodels.NewExpressionListNode([]datamodels.Node{})
	currentParameterIndex := 1
	variableMap := map[string] datamodels.Node{}
	functionName := "functionName"
	programItem := datamodels.NewProgramItem(
		functionExpressionNode,
		currentParameterIndex,
		variableMap,
		functionName,
	)
	programStack := datamodels.NewProgramItemStack()
	programStack.Push(programItem)


	topProgramItemUpdater := NewTopProgramItemUpdater()

	topProgramItemUpdater.UpdateTopProgramItemToNextChild(programStack)

	assert.Equal(t, 1, programStack.Size())
	assert.Equal(t, functionExpressionNode, programStack.Peek().GetFunctionExpressionNode())
	assert.Equal(t, currentParameterIndex + 1, programStack.Peek().GetCurrentParameterIndex())
	assert.Equal(t, variableMap, programStack.Peek().GetVariableMap())
	assert.Equal(t, functionName, programStack.Peek().GetFunctionName())
}

func TestUpdateTopProgramItemIsEmpty(t *testing.T) {
	programStack := datamodels.NewProgramItemStack()

	topProgramItemUpdater := NewTopProgramItemUpdater()
	topProgramItemUpdater.UpdateTopProgramItemToNextChild(programStack)

	assert.Equal(t, 0, programStack.Size())
}
