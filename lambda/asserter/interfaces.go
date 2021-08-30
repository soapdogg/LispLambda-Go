package asserter

import "lisp_lambda-go/lambda/core/datamodels"

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