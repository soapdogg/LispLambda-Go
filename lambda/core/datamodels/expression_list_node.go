package datamodels

type ExpressionListNode struct {
	children []Node
}

func NewExpressionListNode(children []Node) *ExpressionListNode{
	return &ExpressionListNode{children}
}

func (x *ExpressionListNode) GetChildren() []Node{
	return x.children
}

var _ Node = ExpressionListNode{}

