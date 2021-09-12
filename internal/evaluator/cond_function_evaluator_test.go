package evaluator

import (
	"github.com/stretchr/testify/assert"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/core/datamodels/fakes"
	internalFakes "lisp_lambda-go/internal/evaluator/internal/fakes"
	"testing"
)

func TestCurrentParameterIndexIsNotZero(t *testing.T) {
	top := &fakes.FakeProgramItem{}
	programStack := datamodels.NewProgramItemStack()

	top.GetCurrentParameterIndexReturns(1)

	topProgramItemUpdater := &internalFakes.FakeTopProgramItemUpdater{}
	condChildStackItemBuilder := &internalFakes.FakeCondChildStackItemBuilder{}

	condFunctionEvaluator := NewCondFunctionEvaluator(
		topProgramItemUpdater,
		condChildStackItemBuilder,
	)

	actual := condFunctionEvaluator.EvaluateCondProgramItem(
		top,
		programStack,
	)

	assert.NotNil(t, actual)
}

func TestCurrentParameterIndexIsZero(t *testing.T) {
	top := &fakes.FakeProgramItem{}
	programStack := datamodels.NewProgramItemStack()

	top.GetCurrentParameterIndexReturns(0)

	topProgramItemUpdater := &internalFakes.FakeTopProgramItemUpdater{}
	condChildStackItemBuilder := &internalFakes.FakeCondChildStackItemBuilder{}

	condFunctionEvaluator := NewCondFunctionEvaluator(
		topProgramItemUpdater,
		condChildStackItemBuilder,
	)

	actual := condFunctionEvaluator.EvaluateCondProgramItem(
		top,
		programStack,
	)

	assert.Nil(t, actual)
}