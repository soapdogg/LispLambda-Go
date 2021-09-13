package asserter

import (
	"lisp_lambda-go/internal/core/datamodels"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate -o fakes/fake_expression_list_length_asserter.go . ExpressionListLengthAsserter
type ExpressionListLengthAsserter interface{
	AssertLengthIsAsExpected(
		nodes []datamodels.Node,
		userDefinedFunctions map[string] datamodels.DefunFunction,
	) error
}

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
