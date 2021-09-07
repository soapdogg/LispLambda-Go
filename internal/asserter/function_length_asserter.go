package asserter

import (
	"errors"
	top "lisp_lambda-go/internal"
	"lisp_lambda-go/internal/asserter/internal"
	"lisp_lambda-go/internal/core/datamodels"
	"strconv"
)

type functionLengthAsserter struct {
	functionLengthDeterminer internal.FunctionLengthDeterminer
}

func NewFunctionLengthAsserter(
	functionLengthDeterminer internal.FunctionLengthDeterminer,
) *functionLengthAsserter {
	return &functionLengthAsserter{
		functionLengthDeterminer,
	}
}

func (f functionLengthAsserter) AssertLengthIsAsExpected(
	functionName string,
	expected int,
	node datamodels.Node,
) error {
	actual := f.functionLengthDeterminer.DetermineFunctionLength(node)
	if actual != expected {
		return errors.New("Error! Expected length of " + functionName + " list is " + strconv.Itoa(expected)  + "!    Actual: " + strconv.Itoa(actual) + "\n")
	}
	return nil
}

func (f functionLengthAsserter) AssertLengthIsAtLeastMinimum(
	functionName string,
	expected int,
	node datamodels.Node,
) error {
	actual := f.functionLengthDeterminer.DetermineFunctionLength(node)
	if actual < expected {
		return errors.New("Error! Expected length of " + functionName + " list is to be at least " + strconv.Itoa(expected)  + "!    Actual: " + strconv.Itoa(actual) + "\n")
	}
	return nil
}

var _ top.FunctionLengthAsserter = &functionLengthAsserter{}
