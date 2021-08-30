package datamodels

type AtomNode struct {
	value string
}

func NewAtomNode(value string) *AtomNode {
	return &AtomNode{value}
}

func (a *AtomNode) GetValue() string {
	return a.value
}

var _ Node = AtomNode{}
