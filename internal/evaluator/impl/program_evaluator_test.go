package impl

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/core/datamodels/fakes"
	fakes2 "lisp_lambda-go/internal/evaluator/impl/internal/fakes"
	"testing"
)

func TestRootNodeIsAtomNode(t *testing.T) {
	atomNode := &fakes.FakeAtomNode{}
	rootNodes := []datamodels.Node{atomNode}
	userDefinedFunctions := map[string]datamodels.DefunFunction{}

	atomRootNodeAsserter := &fakes2.FakeAtomRootNodeAsserter{}
	err := errors.New("error")
	atomRootNodeAsserter.AssertAtomRootNodeReturns(err)

	rootNodeEvaluator := &fakes2.FakeRootNodeEvaluator{}

	programEvaluator := NewProgramEvaluator(
		atomRootNodeAsserter,
		rootNodeEvaluator,
	)

	actual, actualErr := programEvaluator.Evaluate(
		rootNodes,
		userDefinedFunctions,
	)

	assert.Nil(t, actual)
	assert.Equal(t, err, actualErr)
}

func TestRootNodeIsExpressionListNode(t *testing.T) {
	expressionListNode := &fakes.FakeExpressionListNode{}
	rootNodes := []datamodels.Node{expressionListNode}
	userDefinedFunctions := map[string]datamodels.DefunFunction{}

	atomRootNodeAsserter := &fakes2.FakeAtomRootNodeAsserter{}

	rootNodeEvaluator := &fakes2.FakeRootNodeEvaluator{}
	err := errors.New("error")
	rootNodeEvaluator.EvaluateReturns(nil, err)

	programEvaluator := NewProgramEvaluator(
		atomRootNodeAsserter,
		rootNodeEvaluator,
	)

	actual, actualErr := programEvaluator.Evaluate(
		rootNodes,
		userDefinedFunctions,
	)

	assert.Nil(t, actual)
	assert.Equal(t, err, actualErr)
}


func TestNoErrorsThrown(t *testing.T) {
	atomNode := &fakes.FakeAtomNode{}
	expressionListNode := &fakes.FakeExpressionListNode{}
	rootNodes := []datamodels.Node{atomNode, expressionListNode}
	userDefinedFunctions := map[string]datamodels.DefunFunction{}

	atomRootNodeAsserter := &fakes2.FakeAtomRootNodeAsserter{}
	atomRootNodeAsserter.AssertAtomRootNodeReturns(nil)

	rootNodeEvaluator := &fakes2.FakeRootNodeEvaluator{}
	result := &fakes.FakeAtomNode{}
	rootNodeEvaluator.EvaluateReturns(result, nil)

	programEvaluator := NewProgramEvaluator(
		atomRootNodeAsserter,
		rootNodeEvaluator,
	)

	actual, actualErr := programEvaluator.Evaluate(
		rootNodes,
		userDefinedFunctions,
	)

	assert.Equal(t, atomNode, actual[0])
	assert.Equal(t, result, actual[1])
	assert.Nil(t, actualErr)
}