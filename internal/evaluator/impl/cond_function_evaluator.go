package impl

import (
	"errors"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/evaluator/impl/internal"
)

type condFunctionEvaluator struct {
	topProgramItemUpdater     internal.TopProgramItemUpdater
	condChildStackItemBuilder internal.CondChildStackItemBuilder
}

func NewCondFunctionEvaluator(
	topProgramItemUpdater internal.TopProgramItemUpdater,
	condChildStackItemBuilder internal.CondChildStackItemBuilder,
) *condFunctionEvaluator {
	return &condFunctionEvaluator{
		topProgramItemUpdater,
		condChildStackItemBuilder,
	}
}

func (c condFunctionEvaluator) EvaluateCondProgramItem(
	top datamodels.ProgramItem,
	programStack datamodels.ProgramItemStack,
) error {
	programStack.Push(top)
	if top.GetCurrentParameterIndex() != 0 {
		return errors.New("Error! None of the conditions in the cond function evaluated to true.")
	}
	c.topProgramItemUpdater.UpdateTopProgramItemToNextChild(
		programStack,
	)
	c.condChildStackItemBuilder.BuildCondChildStackItems(
		top,
		programStack,
	)
	return nil
}

var _ internal.CondFunctionEvaluator = &condFunctionEvaluator{}