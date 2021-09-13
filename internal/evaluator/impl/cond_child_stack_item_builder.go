package impl

import (
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/evaluator/impl/internal"
)

type condChildStackItemBuilder struct {
	topProgramItemCreator internal.TopProgramItemCreator
}

func NewCondChildStackItemBuilder(
	topProgramItemCreator internal.TopProgramItemCreator,
) *condChildStackItemBuilder {
	return &condChildStackItemBuilder{
		topProgramItemCreator,
	}
}

func (c condChildStackItemBuilder) BuildCondChildStackItems(
	condProgramItem datamodels.ProgramItem,
	programStack datamodels.ProgramItemStack,
) {
	condsChildren := condProgramItem.GetFunctionExpressionNode().GetChildren()
	for i := len(condsChildren) - 2; i > 0; i-- {
		condChildAtomNode := datamodels.NewAtomNode(constants.COND_CHILD)
		childExpressionList, _ := condsChildren[i].(datamodels.ExpressionListNode)
		grandchildren := childExpressionList.GetChildren()
		condGrandChildren := []datamodels.Node{condChildAtomNode, grandchildren}
		expressionList := datamodels.NewExpressionListNode(condGrandChildren)
		c.topProgramItemCreator.CreateTopProgramItem(
			expressionList,
			condProgramItem.GetVariableMap(),
			programStack,
		)
	}
}

var _ internal.CondChildStackItemBuilder = &condChildStackItemBuilder{}