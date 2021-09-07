package internal

import (
	"lisp_lambda-go/internal/core/datamodels"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate -o fakes/fake_function_length_determiner.go . FunctionLengthDeterminer
type FunctionLengthDeterminer interface {
	DetermineFunctionLength(node datamodels.Node) int
}
