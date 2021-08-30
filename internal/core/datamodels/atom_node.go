package datamodels

import "lisp_lambda-go/internal/core/constants"

type AtomNode struct {
	value string
}

func NewAtomNode(value string) *AtomNode {
	return &AtomNode{value}
}

func NewAtomNodeFromBool(v bool) *AtomNode {
	var value string
	if v {
		value = constants.T
	} else {
		value = constants.NIL
	}
	return &AtomNode{value}
}

func (a *AtomNode) GetValue() string {
	return a.value
}

var _ Node = AtomNode{}
