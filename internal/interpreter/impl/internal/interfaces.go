package internal

import "lisp_lambda-go/internal/core/datamodels"

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate -o fakes/fake_root_node_partitioner.go . RootNodePartitioner
type RootNodePartitioner interface {
	PartitionRootNodes(nodes[]datamodels.Node) datamodels.PartitionedRootNodes
}