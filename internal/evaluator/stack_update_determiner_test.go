package evaluator

import (
	"github.com/stretchr/testify/assert"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/core/datamodels/fakes"
	internalFakes "lisp_lambda-go/internal/evaluator/internal/fakes"
	"testing"
)

func TestNodeIsExpressionListWithMoreThanOneChild(t *testing.T) {
	node := &fakes.FakeExpressionListNode{}
	variableMap := map[string]datamodels.Node{}
	evalStack := datamodels.NewNodeStack()
	programStack := datamodels.NewProgramItemStack()

	child0 := &fakes.FakeNode{}
	child1 := &fakes.FakeNode{}
	children := []datamodels.Node{child0, child1}
	node.GetChildrenReturns(children)

	topProgramItemCreator := &internalFakes.FakeTopProgramItemCreator{}
	postEvaluationStackUpdater := &internalFakes.FakePostEvaluationStackUpdater{}

	stackUpdateDeterminer := NewStackUpdateDeterminer(
		topProgramItemCreator,
		postEvaluationStackUpdater,
	)

	stackUpdateDeterminer.DetermineHowToUpdateStacks(
		node,
		variableMap,
		evalStack,
		programStack,
	)

	actualNode, actualVariableMap, actualProgramStack := topProgramItemCreator.CreateTopProgramItemArgsForCall(0)
	assert.Equal(t, node, actualNode)
	assert.Equal(t, variableMap, actualVariableMap)
	assert.Equal(t, programStack, actualProgramStack)
}

func TestNodeIsExpressionListWithOneChild(t *testing.T) {
	node := &fakes.FakeExpressionListNode{}
	variableMap := map[string]datamodels.Node{}
	evalStack := datamodels.NewNodeStack()
	programStack := datamodels.NewProgramItemStack()

	child0 := &fakes.FakeNode{}
	children := []datamodels.Node{child0}
	node.GetChildrenReturns(children)

	topProgramItemCreator := &internalFakes.FakeTopProgramItemCreator{}
	postEvaluationStackUpdater := &internalFakes.FakePostEvaluationStackUpdater{}

	stackUpdateDeterminer := NewStackUpdateDeterminer(
		topProgramItemCreator,
		postEvaluationStackUpdater,
	)

	stackUpdateDeterminer.DetermineHowToUpdateStacks(
		node,
		variableMap,
		evalStack,
		programStack,
	)

	actualNode, actualVariableMap, actualEvalStack, actualProgramStack := postEvaluationStackUpdater.UpdateStacksAfterEvaluationArgsForCall(0)
	assert.Equal(t, child0, actualNode)
	assert.Equal(t, variableMap, actualVariableMap)
	assert.Equal(t, evalStack, actualEvalStack)
	assert.Equal(t, programStack, actualProgramStack)
}

func TestNodeIsAtomNode(t *testing.T) {
	node := &fakes.FakeAtomNode{}
	variableMap := map[string]datamodels.Node{}
	evalStack := datamodels.NewNodeStack()
	programStack := datamodels.NewProgramItemStack()

	topProgramItemCreator := &internalFakes.FakeTopProgramItemCreator{}
	postEvaluationStackUpdater := &internalFakes.FakePostEvaluationStackUpdater{}

	stackUpdateDeterminer := NewStackUpdateDeterminer(
		topProgramItemCreator,
		postEvaluationStackUpdater,
	)

	stackUpdateDeterminer.DetermineHowToUpdateStacks(
		node,
		variableMap,
		evalStack,
		programStack,
	)

	actualNode, actualVariableMap, actualEvalStack, actualProgramStack := postEvaluationStackUpdater.UpdateStacksAfterEvaluationArgsForCall(0)
	assert.Equal(t, node, actualNode)
	assert.Equal(t, variableMap, actualVariableMap)
	assert.Equal(t, evalStack, actualEvalStack)
	assert.Equal(t, programStack, actualProgramStack)
}