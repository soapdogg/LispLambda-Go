package datamodels

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate -o fakes/fake_partitioned_root_nodes.go . PartitionedRootNodes
type PartitionedRootNodes interface {
	GetDefunNodes() []ExpressionListNode
	GetEvaluatableNodes() []Node
}

type partitionedRootNodes struct {
	defunNodes []ExpressionListNode
	evaluatableNodes []Node
}

func NewPartitionedRootNodes(
	defunNodes []ExpressionListNode,
	evaluatableNodes []Node,
) *partitionedRootNodes {
	return &partitionedRootNodes{
		defunNodes,
		evaluatableNodes,
	}
}

func (p *partitionedRootNodes) GetDefunNodes() []ExpressionListNode {
	return p.defunNodes
}

func (p *partitionedRootNodes) GetEvaluatableNodes() []Node {
	return p.evaluatableNodes
}

var _ PartitionedRootNodes = &partitionedRootNodes{}