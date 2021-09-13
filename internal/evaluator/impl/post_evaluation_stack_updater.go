package impl

import (
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/evaluator/impl/internal"
)

type postEvaluationStackUpdater struct {
	topProgramItemUpdater internal.TopProgramItemUpdater
}

func NewPostEvaluationStackUpdater(
	topProgramItemUpdater internal.TopProgramItemUpdater,
) *postEvaluationStackUpdater {
	return &postEvaluationStackUpdater{
		topProgramItemUpdater,
	}
}

func (p postEvaluationStackUpdater) UpdateStacksAfterEvaluation(
	evaluatedNode datamodels.Node,
	variableMap map[string]datamodels.Node,
	evalStack datamodels.NodeStack,
	programStack datamodels.ProgramItemStack,
) {
	var nodeToPush datamodels.Node
	evaluatedAtomNode, isAtomNode := evaluatedNode.(datamodels.AtomNode)
	if isAtomNode {
		variableNode, isVariable := variableMap[evaluatedAtomNode.GetValue()]
		if isVariable {
			nodeToPush = variableNode
		} else {
			nodeToPush = evaluatedNode
		}
	} else {
		nodeToPush = evaluatedNode
	}
	evalStack.Push(nodeToPush)
	p.topProgramItemUpdater.UpdateTopProgramItemToNextChild(
		programStack,
	)
}

var _ internal.PostEvaluationStackUpdater = &postEvaluationStackUpdater{}