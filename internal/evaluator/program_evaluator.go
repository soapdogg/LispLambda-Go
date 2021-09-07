package evaluator

import (
	top "lisp_lambda-go/internal"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/evaluator/internal"
)

type programEvaluator struct {
	atomRootNodeAsserter internal.AtomRootNodeAsserter
	rootNodeEvaluator internal.RootNodeEvaluator
}

func NewProgramEvaluator(
	atomRootNodeAsserter internal.AtomRootNodeAsserter,
	rootNodeEvaluator internal.RootNodeEvaluator,
) *programEvaluator {
	return &programEvaluator{
		atomRootNodeAsserter,
		rootNodeEvaluator,
	}
}

func (p programEvaluator) Evaluate(
	nodes []datamodels.Node,
	userDefinedFunctions map[string]datamodels.DefunFunction,
) ([]datamodels.Node, error) {
	result := [] datamodels.Node {}
	for _, node := range nodes {
		atomNode, isAtomNode := node.(datamodels.AtomNode)
		if isAtomNode {
			err := p.atomRootNodeAsserter.AssertAtomRootNode(atomNode)
			if err != nil {
				return nil, err
			}
			result = append(result, atomNode)
		} else {
			expressionListNode := node.(datamodels.ExpressionListNode)
			programStack := datamodels.NewProgramItemStack()
			evalStack := datamodels.NewNodeStack()
			evaluated, err := p.rootNodeEvaluator.Evaluate(
				expressionListNode,
				userDefinedFunctions,
				programStack,
				evalStack,
			)
			if err != nil {
				return nil, err
			}
			result = append(result, evaluated)
		}
	}
	return result, nil
}

var _ top.ProgramEvaluator = &programEvaluator{}