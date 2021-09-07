package datamodels

type ProgramItemStack struct {
	elements []*ProgramItem
}

func NewProgramItemStack() *ProgramItemStack {
	return &ProgramItemStack{}
}

func (s *ProgramItemStack) Push(item *ProgramItem) {
	s.elements = append(s.elements, item)
}

func (s *ProgramItemStack) Pop() *ProgramItem {
	lastIndex := len(s.elements) - 1
	res := s.elements[lastIndex]
	s.elements = s.elements[:lastIndex]
	return res
}

func (s *ProgramItemStack) Peek() *ProgramItem {
	lastIndex := len(s.elements) - 1
	return s.elements[lastIndex]
}

func (s *ProgramItemStack) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *ProgramItemStack) IsNotEmpty() bool {
	return len(s.elements) != 0
}

func (s *ProgramItemStack) Size() int {
	return len(s.elements)
}