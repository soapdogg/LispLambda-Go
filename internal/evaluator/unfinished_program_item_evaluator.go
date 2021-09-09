package evaluator

import (
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/evaluator/internal"
)

type unfinishedProgramItemEvaluator struct {
	stackUpdateDeterminer internal.StackUpdateDeterminer
}

func NewUnfinishedProgramItemEvaluator(
	stackUpdateDeterminer internal.StackUpdateDeterminer,
) *unfinishedProgramItemEvaluator {
	return &unfinishedProgramItemEvaluator{
		stackUpdateDeterminer,
	}
}

func (u unfinishedProgramItemEvaluator) EvaluateUnfinishedProgramItem(
	top datamodels.ProgramItem,
	evalStack datamodels.NodeStack,
	programStack datamodels.ProgramItemStack,
) {
	nthChild := top.GetFunctionExpressionNode().GetChildren()[top.GetCurrentParameterIndex()]
	programStack.Push(top)
	u.stackUpdateDeterminer.DetermineHowToUpdateStacks(
		nthChild,
		top.GetVariableMap(),
		evalStack,
		programStack,
	)
}

var _ internal.UnfinishedProgramStackItemEvaluator = &unfinishedProgramItemEvaluator{}