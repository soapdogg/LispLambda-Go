package impl

import (
	"github.com/stretchr/testify/assert"
	"lisp_lambda-go/internal/core/datamodels"
	datamodelsFakes "lisp_lambda-go/internal/core/datamodels/fakes"
	"lisp_lambda-go/internal/evaluator/impl/internal/fakes"
	"testing"
)

func TestEvaluateUserDefinedFunction(t *testing.T) {
	userDefinedFunction := &datamodelsFakes.FakeDefunFunction{}
	variableMap := map[string] datamodels.Node{}
	functionStack := datamodels.NewNodeStack()
	evalStack := datamodels.NewNodeStack()
	programStack := datamodels.NewProgramItemStack()

	formalParameter0 := "formalParameter0"
	userDefinedFunction.GetFormalParametersReturns([]string{formalParameter0})

	paramValue := &datamodelsFakes.FakeNode{}
	functionStack.Push(paramValue)

	userDefinedFunctionBody := &datamodelsFakes.FakeNode{}
	userDefinedFunction.GetBodyReturns(userDefinedFunctionBody)

	stackUpdateDeterminer := &fakes.FakeStackUpdateDeterminer{}

	userDefinedFunctionEvaluator := NewUserDefinedFunctionEvaluator(stackUpdateDeterminer)

	userDefinedFunctionEvaluator.EvaluateUserDefinedFunction(
		userDefinedFunction,
		variableMap,
		functionStack,
		evalStack,
		programStack,
	)

	actualBody, actualMap, actualEvalStack, actualProgramStack := stackUpdateDeterminer.DetermineHowToUpdateStacksArgsForCall(0)

	assert.Equal(t, userDefinedFunctionBody, actualBody)
	assert.Equal(t, paramValue, actualMap[formalParameter0])
	assert.Equal(t, evalStack, actualEvalStack)
	assert.Equal(t, programStack, actualProgramStack)
}
