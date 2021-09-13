package function

import "lisp_lambda-go/internal/core/datamodels"

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate -o fakes/fake_function.go . Function
type Function interface {
	Evaluate(params datamodels.NodeStack) (datamodels.Node, error)
}