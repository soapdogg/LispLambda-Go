package impl

import (
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/function"
)

type singleton struct {
	functionMap map[string]function.Function
}

func NewSingleton() *singleton {
	atomFunction := NewAtomFunction()
	nullFunction := NewNullFunction()

	functionMap := map[string]function.Function {
		constants.ATOM: atomFunction,
		constants.NULL: nullFunction,
	}

	return &singleton{functionMap}
}

func (s *singleton) GetFunctionMap() map[string]function.Function {
	return s.functionMap
}