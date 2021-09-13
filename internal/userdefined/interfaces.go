package userdefined

import "lisp_lambda-go/internal/core/datamodels"

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate -o fakes/fake_function_generator.go . FunctionGenerator
type FunctionGenerator interface {
	GenerateFunction(params datamodels.ExpressionListNode) (datamodels.DefunFunction, error)
}