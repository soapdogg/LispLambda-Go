package datamodels

type ProgramItemStack interface {
	Push(item ProgramItem)
	Pop() ProgramItem
	Peek() ProgramItem
	IsNotEmpty() bool
	Size() int
}

type programItemStack struct {
	elements []ProgramItem
}

func NewProgramItemStack() ProgramItemStack {
	return &programItemStack{}
}

func (s *programItemStack) Push(item ProgramItem) {
	s.elements = append(s.elements, item)
}

func (s *programItemStack) Pop() ProgramItem {
	lastIndex := len(s.elements) - 1
	res := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]
	return res
}

func (s *programItemStack) Peek() ProgramItem {
	lastIndex := len(s.elements) - 1
	return s.elements[lastIndex]
}

func (s *programItemStack) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *programItemStack) IsNotEmpty() bool {
	return len(s.elements) != 0
}

func (s *programItemStack) Size() int {
	return len(s.elements)
}

var _ ProgramItemStack = &programItemStack{}