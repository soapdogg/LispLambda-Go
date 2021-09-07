package evaluator

import (
	top "lisp_lambda-go/internal"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/evaluator/internal"
)

type builtInFunctionEvaluator struct {
	functionMap map[string] top.Function
	postEvaluationStackUpdater internal.PostEvaluationStackUpdater
}

func NewBuiltInFunctionEvaluator(
	functionMap map[string] top.Function,
	postEvaluationStackUpdater internal.PostEvaluationStackUpdater,
) *builtInFunctionEvaluator {
	return &builtInFunctionEvaluator{
		functionMap,
		postEvaluationStackUpdater,
	}
}

func (b builtInFunctionEvaluator) EvaluateBuiltInFunction(
	functionName string,
	functionStack datamodels.NodeStack,
	top datamodels.ProgramItem,
	evalStack datamodels.NodeStack,
	programStack datamodels.ProgramItemStack,
) error {
	function := b.functionMap[functionName]
	evaluatedFunctionResult, err := function.Evaluate(functionStack)
	if err != nil {
		return err
	}
	b.postEvaluationStackUpdater.UpdateStacksAfterEvaluation(
		evaluatedFunctionResult,
		top.GetVariableMap(),
		evalStack,
		programStack,
	)
	return nil
}


var _ internal.BuiltInFunctionEvaluator = &builtInFunctionEvaluator{}
