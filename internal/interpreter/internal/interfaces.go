package internal

import "lisp_lambda-go/internal/core/datamodels"


type RootNodePartitioner interface {
	PartitionRootNodes(nodes[]datamodels.Node) *datamodels.PartitionedRootNodes
}