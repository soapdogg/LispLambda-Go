package internal

import (
	"lisp_lambda-go/internal/core/datamodels"
)

type NodeParser interface {
	ParseIntoNode(tokens []string) datamodels.Node
}
