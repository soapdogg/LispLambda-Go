package internal

import (
	"lisp_lambda-go/internal/core/datamodels"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate -o fakes/fake_parser.go . Parser
type Parser interface {
	Parse(tokens []string) ([]datamodels.Node, error)
}
