package internal

import (
	"lisp_lambda-go/internal/core/datamodels"
)

type FunctionLengthDeterminer interface {
	DetermineFunctionLength(node datamodels.Node) int
}
