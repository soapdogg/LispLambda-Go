package internal

import "lisp_lambda-go/internal/core/datamodels"

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

type GCDCalculator interface {
	CalculateGCD(i int, j int) int
}

//counterfeiter:generate -o fakes/fake_list_value_retriever.go . ListValueRetriever
type ListValueRetriever interface {
	RetrieveListValue(
		node datamodels.Node,
		functionName string,
	) (datamodels.ExpressionListNode, error)
}

//counterfeiter:generate -o fakes/fake_numeric_value_retriever.go . NumericValueRetriever
type NumericValueRetriever interface {
	RetrieveNumericValue(
		node datamodels.Node,
		functionName string,
		index int,
	) (int, error)
}
