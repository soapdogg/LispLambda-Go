package evaluator

import (
	"lisp_lambda-go/internal/core/datamodels"
	"lisp_lambda-go/internal/evaluator/internal"
)

type topProgramItemCreator struct {}

func NewTopProgramItemCreator() *topProgramItemCreator {
	return &topProgramItemCreator{}
}

func (t topProgramItemCreator) CreateTopProgramItem(
	expressionListNode *datamodels.ExpressionListNode,
	variableMap map[string]datamodels.Node,
	programStack *datamodels.ProgramItemStack,
) {
 	top := datamodels.NewProgramItemFromScratch(
 		expressionListNode,
 		variableMap,
	)
 	programStack.Push(top)
}

var _ internal.TopProgramItemCreator = &topProgramItemCreator{}
