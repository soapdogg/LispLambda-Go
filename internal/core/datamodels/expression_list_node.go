package datamodels

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate -o fakes/fake_expression_list_node.go . ExpressionListNode
type ExpressionListNode interface {
	GetChildren() []Node
}

type expressionListNode struct {
	children []Node
}

func NewExpressionListNode(children []Node) *expressionListNode {
	return &expressionListNode{children}
}

func (x *expressionListNode) GetChildren() []Node {
	return x.children
}

var _ Node = &expressionListNode{}
var _ ExpressionListNode = &expressionListNode{}

