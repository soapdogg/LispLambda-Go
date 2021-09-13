package impl

import (
	"errors"
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/evaluator/impl/internal"
	"strconv"
)

type atomRootNodeAsserter struct {}

func NewAtomRootNodeAsserter()*atomRootNodeAsserter {
	return &atomRootNodeAsserter{}
}

func (a atomRootNodeAsserter) AssertAtomRootNode(atomNode datamodels.AtomNode) error {
	value := atomNode.GetValue()
	_, convErr :=  strconv.Atoi(value)
	isNotNumeric := convErr != nil
	isNotT := value != constants.T
	isNotNil := value != constants.NIL
	if isNotNumeric && isNotT && isNotNil {
		return errors.New("Error! " + value + " is not a valid atomic value")
	}
	return nil
}

var _ internal.AtomRootNodeAsserter = &atomRootNodeAsserter{}
