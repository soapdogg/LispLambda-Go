package parser

import "lisp_lambda-go/lambda/core/datamodels"

type Parser interface {
	Parse(tokens []string) []datamodels.Node
}