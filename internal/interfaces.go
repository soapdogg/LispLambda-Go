package internal

import "lisp_lambda-go/internal/core/datamodels"

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate -o fakes/fake_function_length_asserter.go . FunctionLengthAsserter
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
		userDefinedFunctions map[string] *datamodels.DefunFunction,
	) error
}

//counterfeiter:generate -o fakes/fake_function.go . Function
type Function interface {
	Evaluate(params *datamodels.NodeStack) (datamodels.Node, error)
}

type FunctionGenerator interface {
	GenerateFunction(params *datamodels.ExpressionListNode) (*datamodels.DefunFunction, error)
}

//counterfeiter:generate -o fakes/fake_list_notation_printer.go . ListNotationPrinter
type ListNotationPrinter interface{
	PrintAllInListNotation(nodes []datamodels.Node) string
	PrintInListNotation(node datamodels.Node) string
}

type ProgramEvaluator interface {
	Evaluate(nodes []datamodels.Node, userDefinedFunctions map[string]*datamodels.DefunFunction) ([]datamodels.Node, error)
}

type Parser interface {
	Parse(tokens []string) ([]datamodels.Node, error)
}

type Tokenizer interface {
	Tokenize(input string) []string
}