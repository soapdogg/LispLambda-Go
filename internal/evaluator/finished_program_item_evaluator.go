package evaluator

import (
	"errors"
	top "lisp_lambda-go/internal"
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/evaluator/internal"
)

type finishedProgramItemEvaluator struct {
	postEvaluationStackUpdater internal.PostEvaluationStackUpdater
	functionMap map[string]top.Function
	builtInFunctionEvaluator internal.BuiltInFunctionEvaluator
	userDefinedFunctionEvaluator internal.UserDefinedFunctionEvaluator	
}


func NewFinishedProgramItemEvaluator(
	postEvaluationStackUpdater internal.PostEvaluationStackUpdater,
	functionMap map[string]top.Function,
	builtInFunctionEvaluator internal.BuiltInFunctionEvaluator,
	userDefinedFunctionEvaluator internal.UserDefinedFunctionEvaluator,
) *finishedProgramItemEvaluator {
	return &finishedProgramItemEvaluator{
		postEvaluationStackUpdater,
		functionMap,
		builtInFunctionEvaluator,
		userDefinedFunctionEvaluator,
	}
}

func (f finishedProgramItemEvaluator) EvaluateFinishedProgramItem(
	top datamodels.ProgramItem,
	userDefinedFunctions map[string]datamodels.DefunFunction,
	evalStack datamodels.NodeStack,
	programStack datamodels.ProgramItemStack,
) error {
	
	functionStack := datamodels.NewNodeStack()
	for i := 0; i < len(top.GetFunctionExpressionNode().GetChildren()) - 1; i++ {
		functionStack.Push(evalStack.Pop())
	}
	
	functionNameNode := functionStack.Pop()
	functionName := top.GetFunctionName()
	
	_, isInFunctionMap := f.functionMap[functionName]
	userDefinedFunction, isUserDefinedFunction := userDefinedFunctions[functionName]
	if functionName == constants.NIL {
		f.postEvaluationStackUpdater.UpdateStacksAfterEvaluation(
			functionNameNode,
			top.GetVariableMap(),
			evalStack,
			programStack,
		)
	} else if isInFunctionMap {
		err := f.builtInFunctionEvaluator.EvaluateBuiltInFunction(
			functionName,
			functionStack,
			top,
			evalStack,
			programStack,
		)
		if err != nil {
			return err
		}
	} else if isUserDefinedFunction {
		f.userDefinedFunctionEvaluator.EvaluateUserDefinedFunction(
			userDefinedFunction,
			top.GetVariableMap(),
			functionStack,
			evalStack,
			programStack,
		)
	} else {
		return errors.New("Error! Invalid CAR value: " + functionName)
	}
	return nil
}

var _ internal.FinishedProgramItemEvaluator = &finishedProgramItemEvaluator{}