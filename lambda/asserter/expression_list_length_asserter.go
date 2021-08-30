package asserter

import "lisp_lambda-go/lambda/core/datamodels"


type expressionListLengthAsserter struct{}

func (e expressionListLengthAsserter) AssertLengthIsAsExpected(nodes []datamodels.Node, userDefinedFunctions map[string]*datamodels.UserDefinedFunction) error {
	panic("implement me")
}

var _ ExpressionListLengthAsserter = &expressionListLengthAsserter{}
