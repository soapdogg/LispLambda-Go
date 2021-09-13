package impl

import (
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/function"
	"lisp_lambda-go/internal/printer"
)

type singleton struct {
	functionMap map[string]function.Function
}

func NewSingleton(
	listNotationPrinter printer.ListNotationPrinter,
) *singleton {
	listValueRetriever := NewListValueRetriever()
	numericValueRetriever := NewNumericValueRetriever(listNotationPrinter)
	gcdCalculator := NewGcdCalculator()

	atomFunction := NewAtomFunction()
	carFunction := NewCarFunction(listValueRetriever)
	cdrFunction := NewCdrFunction(listValueRetriever)
	consFunction := NewConsFunction()
	eqFunction := NewEqFunction(listNotationPrinter)
	evenPFunction := NewEvenPFunction(numericValueRetriever)
	gcdFunction := NewGCDFunction(
		numericValueRetriever,
		gcdCalculator,
	)
	greaterFunction := NewGreaterFunction(numericValueRetriever)
	greaterEqFunction := NewGreaterEqFunction(numericValueRetriever)
	intFunction := NewIntFunction()
	lcmFunction := NewLCMFunction(
		numericValueRetriever,
		gcdCalculator,
	)
	lessFunction := NewLessFunction(numericValueRetriever)
	lessEqFunction := NewLessEqFunction(numericValueRetriever)
	maxFunction := NewMaxFunction(numericValueRetriever)
	minFunction := NewMinFunction(numericValueRetriever)
	minusFunction := NewMinusFunction(numericValueRetriever)
	minusPFunction := NewMinusPFunction(numericValueRetriever)
	notEqFunction := NewNotEqFunction()
	nullFunction := NewNullFunction()
	oddPFunction := NewOddPFunction(numericValueRetriever)
	oneMinusFunction := NewOneMinusFunction(numericValueRetriever)
	onePlusFunction := NewOnePlusFunction(numericValueRetriever)
	plusFunction := NewPlusFunction(numericValueRetriever)
	plusPFunction := NewPlusPFunction(numericValueRetriever)
	timesFunction := NewTimesFunction(numericValueRetriever)
	zeroPFunction := NewZeroPFunction(numericValueRetriever)

	functionMap := map[string]function.Function {
		constants.ATOM: atomFunction,
		constants.CAR: carFunction,
		constants.CDR: cdrFunction,
		constants.CONS: consFunction,
		constants.EQ: eqFunction,
		constants.EVEN_P: evenPFunction,
		constants.GCD: gcdFunction,
		constants.GREATER: greaterFunction,
		constants.GREATER_EQ: greaterEqFunction,
		constants.INT: intFunction,
		constants.LCM: lcmFunction,
		constants.LESS: lessFunction,
		constants.LESS_EQ: lessEqFunction,
		constants.MAX: maxFunction,
		constants.MIN: minFunction,
		constants.NOT_EQ: notEqFunction,
		constants.NULL: nullFunction,
		constants.MINUS: minusFunction,
		constants.MINUS_P: minusPFunction,
		constants.ODD_P: oddPFunction,
		constants.ONE_MINUS: oneMinusFunction,
		constants.ONE_PLUS: onePlusFunction,
		constants.PLUS: plusFunction,
		constants.PLUS_P: plusPFunction,
		constants.TIMES: timesFunction,
		constants.ZERO_P: zeroPFunction,
	}

	return &singleton{functionMap}
}

func (s *singleton) GetFunctionMap() map[string]function.Function {
	return s.functionMap
}