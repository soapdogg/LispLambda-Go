package datamodels

type ExpressionListNode struct {
	children []Node
}

func NewExpressionListNode(children []Node) *ExpressionListNode{
	return &ExpressionListNode{children}
}

var _ Node = ExpressionListNode{}

