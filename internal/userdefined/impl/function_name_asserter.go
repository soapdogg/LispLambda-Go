package impl

import (
	"errors"
	"lisp_lambda-go/internal/userdefined/impl/internal"
	"strconv"
)

type functionNameAsserter struct {
	keywords map[string]bool
}

func NewFunctionNameAsserter(
	keywords map[string]bool,
) *functionNameAsserter {
	return &functionNameAsserter{keywords}
}

func (f functionNameAsserter) AssertFunctionNameIsValid(functionName string) error {
	_, isKeyword := f.keywords[functionName]

	_, err := strconv.Atoi(functionName)
	isNumeric := err == nil

	if isKeyword || isNumeric {
		return errors.New("Error! Invalid function name: " + functionName + "\n")
	}
	return nil
}

var _ internal.FunctionNameAsserter = &functionNameAsserter{}
