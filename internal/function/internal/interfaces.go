package internal

import "lisp_lambda-go/internal/core/datamodels"

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
