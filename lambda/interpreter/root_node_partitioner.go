package interpreter

import (
	"lisp_lambda-go/lambda/core/datamodels"
	"lisp_lambda-go/lambda/interpreter/internal"
)

type rootNodePartitioner struct {}

func NewRootNodePartitioner() *rootNodePartitioner {
	return &rootNodePartitioner{}
}

func (r rootNodePartitioner) PartitionRootNodes(nodes []datamodels.Node) *datamodels.PartitionedRootNodes {
	panic("implement me")
}

var _ internal.RootNodePartitioner = &rootNodePartitioner{}


