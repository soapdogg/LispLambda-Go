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

	atomFunction := NewAtomFunction()
	carFunction := NewCarFunction(listValueRetriever)
	cdrFunction := NewCdrFunction(listValueRetriever)
	eqFunction := NewEqFunction(listNotationPrinter)
	nullFunction := NewNullFunction()

	functionMap := map[string]function.Function {
		constants.ATOM: atomFunction,
		constants.CAR: carFunction,
		constants.CDR: cdrFunction,
		constants.EQ: eqFunction,
		constants.NULL: nullFunction,
	}

	return &singleton{functionMap}
}

func (s *singleton) GetFunctionMap() map[string]function.Function {
	return s.functionMap
}