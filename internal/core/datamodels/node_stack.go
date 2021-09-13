package datamodels

type NodeStack interface {
	Push(item Node)
	Pop() Node
	IsNotEmpty() bool
}

type nodeStack struct {
	elements []Node
}

func NewNodeStack() NodeStack {
	return &nodeStack{}
}

func (s *nodeStack) Push(item Node) {
	s.elements = append(s.elements, item)
}

func (s *nodeStack) Pop() Node {
	lastIndex := len(s.elements) - 1
	res := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]
	return res
}

func (s *nodeStack) Peek() Node {
	lastIndex := len(s.elements) - 1
	return s.elements[lastIndex]
}

func (s *nodeStack) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *nodeStack) IsNotEmpty() bool {
	return len(s.elements) != 0
}

func (s *nodeStack) Size() int {
	return len(s.elements)
}

var _ NodeStack = &nodeStack{}