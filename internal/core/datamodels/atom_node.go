package datamodels

import (
	"lisp_lambda-go/internal/core/constants"
	"strconv"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate -o fakes/fake_atom_node.go . AtomNode
type AtomNode interface {
	GetValue() string
}

type atomNode struct {
	value string
}

func NewAtomNode(value string) *atomNode {
	return &atomNode{value}
}

func NewAtomNodeFromBool(v bool) *atomNode {
	var value string
	if v {
		value = constants.T
	} else {
		value = constants.NIL
	}
	return &atomNode{value}
}

func NewAtomNodeFromInt(i int) *atomNode {
	return &atomNode{
		strconv.Itoa(i),
	}
}

func (a *atomNode) GetValue() string {
	return a.value
}

var _ Node = &atomNode{}
var _ AtomNode = &atomNode{}
