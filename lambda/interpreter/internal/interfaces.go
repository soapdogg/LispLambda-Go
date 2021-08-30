package internal

import "lisp_lambda-go/lambda/core/datamodels"

type RootNodePartitioner interface {
	PartitionRootNodes(nodes[]datamodels.Node) *datamodels.PartitionedRootNodes
}