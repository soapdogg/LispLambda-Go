package impl

import (
	"github.com/stretchr/testify/assert"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/core/datamodels/fakes"
	fakes2 "lisp_lambda-go/internal/evaluator/impl/internal/fakes"
	"testing"
)

func TestCurrentParameterIndexIsNotZero(t *testing.T) {
	top := &fakes.FakeProgramItem{}
	programStack := datamodels.NewProgramItemStack()

	top.GetCurrentParameterIndexReturns(1)

	topProgramItemUpdater := &fakes2.FakeTopProgramItemUpdater{}
	condChildStackItemBuilder := &fakes2.FakeCondChildStackItemBuilder{}

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

	topProgramItemUpdater := &fakes2.FakeTopProgramItemUpdater{}
	condChildStackItemBuilder := &fakes2.FakeCondChildStackItemBuilder{}

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