package printer

import "lisp_lambda-go/internal/core/datamodels"

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate -o fakes/fake_list_notation_printer.go . ListNotationPrinter
type ListNotationPrinter interface{
	PrintAllInListNotation(nodes []datamodels.Node) string
	PrintInListNotation(node datamodels.Node) string
}