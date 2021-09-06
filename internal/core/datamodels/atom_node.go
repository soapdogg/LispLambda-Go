package datamodels

import (
	"lisp_lambda-go/internal/core/constants"
	"strconv"
)

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

func NewAtomNodeFromInt(i int) *AtomNode {
	return &AtomNode{
		strconv.Itoa(i),
	}
}

func (a *AtomNode) GetValue() string {
	return a.value
}

var _ Node = AtomNode{}
