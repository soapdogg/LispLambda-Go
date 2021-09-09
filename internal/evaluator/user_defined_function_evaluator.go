package evaluator

import (
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/evaluator/internal"
)

type userDefinedFunctionEvaluator struct {
	stackUpdateDeterminer internal.StackUpdateDeterminer
}

func NewUserDefinedFunctionEvaluator(
	stackUpdateDeterminer internal.StackUpdateDeterminer,
) *userDefinedFunctionEvaluator {
	return &userDefinedFunctionEvaluator{
		stackUpdateDeterminer,
	}
}

func (u userDefinedFunctionEvaluator) EvaluateUserDefinedFunction(
	userDefinedFunction datamodels.DefunFunction,
	variableMap map[string]datamodels.Node,
	functionStack datamodels.NodeStack,
	evalStack datamodels.NodeStack,
	programStack datamodels.ProgramItemStack,
) {
	mapCopy := map[string]datamodels.Node{}
	for k, v := range variableMap {
		mapCopy[k] = v
	}
	for _, formalParameter := range userDefinedFunction.GetFormalParameters() {
		param := functionStack.Pop()
		mapCopy[formalParameter] = param
	}
	u.stackUpdateDeterminer.DetermineHowToUpdateStacks(
		userDefinedFunction.GetBody(),
		mapCopy,
		evalStack,
		programStack,
	)
}

var _ internal.UserDefinedFunctionEvaluator = &userDefinedFunctionEvaluator{}


