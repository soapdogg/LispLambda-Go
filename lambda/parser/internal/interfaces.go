package internal

import "lisp_lambda-go/lambda/core/datamodels"

type NodeParser interface {
	ParseIntoNode(tokens []string) datamodels.Node
}
