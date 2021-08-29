package datamodels

type NodeListStack struct {
	elements [][]Node
}

func NewNodeListStack() *NodeListStack {
	return &NodeListStack{}
}

func (s *NodeListStack) Push(item []Node) {
	s.elements = append(s.elements, item)
}

func (s *NodeListStack) Pop() []Node {
	lastIndex := len(s.elements) - 1
	res := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]
	return res
}

func (s *NodeListStack) Peek() []Node {
	lastIndex := len(s.elements) - 1
	return s.elements[lastIndex]
}

func (s *NodeListStack) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *NodeListStack) IsNotEmpty() bool {
	return len(s.elements) != 0
}

func (s *NodeListStack) Size() int {
	return len(s.elements)
}