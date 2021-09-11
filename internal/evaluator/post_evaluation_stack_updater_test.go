package evaluator

import (
	"github.com/stretchr/testify/assert"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/core/datamodels/fakes"
	internalFakes "lisp_lambda-go/internal/evaluator/internal/fakes"
	"testing"
)

func TestAtomNodeIsVariable(t *testing.T) {
	evaluatedNode := &fakes.FakeAtomNode{}
	variableName := "variableName"
	evaluatedNode.GetValueReturns(variableName)
	variable := &fakes.FakeNode{}
	variableMap := map[string]datamodels.Node{variableName: variable}
	evalStack := datamodels.NewNodeStack()
	programStack := datamodels.NewProgramItemStack()

	topProgramItemUpdater := &internalFakes.FakeTopProgramItemUpdater{}
	postEvaluationStackUpdater := NewPostEvaluationStackUpdater(
		topProgramItemUpdater,
	)

	postEvaluationStackUpdater.UpdateStacksAfterEvaluation(
		evaluatedNode,
		variableMap,
		evalStack,
		programStack,
	)

	assert.Equal(t, evalStack.Pop(), variable)
}

func TestAtomNodeIsNotVariable(t *testing.T) {
	evaluatedNode := &fakes.FakeAtomNode{}
	value := "value"
	evaluatedNode.GetValueReturns(value)
	variableMap := map[string]datamodels.Node{}
	evalStack := datamodels.NewNodeStack()
	programStack := datamodels.NewProgramItemStack()

	topProgramItemUpdater := &internalFakes.FakeTopProgramItemUpdater{}
	postEvaluationStackUpdater := NewPostEvaluationStackUpdater(
		topProgramItemUpdater,
	)

	postEvaluationStackUpdater.UpdateStacksAfterEvaluation(
		evaluatedNode,
		variableMap,
		evalStack,
		programStack,
	)

	assert.Equal(t, evalStack.Pop(), evaluatedNode)
}

func TestExpressionNode(t *testing.T) {
	evaluatedNode := &fakes.FakeExpressionListNode{}
	variableMap := map[string]datamodels.Node{}
	evalStack := datamodels.NewNodeStack()
	programStack := datamodels.NewProgramItemStack()

	topProgramItemUpdater := &internalFakes.FakeTopProgramItemUpdater{}
	postEvaluationStackUpdater := NewPostEvaluationStackUpdater(
		topProgramItemUpdater,
	)

	postEvaluationStackUpdater.UpdateStacksAfterEvaluation(
		evaluatedNode,
		variableMap,
		evalStack,
		programStack,
	)

	assert.Equal(t, evalStack.Pop(), evaluatedNode)
}