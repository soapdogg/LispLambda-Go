package interpreter

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/core/datamodels/fakes"
	topFakes "lisp_lambda-go/internal/fakes"
	internalFakes "lisp_lambda-go/internal/interpreter/internal/fakes"
	"testing"
)

func TestParserThrowsError(t *testing.T) {
	input := "input"

	var tokens []string
	tokenizer := &topFakes.FakeTokenizer{}
	tokenizer.TokenizeReturns(tokens)

	err := errors.New("parser error")
	parser := &topFakes.FakeParser{}
	parser.ParseReturns(nil, err)

	rootNodePartitioner := &internalFakes.FakeRootNodePartitioner{}
	functionGenerator := &topFakes.FakeFunctionGenerator{}
	expressionListLengthAsserter := &topFakes.FakeExpressionListLengthAsserter{}
	programEvaluator := &topFakes.FakeProgramEvaluator{}
	printer := &topFakes.FakeListNotationPrinter{}

	interpreter := NewInterpreter(
		tokenizer,
		parser,
		rootNodePartitioner,
		functionGenerator,
		expressionListLengthAsserter,
		programEvaluator,
		printer,
	)

	actual, actualError := interpreter.Interpret(input)

	assert.Empty(t, actual)
	assert.Equal(t, err, actualError)
}

func TestFunctionGeneratorThrowsError(t *testing.T) {
	input := "input"

	var tokens []string
	tokenizer := &topFakes.FakeTokenizer{}
	tokenizer.TokenizeReturns(tokens)

	var rootNodes []datamodels.Node
	parser := &topFakes.FakeParser{}
	parser.ParseReturns(rootNodes, nil)

	partitionedRootNodes := &fakes.FakePartitionedRootNodes{}
	rootNodePartitioner := &internalFakes.FakeRootNodePartitioner{}
	rootNodePartitioner.PartitionRootNodesReturns(partitionedRootNodes)

	defunNode := &fakes.FakeExpressionListNode{}
	defunNodes := []datamodels.ExpressionListNode {defunNode}
	partitionedRootNodes.GetDefunNodesReturns(defunNodes)

	err := errors.New("function generator error")
	functionGenerator := &topFakes.FakeFunctionGenerator{}
	functionGenerator.GenerateFunctionReturns(nil, err)

	expressionListLengthAsserter := &topFakes.FakeExpressionListLengthAsserter{}
	programEvaluator := &topFakes.FakeProgramEvaluator{}
	printer := &topFakes.FakeListNotationPrinter{}

	interpreter := NewInterpreter(
		tokenizer,
		parser,
		rootNodePartitioner,
		functionGenerator,
		expressionListLengthAsserter,
		programEvaluator,
		printer,
	)

	actual, actualError := interpreter.Interpret(input)

	assert.Empty(t, actual)
	assert.Equal(t, err, actualError)
}

func TestExpressionListLengthAsserterThrowsError(t *testing.T) {
	input := "input"

	var tokens []string
	tokenizer := &topFakes.FakeTokenizer{}
	tokenizer.TokenizeReturns(tokens)

	var rootNodes []datamodels.Node
	parser := &topFakes.FakeParser{}
	parser.ParseReturns(rootNodes, nil)

	partitionedRootNodes := &fakes.FakePartitionedRootNodes{}
	rootNodePartitioner := &internalFakes.FakeRootNodePartitioner{}
	rootNodePartitioner.PartitionRootNodesReturns(partitionedRootNodes)

	defunNode := &fakes.FakeExpressionListNode{}
	defunNodes := []datamodels.ExpressionListNode {defunNode}
	partitionedRootNodes.GetDefunNodesReturns(defunNodes)

	userDefinedFunction := &fakes.FakeDefunFunction{}
	functionGenerator := &topFakes.FakeFunctionGenerator{}
	functionGenerator.GenerateFunctionReturns(userDefinedFunction, nil)

	functionName := "functionName"
	userDefinedFunction.GetFunctionNameReturns(functionName)

	var evaluatableNodes []datamodels.Node
	partitionedRootNodes.GetEvaluatableNodesReturns(evaluatableNodes)

	err := errors.New("expression list length error")
	expressionListLengthAsserter := &topFakes.FakeExpressionListLengthAsserter{}
	expressionListLengthAsserter.AssertLengthIsAsExpectedReturns(err)

	programEvaluator := &topFakes.FakeProgramEvaluator{}
	printer := &topFakes.FakeListNotationPrinter{}

	interpreter := NewInterpreter(
		tokenizer,
		parser,
		rootNodePartitioner,
		functionGenerator,
		expressionListLengthAsserter,
		programEvaluator,
		printer,
	)

	actual, actualError := interpreter.Interpret(input)

	assert.Empty(t, actual)
	assert.Equal(t, err, actualError)
}

func TestProgramEvaluatorThrowsError(t *testing.T) {
	input := "input"

	var tokens []string
	tokenizer := &topFakes.FakeTokenizer{}
	tokenizer.TokenizeReturns(tokens)

	var rootNodes []datamodels.Node
	parser := &topFakes.FakeParser{}
	parser.ParseReturns(rootNodes, nil)

	partitionedRootNodes := &fakes.FakePartitionedRootNodes{}
	rootNodePartitioner := &internalFakes.FakeRootNodePartitioner{}
	rootNodePartitioner.PartitionRootNodesReturns(partitionedRootNodes)

	defunNode := &fakes.FakeExpressionListNode{}
	defunNodes := []datamodels.ExpressionListNode {defunNode}
	partitionedRootNodes.GetDefunNodesReturns(defunNodes)

	userDefinedFunction := &fakes.FakeDefunFunction{}
	functionGenerator := &topFakes.FakeFunctionGenerator{}
	functionGenerator.GenerateFunctionReturns(userDefinedFunction, nil)

	functionName := "functionName"
	userDefinedFunction.GetFunctionNameReturns(functionName)

	var evaluatableNodes []datamodels.Node
	partitionedRootNodes.GetEvaluatableNodesReturns(evaluatableNodes)

	expressionListLengthAsserter := &topFakes.FakeExpressionListLengthAsserter{}
	expressionListLengthAsserter.AssertLengthIsAsExpectedReturns(nil)

	err := errors.New("program evaluator error")
	programEvaluator := &topFakes.FakeProgramEvaluator{}
	programEvaluator.EvaluateReturns(nil, err)

	printer := &topFakes.FakeListNotationPrinter{}

	interpreter := NewInterpreter(
		tokenizer,
		parser,
		rootNodePartitioner,
		functionGenerator,
		expressionListLengthAsserter,
		programEvaluator,
		printer,
	)

	actual, actualError := interpreter.Interpret(input)

	assert.Empty(t, actual)
	assert.Equal(t, err, actualError)
}

func TestNoErrorsThrown(t *testing.T) {
	input := "input"

	var tokens []string
	tokenizer := &topFakes.FakeTokenizer{}
	tokenizer.TokenizeReturns(tokens)

	var rootNodes []datamodels.Node
	parser := &topFakes.FakeParser{}
	parser.ParseReturns(rootNodes, nil)

	partitionedRootNodes := &fakes.FakePartitionedRootNodes{}
	rootNodePartitioner := &internalFakes.FakeRootNodePartitioner{}
	rootNodePartitioner.PartitionRootNodesReturns(partitionedRootNodes)

	defunNode := &fakes.FakeExpressionListNode{}
	defunNodes := []datamodels.ExpressionListNode {defunNode}
	partitionedRootNodes.GetDefunNodesReturns(defunNodes)

	userDefinedFunction := &fakes.FakeDefunFunction{}
	functionGenerator := &topFakes.FakeFunctionGenerator{}
	functionGenerator.GenerateFunctionReturns(userDefinedFunction, nil)

	functionName := "functionName"
	userDefinedFunction.GetFunctionNameReturns(functionName)

	var evaluatableNodes []datamodels.Node
	partitionedRootNodes.GetEvaluatableNodesReturns(evaluatableNodes)

	expressionListLengthAsserter := &topFakes.FakeExpressionListLengthAsserter{}
	expressionListLengthAsserter.AssertLengthIsAsExpectedReturns(nil)

	evaluatedNodes := []datamodels.Node{}
	programEvaluator := &topFakes.FakeProgramEvaluator{}
	programEvaluator.EvaluateReturns(evaluatedNodes, nil)

	result := "result"
	printer := &topFakes.FakeListNotationPrinter{}
	printer.PrintAllInListNotationReturns(result)

	interpreter := NewInterpreter(
		tokenizer,
		parser,
		rootNodePartitioner,
		functionGenerator,
		expressionListLengthAsserter,
		programEvaluator,
		printer,
	)

	actual, actualError := interpreter.Interpret(input)

	assert.Equal(t, result, actual)
	assert.Nil(t, actualError)
}