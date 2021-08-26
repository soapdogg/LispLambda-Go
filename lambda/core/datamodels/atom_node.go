package datamodels

type AtomNode struct {
	value string
}

var _ Node = AtomNode{}
