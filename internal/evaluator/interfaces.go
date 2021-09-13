package evaluator

import "lisp_lambda-go/internal/core/datamodels"

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate -o fakes/fake_program_evaluator.go . ProgramEvaluator
type ProgramEvaluator interface {
	Evaluate(nodes []datamodels.Node, userDefinedFunctions map[string]datamodels.DefunFunction) ([]datamodels.Node, error)
}