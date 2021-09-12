package evaluator

import (
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/evaluator/internal"
)

type condChildFunctionEvaluator struct {
	stackUpdateDeterminer internal.StackUpdateDeterminer
}

func NewCondChildFunctionEvaluator(
	stackUpdateDeterminer internal.StackUpdateDeterminer,
) *condChildFunctionEvaluator {
	return &condChildFunctionEvaluator{
		stackUpdateDeterminer,
	}
}

func (c condChildFunctionEvaluator) EvaluateCondChildFunction(
	top datamodels.ProgramItem,
	evalStack datamodels.NodeStack,
	programStack datamodels.ProgramItemStack,
) {
	programStack.Push(top)
	condChildCurrentParam := top.GetFunctionExpressionNode().GetChildren()[top.GetCurrentParameterIndex() + 1]
	if top.GetCurrentParameterIndex() == 0 {
		c.stackUpdateDeterminer.DetermineHowToUpdateStacks(
			condChildCurrentParam,
			top.GetVariableMap(),
			evalStack,
			programStack,
		)
	} else {
		programStack.Pop()
		evaluatedCondChild := evalStack.Pop().(datamodels.AtomNode)
		if evaluatedCondChild.GetValue() != constants.NIL {
			for programStack.Peek().GetFunctionName() != constants.COND {
				programStack.Pop()
			}
			programStack.Pop()
			c.stackUpdateDeterminer.DetermineHowToUpdateStacks(
				condChildCurrentParam,
				top.GetVariableMap(),
				evalStack,
				programStack,
			)
		}
	}
}

var _ internal.CondChildFunctionEvaluator = &condChildFunctionEvaluator{}