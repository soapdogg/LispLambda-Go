package impl

import (
	"lisp_lambda-go/internal/asserter"
	"lisp_lambda-go/internal/core/constants"
)

type singleton struct {
	expressionListLengthAsserter asserter.ExpressionListLengthAsserter
	functionLengthAsserter       asserter.FunctionLengthAsserter
}

func NewSingleton() *singleton {
	functionLengthDeterminer := NewFunctionLengthDeterminer()
	functionLengthAsserter := NewFunctionLengthAsserter(functionLengthDeterminer)

	functionLengthMap := map[string]int {
		constants.ATOM: constants.TWO,
		constants.CAR: constants.TWO,
		constants.NULL: constants.TWO,
	}

	minimumFunctionLengthMap := map[string]int {

	}

	expressionListLengthAsserter := NewExpressionListLengthAsserter(
		functionLengthAsserter,
		functionLengthMap,
		minimumFunctionLengthMap,
	)

	return &singleton{
		expressionListLengthAsserter,
		functionLengthAsserter,
	}
}

func (s *singleton) GetExpressionListLengthAsserter() asserter.ExpressionListLengthAsserter {
	return s.expressionListLengthAsserter
}

func (s *singleton) GetFunctionLengthAsserter() asserter.FunctionLengthAsserter {
	return s.functionLengthAsserter
}
