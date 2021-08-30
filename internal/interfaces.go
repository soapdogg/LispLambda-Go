package internal

import "lisp_lambda-go/internal/core/datamodels"

type FunctionLengthAsserter interface {
	AssertLengthIsAsExpected(
		functionName string,
		expected int,
		node datamodels.Node,
	) error
	AssertLengthIsAtLeastMinimum(
		functionName string,
		expected int,
		node datamodels.Node,
	) error
}

type ExpressionListLengthAsserter interface{
	AssertLengthIsAsExpected(
		nodes []datamodels.Node,
		userDefinedFunctions map[string] *datamodels.UserDefinedFunction,
	) error
}

type Function interface {
	Evaluate(params *datamodels.NodeStack) datamodels.Node
}

type ListNotationPrinter interface{
	PrintAllInListNotation(nodes []datamodels.Node) string
	PrintInListNotation(node datamodels.Node) string
}

type Parser interface {
	Parse(tokens []string) []datamodels.Node
}

type Tokenizer interface {
	Tokenize(input string) []string
}