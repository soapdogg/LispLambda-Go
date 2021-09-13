package impl

import (
	"lisp_lambda-go/internal/asserter"
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/function"
	"lisp_lambda-go/internal/userdefined"
)

type singleton struct {
	functionGenerator userdefined.FunctionGenerator
}

func NewSingleton(
	functionMap map[string] function.Function,
	functionLengthAsserter asserter.FunctionLengthAsserter,
) *singleton {
	invalidNames := map[string]bool{}
	for key := range functionMap {
		invalidNames[key] = true
	}
	invalidNames[constants.T] = true
	invalidNames[constants.NIL] = true

	functionNameAsserter := NewFunctionNameAsserter(
		invalidNames,
	)
	formalParameterAsserter := NewUserDefinedFormalParametersAsserter(
		invalidNames,
	)

	functionGenerator := NewFunctionGenerator(
		functionLengthAsserter,
		functionNameAsserter,
		formalParameterAsserter,
	)

	return &singleton{
		functionGenerator,
	}
}

func (s *singleton) GetFunctionGenerator() userdefined.FunctionGenerator {
	return s.functionGenerator
}