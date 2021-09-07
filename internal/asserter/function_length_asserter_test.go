package asserter

import (
	"github.com/stretchr/testify/assert"
	"lisp_lambda-go/internal/asserter/internal/fakes"
	"lisp_lambda-go/internal/core/datamodels"
	"testing"
)

func TestNotEquals(t *testing.T) {
	functionName := "functionName"
	expected := 1
	node := datamodels.NewAtomNode("value")

	functionLengthDeterminer := &fakes.FakeFunctionLengthDeterminer{}
	functionLengthDeterminer.DetermineFunctionLengthReturns(expected + 1)

	functionLengthAsserter := NewFunctionLengthAsserter(functionLengthDeterminer)

	actual := functionLengthAsserter.AssertLengthIsAsExpected(
		functionName,
		expected,
		node,
	)

	assert.NotNil(t, actual)
}

func TestEquals(t *testing.T) {
	functionName := "functionName"
	expected := 1
	node := datamodels.NewAtomNode("value")

	functionLengthDeterminer := &fakes.FakeFunctionLengthDeterminer{}
	functionLengthDeterminer.DetermineFunctionLengthReturns(expected)

	functionLengthAsserter := NewFunctionLengthAsserter(functionLengthDeterminer)

	actual := functionLengthAsserter.AssertLengthIsAsExpected(
		functionName,
		expected,
		node,
	)

	assert.Nil(t, actual)
}

func TestNotAtLeastMinimum(t *testing.T) {
	functionName := "functionName"
	expected := 1
	node := datamodels.NewAtomNode("value")

	functionLengthDeterminer := &fakes.FakeFunctionLengthDeterminer{}
	functionLengthDeterminer.DetermineFunctionLengthReturns(expected - 1)

	functionLengthAsserter := NewFunctionLengthAsserter(functionLengthDeterminer)

	actual := functionLengthAsserter.AssertLengthIsAtLeastMinimum(
		functionName,
		expected,
		node,
	)

	assert.NotNil(t, actual)
}

func TestAtLeastMinimum(t *testing.T) {
	functionName := "functionName"
	expected := 1
	node := datamodels.NewAtomNode("value")

	functionLengthDeterminer := &fakes.FakeFunctionLengthDeterminer{}
	functionLengthDeterminer.DetermineFunctionLengthReturns(expected)

	functionLengthAsserter := NewFunctionLengthAsserter(functionLengthDeterminer)

	actual := functionLengthAsserter.AssertLengthIsAtLeastMinimum(
		functionName,
		expected,
		node,
	)

	assert.Nil(t, actual)
}
