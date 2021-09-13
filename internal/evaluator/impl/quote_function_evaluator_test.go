package impl

import (
	"github.com/stretchr/testify/assert"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/evaluator/impl/internal/fakes"
	"testing"
)

func TestEvaluateQuoteFunction(t *testing.T) {

	evalStack := datamodels.NewNodeStack()
	programStack := datamodels.NewProgramItemStack()

	child0 := datamodels.NewAtomNode("value")
	quoted := datamodels.NewAtomNode("quoted")
	quoteExprNodeChildren := []datamodels.Node{child0, quoted}
	quoteExprNode := datamodels.NewExpressionListNode(quoteExprNodeChildren)

	variableMap := map[string]datamodels.Node{}
	top := datamodels.NewProgramItem(
		quoteExprNode,
		1,
		variableMap,
		"functionName",
	)

	postEvaluationStackUpdater := &fakes.FakePostEvaluationStackUpdater{}
	quoteFunctionEvaluator := NewQuoteFunctionEvaluator(
		postEvaluationStackUpdater,
	)
	
	quoteFunctionEvaluator.EvaluateQuoteFunction(
		top, 
		evalStack,
		programStack,
	)
	
	actualQuoted, actualVariableMap, actualEvalStack, actualProgramStack := postEvaluationStackUpdater.UpdateStacksAfterEvaluationArgsForCall(0)
	assert.Equal(t, quoted, actualQuoted)
	assert.Equal(t, variableMap, actualVariableMap)
	assert.Equal(t, evalStack, actualEvalStack)
	assert.Equal(t, programStack, actualProgramStack)
}