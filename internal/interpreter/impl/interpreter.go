package impl

import (
	"lisp_lambda-go/internal/asserter"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/evaluator"
	interpreter2 "lisp_lambda-go/internal/interpreter"
	"lisp_lambda-go/internal/interpreter/impl/internal"
	parser2 "lisp_lambda-go/internal/parser"
	"lisp_lambda-go/internal/printer"
	tokenizer2 "lisp_lambda-go/internal/tokenizer"
	"lisp_lambda-go/internal/userdefined"
)

type interpreter struct {
	tokenizer tokenizer2.Tokenizer
	parser              parser2.Parser
	rootNodePartitioner internal.RootNodePartitioner
	functionGenerator   userdefined.FunctionGenerator
	expressionListLengthAsserter asserter.ExpressionListLengthAsserter
	programEvaluator evaluator.ProgramEvaluator
	printer printer.ListNotationPrinter
}


func NewInterpreter(
	tokenizer tokenizer2.Tokenizer,
	parser parser2.Parser,
	rootNodePartitioner internal.RootNodePartitioner,
	functionGenerator userdefined.FunctionGenerator,
	expressionListLengthAsserter asserter.ExpressionListLengthAsserter,
	programEvaluator evaluator.ProgramEvaluator,
	printer printer.ListNotationPrinter,
)*interpreter {
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


var _ interpreter2.Interpreter = &interpreter{}
