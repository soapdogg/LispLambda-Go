package internal

import (
	"lisp_lambda-go/internal/core/datamodels"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate -o fakes/fake_node_parser.go . NodeParser
type NodeParser interface {
	ParseIntoNode(tokens []string) datamodels.Node
}
