package printer

import (
	"lisp_lambda-go/internal"
	"lisp_lambda-go/internal/core/datamodels"
)

type listNotationPrinter struct {}

func NewListNotationPrinter() *listNotationPrinter {
	return &listNotationPrinter{}
}

func (l listNotationPrinter) PrintAllInListNotation(nodes []datamodels.Node) string {
	panic("implement me")
}

func (l listNotationPrinter) PrintInListNotation(node datamodels.Node) string {
	panic("implement me")
}

var _ internal.ListNotationPrinter = &listNotationPrinter{}