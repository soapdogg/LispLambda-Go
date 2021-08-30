package datamodels

type NodeStack struct {
	elements []Node
}

func NewNodeStack() *NodeStack {
	return &NodeStack{}
}

func (s *NodeStack) Push(item Node) {
	s.elements = append(s.elements, item)
}

func (s *NodeStack) Pop() Node {
	lastIndex := len(s.elements) - 1
	res := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]
	return res
}

func (s *NodeStack) Peek() Node {
	lastIndex := len(s.elements) - 1
	return s.elements[lastIndex]
}

func (s *NodeStack) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *NodeStack) IsNotEmpty() bool {
	return len(s.elements) != 0
}

func (s *NodeStack) Size() int {
	return len(s.elements)
}