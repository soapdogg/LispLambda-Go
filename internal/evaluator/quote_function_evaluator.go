package evaluator

import (
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/evaluator/internal"
)

type quoteFunctionEvaluator struct {
	postEvaluationStackUpdater internal.PostEvaluationStackUpdater
}

func NewQuoteFunctionEvaluator(
	postEvaluationStackUpdater internal.PostEvaluationStackUpdater,
) *quoteFunctionEvaluator {
	return &quoteFunctionEvaluator{
		postEvaluationStackUpdater,
	}
}

func (q quoteFunctionEvaluator) EvaluateQuoteFunction(
	top *datamodels.ProgramItem,
	evalStack *datamodels.NodeStack,
	programStack *datamodels.ProgramItemStack,
) {
	quoteExprNode := top.GetFunctionExpressionNode()
	quoted := quoteExprNode.GetChildren()[1]
	q.postEvaluationStackUpdater.UpdateStacksAfterEvaluation(
		quoted,
		top.GetVariableMap(),
		evalStack,
		programStack,
	)
}

var _ internal.QuoteFunctionEvaluator = &quoteFunctionEvaluator{}