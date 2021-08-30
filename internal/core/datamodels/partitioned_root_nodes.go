package datamodels

type PartitionedRootNodes struct {
	defunNodes []*ExpressionListNode
	evaluatableNodes []Node
}

func NewPartitionedRootNodes(
	defunNodes []*ExpressionListNode,
	evaluatableNodes []Node,
) *PartitionedRootNodes {
	return &PartitionedRootNodes{
		defunNodes,
		evaluatableNodes,
	}
}

func (p *PartitionedRootNodes) GetDefunNodes() []*ExpressionListNode {
	return p.defunNodes
}

func (p *PartitionedRootNodes) GetEvaluatableNodes() []Node {
	return p.evaluatableNodes
}
