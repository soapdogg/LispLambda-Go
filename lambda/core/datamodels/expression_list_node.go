package datamodels

type ExpressionListNode struct {
	children []Node
}

var _ Node = ExpressionListNode{}

