package asserter

import (
	"errors"
	top "lisp_lambda-go/internal"
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
)


type expressionListLengthAsserter struct{
	functionLengthAsserter top.FunctionLengthAsserter
	functionLengthMap map[string]int
	minimumFunctionLengthMap map[string]int
}

func NewExpressionListLengthAsserter(
	functionLengthAsserter top.FunctionLengthAsserter,
	functionLengthMap map[string]int,
	minimumFunctionLengthMap map[string]int,
) * expressionListLengthAsserter {
	return &expressionListLengthAsserter{
		functionLengthAsserter,
		functionLengthMap,
		minimumFunctionLengthMap,
	}
}

func (e expressionListLengthAsserter) AssertLengthIsAsExpected(
	nodes []datamodels.Node,
	userDefinedFunctions map[string]datamodels.DefunFunction,
) error {
	stack := datamodels.NewExpressionListNodeStack()
	for _, node := range nodes {
		expressionListNode, isExpressionListNode := node.(datamodels.ExpressionListNode)
		if isExpressionListNode {
			stack.Push(expressionListNode)
		}
	}
	for stack.IsNotEmpty() {
		node := stack.Pop()
		address := node.GetChildren()[0]
		addressAtomNode, isAtomNode := address.(datamodels.AtomNode)
		if isAtomNode {
			value := addressAtomNode.GetValue()
			functionLength, isInFunctionLengthMap := e.functionLengthMap[value]
			if isInFunctionLengthMap {
				err := e.functionLengthAsserter.AssertLengthIsAsExpected(
					value,
					functionLength,
					node,
				)
				if err != nil {
					return err
				}
			}
			minimumFunctionLength, isInMinimumFunctionLengthMap := e.minimumFunctionLengthMap[value]
			if isInMinimumFunctionLengthMap {
				err := e.functionLengthAsserter.AssertLengthIsAtLeastMinimum(
					value,
					minimumFunctionLength,
					node,
				)
				if err != nil {
					return err
				}
			}
			userDefinedFunction, isUserDefinedFunction := userDefinedFunctions[value]
			if isUserDefinedFunction {
				err := e.functionLengthAsserter.AssertLengthIsAsExpected(
					value,
					len(userDefinedFunction.GetFormalParameters()) + 1,
					node,
				)
				if err != nil {
					return err
				}
			}
			if value == constants.COND {
				for i := 1; i < len(node.GetChildren()) - 1; i++ {
					condParameter := node.GetChildren()[i]
					expressionListNodeCondParameter, isExpressionList := condParameter.(datamodels.ExpressionListNode)
					if isExpressionList {
						err := e.functionLengthAsserter.AssertLengthIsAsExpected(
							constants.COND,
							constants.TWO,
							expressionListNodeCondParameter,
						)
						if err != nil {
							return err
						}
					} else {
						return errors.New("Error! cond parameter is not a list")
					}
				}
			}

			for _, grandChildNode := range node.GetChildren() {
				expressionListGrandChild, isExpressionList := grandChildNode.(datamodels.ExpressionListNode)
				if isExpressionList {
					stack.Push(expressionListGrandChild)
				}
			}
		}
	}
	return nil
}

var _ top.ExpressionListLengthAsserter = &expressionListLengthAsserter{}
