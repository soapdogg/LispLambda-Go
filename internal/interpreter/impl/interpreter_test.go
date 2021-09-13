package impl

import (
	"errors"
	"github.com/stretchr/testify/assert"
	asserterFakes "lisp_lambda-go/internal/asserter/fakes"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/core/datamodels/fakes"
	evaluatorFakes "lisp_lambda-go/internal/evaluator/fakes"
	internalFakes "lisp_lambda-go/internal/interpreter/impl/internal/fakes"
	parserFakes "lisp_lambda-go/internal/parser/fakes"
	printerFakes "lisp_lambda-go/internal/printer/fakes"
	tokenizerFakes "lisp_lambda-go/internal/tokenizer/fakes"
	userDefinedFakes "lisp_lambda-go/internal/userdefined/fakes"
	"testing"
)

func TestParserThrowsError(t *testing.T) {
	input := "input"

	var tokens []string
	tokenizer := &tokenizerFakes.FakeTokenizer{}
	tokenizer.TokenizeReturns(tokens)

	err := errors.New("parser error")
	parser := &parserFakes.FakeParser{}
	parser.ParseReturns(nil, err)

	rootNodePartitioner := &internalFakes.FakeRootNodePartitioner{}
	functionGenerator := &userDefinedFakes.FakeFunctionGenerator{}
	expressionListLengthAsserter := &asserterFakes.FakeExpressionListLengthAsserter{}
	programEvaluator := &evaluatorFakes.FakeProgramEvaluator{}
	printer := &printerFakes.FakeListNotationPrinter{}

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
	tokenizer := &tokenizerFakes.FakeTokenizer{}
	tokenizer.TokenizeReturns(tokens)

	var rootNodes []datamodels.Node
	parser := &parserFakes.FakeParser{}
	parser.ParseReturns(rootNodes, nil)

	partitionedRootNodes := &fakes.FakePartitionedRootNodes{}
	rootNodePartitioner := &internalFakes.FakeRootNodePartitioner{}
	rootNodePartitioner.PartitionRootNodesReturns(partitionedRootNodes)

	defunNode := &fakes.FakeExpressionListNode{}
	defunNodes := []datamodels.ExpressionListNode {defunNode}
	partitionedRootNodes.GetDefunNodesReturns(defunNodes)

	err := errors.New("function generator error")
	functionGenerator := &userDefinedFakes.FakeFunctionGenerator{}
	functionGenerator.GenerateFunctionReturns(nil, err)

	expressionListLengthAsserter := &asserterFakes.FakeExpressionListLengthAsserter{}
	programEvaluator := &evaluatorFakes.FakeProgramEvaluator{}
	printer := &printerFakes.FakeListNotationPrinter{}

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
	tokenizer := &tokenizerFakes.FakeTokenizer{}
	tokenizer.TokenizeReturns(tokens)

	var rootNodes []datamodels.Node
	parser := &parserFakes.FakeParser{}
	parser.ParseReturns(rootNodes, nil)

	partitionedRootNodes := &fakes.FakePartitionedRootNodes{}
	rootNodePartitioner := &internalFakes.FakeRootNodePartitioner{}
	rootNodePartitioner.PartitionRootNodesReturns(partitionedRootNodes)

	defunNode := &fakes.FakeExpressionListNode{}
	defunNodes := []datamodels.ExpressionListNode {defunNode}
	partitionedRootNodes.GetDefunNodesReturns(defunNodes)

	userDefinedFunction := &fakes.FakeDefunFunction{}
	functionGenerator := &userDefinedFakes.FakeFunctionGenerator{}
	functionGenerator.GenerateFunctionReturns(userDefinedFunction, nil)

	functionName := "functionName"
	userDefinedFunction.GetFunctionNameReturns(functionName)

	var evaluatableNodes []datamodels.Node
	partitionedRootNodes.GetEvaluatableNodesReturns(evaluatableNodes)

	err := errors.New("expression list length error")
	expressionListLengthAsserter := &asserterFakes.FakeExpressionListLengthAsserter{}
	expressionListLengthAsserter.AssertLengthIsAsExpectedReturns(err)

	programEvaluator := &evaluatorFakes.FakeProgramEvaluator{}
	printer := &printerFakes.FakeListNotationPrinter{}

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
	tokenizer := &tokenizerFakes.FakeTokenizer{}
	tokenizer.TokenizeReturns(tokens)

	var rootNodes []datamodels.Node
	parser := &parserFakes.FakeParser{}
	parser.ParseReturns(rootNodes, nil)

	partitionedRootNodes := &fakes.FakePartitionedRootNodes{}
	rootNodePartitioner := &internalFakes.FakeRootNodePartitioner{}
	rootNodePartitioner.PartitionRootNodesReturns(partitionedRootNodes)

	defunNode := &fakes.FakeExpressionListNode{}
	defunNodes := []datamodels.ExpressionListNode {defunNode}
	partitionedRootNodes.GetDefunNodesReturns(defunNodes)

	userDefinedFunction := &fakes.FakeDefunFunction{}
	functionGenerator := &userDefinedFakes.FakeFunctionGenerator{}
	functionGenerator.GenerateFunctionReturns(userDefinedFunction, nil)

	functionName := "functionName"
	userDefinedFunction.GetFunctionNameReturns(functionName)

	var evaluatableNodes []datamodels.Node
	partitionedRootNodes.GetEvaluatableNodesReturns(evaluatableNodes)

	expressionListLengthAsserter := &asserterFakes.FakeExpressionListLengthAsserter{}
	expressionListLengthAsserter.AssertLengthIsAsExpectedReturns(nil)

	err := errors.New("program evaluator error")
	programEvaluator := &evaluatorFakes.FakeProgramEvaluator{}
	programEvaluator.EvaluateReturns(nil, err)

	printer := &printerFakes.FakeListNotationPrinter{}

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
	tokenizer := &tokenizerFakes.FakeTokenizer{}
	tokenizer.TokenizeReturns(tokens)

	var rootNodes []datamodels.Node
	parser := &parserFakes.FakeParser{}
	parser.ParseReturns(rootNodes, nil)

	partitionedRootNodes := &fakes.FakePartitionedRootNodes{}
	rootNodePartitioner := &internalFakes.FakeRootNodePartitioner{}
	rootNodePartitioner.PartitionRootNodesReturns(partitionedRootNodes)

	defunNode := &fakes.FakeExpressionListNode{}
	defunNodes := []datamodels.ExpressionListNode {defunNode}
	partitionedRootNodes.GetDefunNodesReturns(defunNodes)

	userDefinedFunction := &fakes.FakeDefunFunction{}
	functionGenerator := &userDefinedFakes.FakeFunctionGenerator{}
	functionGenerator.GenerateFunctionReturns(userDefinedFunction, nil)

	functionName := "functionName"
	userDefinedFunction.GetFunctionNameReturns(functionName)

	var evaluatableNodes []datamodels.Node
	partitionedRootNodes.GetEvaluatableNodesReturns(evaluatableNodes)

	expressionListLengthAsserter := &asserterFakes.FakeExpressionListLengthAsserter{}
	expressionListLengthAsserter.AssertLengthIsAsExpectedReturns(nil)

	evaluatedNodes := []datamodels.Node{}
	programEvaluator := &evaluatorFakes.FakeProgramEvaluator{}
	programEvaluator.EvaluateReturns(evaluatedNodes, nil)

	result := "result"
	printer := &printerFakes.FakeListNotationPrinter{}
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