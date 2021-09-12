package evaluator

import (
	"github.com/stretchr/testify/assert"
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/core/datamodels/fakes"
	internalFakes "lisp_lambda-go/internal/evaluator/internal/fakes"
	"testing"
)

func TestBuildCondChildStackItems(t *testing.T) {
	condProgramItem := &fakes.FakeProgramItem{}
	programStack := datamodels.NewProgramItemStack()

	variableMap := map[string]datamodels.Node{}
	condProgramItem.GetVariableMapReturns(variableMap)

	functionExpressionListNode := &fakes.FakeExpressionListNode{}
	condProgramItem.GetFunctionExpressionNodeReturns(functionExpressionListNode)

	child0 := &fakes.FakeNode{}
	child1 := &fakes.FakeExpressionListNode{}
	child2 := &fakes.FakeNode{}
	condChildren := []datamodels.Node{child0, child1, child2}
	functionExpressionListNode.GetChildrenReturns(condChildren)

	grandchild := &fakes.FakeNode{}
	child1Children := []datamodels.Node{grandchild}
	child1.GetChildrenReturns(child1Children)

	topProgramItemCreator := &internalFakes.FakeTopProgramItemCreator{}
	condChildStackItemBuilder := NewCondChildStackItemBuilder(
		topProgramItemCreator,
	)

 	condChildStackItemBuilder.BuildCondChildStackItems(
 		condProgramItem,
 		programStack,
	)

	actualCondChildExpressionListNode, actualVariableMap, actualProgramStack := topProgramItemCreator.CreateTopProgramItemArgsForCall(0)

	assert.Equal(t, 2, len(actualCondChildExpressionListNode.GetChildren()))
	assert.Equal(t, constants.COND_CHILD, actualCondChildExpressionListNode.GetChildren()[0].(datamodels.AtomNode).GetValue())
	assert.Equal(t, child1Children, actualCondChildExpressionListNode.GetChildren()[1])
	assert.Equal(t, variableMap, actualVariableMap)
	assert.Equal(t, programStack, actualProgramStack)
}