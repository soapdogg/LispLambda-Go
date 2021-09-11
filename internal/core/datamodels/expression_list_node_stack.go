package datamodels

type ExpressionListNodeStack interface {
	Push(item ExpressionListNode)
	Pop() ExpressionListNode
	IsNotEmpty() bool
}

type expressionListNodeStack struct {
	elements []ExpressionListNode
}

func NewExpressionListNodeStack() *expressionListNodeStack {
	return &expressionListNodeStack{}
}

func (s *expressionListNodeStack) Push(item ExpressionListNode) {
	s.elements = append(s.elements, item)
}

func (s *expressionListNodeStack) Pop() ExpressionListNode {
	lastIndex := len(s.elements) - 1
	res := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]
	return res
}

func (s *expressionListNodeStack) IsNotEmpty() bool {
	return len(s.elements) != 0
}

var _ ExpressionListNodeStack = &expressionListNodeStack{}