package userdefined

import (
	top "lisp_lambda-go/internal"
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/userdefined/internal"
)

type functionGenerator struct {
	functionLengthAsserter top.FunctionLengthAsserter
	functionNameAsserter internal.FunctionNameAsserter
	formalParametersAsserter internal.FormalParametersAsserter
}

func NewFunctionGenerator(
	functionLengthAsserter top.FunctionLengthAsserter,
	functionNameAsserter internal.FunctionNameAsserter,
	formalParameterAsserter internal.FormalParametersAsserter,
) *functionGenerator {
	return &functionGenerator{
		functionLengthAsserter,
		functionNameAsserter,
		formalParameterAsserter,
	}
}

func (f functionGenerator) GenerateFunction(params *datamodels.ExpressionListNode) (*datamodels.DefunFunction, error) {
	err := f.functionLengthAsserter.AssertLengthIsAsExpected(
		constants.DEFUN,
		constants.FOUR,
		params,
	)
	if err != nil {
		return nil, err
	}

	functionNameNode := params.GetChildren()[1].(*datamodels.AtomNode)
	functionName := functionNameNode.GetValue()
	err = f.functionNameAsserter.AssertFunctionNameIsValid(functionName)
	if err != nil {
		return nil, err
	}

	var formalParameters []string
	formalParametersNode := params.GetChildren()[2]
	expressionListNode, isExpressionListNode := formalParametersNode.(*datamodels.ExpressionListNode)
	if isExpressionListNode && len(expressionListNode.GetChildren()) > 0 {
		for i := 0; i < len(expressionListNode.GetChildren()) - 1; i++ {
			child := expressionListNode.GetChildren()[i]
			atomNode := child.(*datamodels.AtomNode)
			formalParameters = append(formalParameters, atomNode.GetValue())
		}
	}
	err = f.formalParametersAsserter.AssertFormalParameters(formalParameters)
	if err != nil {
		return nil, err
	}

	return datamodels.NewDefunFunction(
		functionName,
		formalParameters,
		params.GetChildren()[3],
	), nil
}

var _ top.FunctionGenerator = &functionGenerator{}