package internal

import "lisp_lambda-go/internal/core/datamodels"

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

type ListValueRetriever interface {
	RetrieveListValue(
		node datamodels.Node,
		functionName string,
	) (datamodels.ExpressionListNode, error)
}

type NumericValueRetriever interface {
	RetrieveNumericValue(
		node datamodels.Node,
		functionName string,
		index int,
	) (int, error)
}
