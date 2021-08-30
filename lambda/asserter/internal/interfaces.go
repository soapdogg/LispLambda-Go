package internal

import "lisp_lambda-go/lambda/core/datamodels"

type FunctionLengthDeterminer interface {
	DetermineFunctionLength(node datamodels.Node) int
}
