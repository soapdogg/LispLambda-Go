package interpreter

import (
	datamodels2 "lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/interpreter/internal"
)

type rootNodePartitioner struct {}

func NewRootNodePartitioner() *rootNodePartitioner {
	return &rootNodePartitioner{}
}

func (r rootNodePartitioner) PartitionRootNodes(nodes []datamodels2.Node) *datamodels2.PartitionedRootNodes {
	panic("implement me")
}

var _ internal.RootNodePartitioner = &rootNodePartitioner{}


