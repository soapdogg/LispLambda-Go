package asserter

import (
	top "lisp_lambda-go/internal"
	"lisp_lambda-go/internal/core/datamodels"
)


type expressionListLengthAsserter struct{}

func (e expressionListLengthAsserter) AssertLengthIsAsExpected(
	nodes []datamodels.Node,
	userDefinedFunctions map[string]*datamodels.UserDefinedFunction,
) error {
	panic("implement me")
}

var _ top.ExpressionListLengthAsserter = &expressionListLengthAsserter{}
