package evaluator

import (
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/evaluator/internal"
)

type rootNodeEvaluator struct {
	topProgramItemCreator internal.TopProgramItemCreator
	condFunctionEvaluator internal.CondFunctionEvaluator
	condChildFunctionEvaluator internal.CondChildFunctionEvaluator
	quoteFunctionEvaluator internal.QuoteFunctionEvaluator
	unfinishedProgramItemEvaluator internal.UnfinishedProgramStackItemEvaluator
	finishedProgramItemEvaluator internal.FinishedProgramItemEvaluator
}

func NewRootNodeEvaluator(
	topProgramItemCreator internal.TopProgramItemCreator,
	condFunctionEvaluator internal.CondFunctionEvaluator,
	condChildFunctionEvaluator internal.CondChildFunctionEvaluator,
	quoteFunctionEvaluator internal.QuoteFunctionEvaluator,
	unfinishedProgramItemEvaluator internal.UnfinishedProgramStackItemEvaluator,
	finishedProgramItemEvaluator internal.FinishedProgramItemEvaluator,
) *rootNodeEvaluator {
	return &rootNodeEvaluator{
		topProgramItemCreator,
		condFunctionEvaluator,
		condChildFunctionEvaluator,
		quoteFunctionEvaluator,
		unfinishedProgramItemEvaluator,
		finishedProgramItemEvaluator,
	}
}

func (r rootNodeEvaluator) Evaluate(
	rootNode datamodels.ExpressionListNode,
	userDefinedFunctions map[string]datamodels.DefunFunction,
	programStack datamodels.ProgramItemStack,
	evalStack datamodels.NodeStack,
) (
	datamodels.Node,
	error,
) {
	r.topProgramItemCreator.CreateTopProgramItem(
		rootNode,
		map[string]datamodels.Node{},
		programStack,
	)

	for programStack.IsNotEmpty() {
		head := programStack.Pop()

		if head.GetFunctionName() == constants.COND {
			err := r.condFunctionEvaluator.EvaluateCondProgramItem(
				head,
				programStack,
			)
			if err != nil {
				return nil, err
			}
		} else if head.GetFunctionName() == constants.COND_CHILD {
			r.condChildFunctionEvaluator.EvaluateCondChildFunction(
				head,
				evalStack,
				programStack,
			)
		} else if head.GetFunctionName() == constants.QUOTE {
			r.quoteFunctionEvaluator.EvaluateQuoteFunction(
				head,
				evalStack,
				programStack,
			)
		} else if head.GetCurrentParameterIndex() < len(head.GetFunctionExpressionNode().GetChildren()) - 1 {
			r.unfinishedProgramItemEvaluator.EvaluateUnfinishedProgramItem(
				head,
				evalStack,
				programStack,
			)
		} else {
			err := r.finishedProgramItemEvaluator.EvaluateFinishedProgramItem(
				head,
				userDefinedFunctions,
				evalStack,
				programStack,
			)
			if err != nil {
				return nil, err
			}
		}
	}

	return evalStack.Pop(), nil
}

var _ internal.RootNodeEvaluator = &rootNodeEvaluator{}
