package datamodels

type AtomNode struct {
	value string
}

func NewAtomNode(value string) *AtomNode {
	return &AtomNode{value}
}

var _ Node = AtomNode{}
