package interpreter

import (
	top "lisp_lambda-go/internal"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/interpreter/internal"
)

type interpreter struct {
	tokenizer top.Tokenizer
	parser top.Parser
	rootNodePartitioner internal.RootNodePartitioner
	functionGenerator top.FunctionGenerator
	expressionListLengthAsserter top.ExpressionListLengthAsserter
	programEvaluator top.ProgramEvaluator
	printer top.ListNotationPrinter
}


func NewInterpreter(
	tokenizer top.Tokenizer,
	parser top.Parser,
	rootNodePartitioner internal.RootNodePartitioner,
	functionGenerator top.FunctionGenerator,
	expressionListLengthAsserter top.ExpressionListLengthAsserter,
	programEvaluator top.ProgramEvaluator,
	printer top.ListNotationPrinter,
)*interpreter{
	return &interpreter{
		tokenizer,
		parser,
		rootNodePartitioner,
		functionGenerator,
		expressionListLengthAsserter,
		programEvaluator,
		printer,
	}
}

func (i interpreter) Interpret(input string) (string, error) {
	tokens := i.tokenizer.Tokenize(input)
	rootNodes, err := i.parser.Parse(tokens)
	if err != nil {
		return "", err
	}
	partitionedRootNodes := i.rootNodePartitioner.PartitionRootNodes(rootNodes)
	userDefinedFunctions := map[string]datamodels.DefunFunction{}
	for _, rootNode := range partitionedRootNodes.GetDefunNodes() {
		userDefinedFunction, err := i.functionGenerator.GenerateFunction(rootNode)
		if err != nil {
			return "", err
		}
		userDefinedFunctions[userDefinedFunction.GetFunctionName()] = userDefinedFunction
	}

	err = i.expressionListLengthAsserter.AssertLengthIsAsExpected(
		partitionedRootNodes.GetEvaluatableNodes(),
		userDefinedFunctions,
	)
	if err != nil {
		return "", err
	}

	evaluatedNodes, err := i.programEvaluator.Evaluate(
		partitionedRootNodes.GetEvaluatableNodes(),
		userDefinedFunctions,
	)
	if err != nil {
		return "", err
	}
	return i.printer.PrintAllInListNotation(evaluatedNodes), nil
}


var _ top.Interpreter = &interpreter{}
