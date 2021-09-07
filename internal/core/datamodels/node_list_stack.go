package datamodels

type NodeListStack interface {
	Push(item []Node)
	Pop() []Node
	IsNotEmpty() bool
}

type nodeListStack struct {
	elements [][]Node
}

func NewNodeListStack() NodeListStack {
	return &nodeListStack{}
}

func (s *nodeListStack) Push(item []Node) {
	s.elements = append(s.elements, item)
}

func (s *nodeListStack) Pop() []Node {
	lastIndex := len(s.elements) - 1
	res := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]
	return res
}

func (s *nodeListStack) Peek() []Node {
	lastIndex := len(s.elements) - 1
	return s.elements[lastIndex]
}

func (s *nodeListStack) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *nodeListStack) IsNotEmpty() bool {
	return len(s.elements) != 0
}

func (s *nodeListStack) Size() int {
	return len(s.elements)
}

var _ NodeListStack = &nodeListStack{}