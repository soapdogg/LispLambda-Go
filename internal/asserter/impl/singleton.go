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
		constants.CDR: constants.TWO,
		constants.CONS: constants.THREE,
		constants.EVEN_P: constants.TWO,
		constants.INT: constants.TWO,
		constants.MINUS_P: constants.TWO,
		constants.NULL: constants.TWO,
		constants.ODD_P: constants.TWO,
		constants.ONE_MINUS: constants.TWO,
		constants.ONE_PLUS: constants.TWO,
		constants.PLUS_P: constants.TWO,
		constants.QUOTE: constants.TWO,
		constants.ZERO_P: constants.TWO,
	}

	minimumFunctionLengthMap := map[string]int {
		constants.EQ: constants.TWO,
		constants.GREATER: constants.TWO,
		constants.GREATER_EQ: constants.TWO,
		constants.LCM: constants.TWO,
		constants.LESS: constants.TWO,
		constants.LESS_EQ: constants.TWO,
		constants.MAX: constants.TWO,
		constants.MIN: constants.TWO,
		constants.MINUS: constants.TWO,
		constants.NOT_EQ: constants.TWO,
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
