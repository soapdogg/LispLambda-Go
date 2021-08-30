package printer

import "lisp_lambda-go/lambda/core/datamodels"

type ListNotationPrinter interface{
	PrintAllInListNotation(nodes []datamodels.Node) string
	PrintInListNotation(node datamodels.Node) string
}

