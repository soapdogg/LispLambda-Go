package evaluator

import (
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/evaluator/internal"
)

type stackUpdateDeterminer struct {
	topProgramItemCreator internal.TopProgramItemCreator
	postEvaluationStackUpdater internal.PostEvaluationStackUpdater
}

func NewStackUpdateDeterminer(
	topProgramItemCreator internal.TopProgramItemCreator,
	postEvaluationStackUpdater internal.PostEvaluationStackUpdater,
) *stackUpdateDeterminer {
	return &stackUpdateDeterminer{
		topProgramItemCreator,
		postEvaluationStackUpdater,
	}
}

func (s stackUpdateDeterminer) DetermineHowToUpdateStacks(
	node datamodels.Node,
	variableMap map[string]datamodels.Node,
	evalStack datamodels.NodeStack,
	programStack datamodels.ProgramItemStack,
) {
	expressionListNode, isExpressionListNode := node.(datamodels.ExpressionListNode)
	if isExpressionListNode {
		if len(expressionListNode.GetChildren()) > 1 {
			s.topProgramItemCreator.CreateTopProgramItem(
				expressionListNode,
				variableMap,
				programStack,
			)
		} else {
			s.postEvaluationStackUpdater.UpdateStacksAfterEvaluation(
				expressionListNode.GetChildren()[0],
				variableMap,
				evalStack,
				programStack,
			)
		}
	} else {
		s.postEvaluationStackUpdater.UpdateStacksAfterEvaluation(
			node,
			variableMap,
			evalStack,
			programStack,
		)
	}

}

var _ internal.StackUpdateDeterminer = &stackUpdateDeterminer{}