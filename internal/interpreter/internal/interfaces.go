package internal

import (
	datamodels2 "lisp_lambda-go/internal/core/datamodels"
)

type RootNodePartitioner interface {
	PartitionRootNodes(nodes[]datamodels2.Node) *datamodels2.PartitionedRootNodes
}