package internal

import "lisp_lambda-go/internal/core/datamodels"

type AtomRootNodeAsserter interface {
	AssertAtomRootNode(atomNode *datamodels.AtomNode) error
}

type BuiltInFunctionEvaluator interface {
	EvaluateBuiltInFunction(
		functionName string,
		functionStack *datamodels.NodeStack,
		top *datamodels.ProgramItem,
		evalStack *datamodels.NodeStack,
		programStack *datamodels.ProgramItemStack,
	) error
}

type CondFunctionEvaluator interface {
	EvaluateCondProgramItem(
		top *datamodels.ProgramItem,
		programStack *datamodels.ProgramItemStack,
	) error
}

type CondChildFunctionEvaluator interface {
	EvaluateCondChildFunction(
		top *datamodels.ProgramItem,
		evalStack *datamodels.NodeStack,
		programStack *datamodels.ProgramItemStack,
	) error
}

type QuoteFunctionEvaluator interface{
	EvaluateQuoteFunction(
		top *datamodels.ProgramItem,
		evalStack *datamodels.NodeStack,
		programStack *datamodels.ProgramItemStack,
	)
}

type RootNodeEvaluator interface {
	Evaluate(
		rootNode *datamodels.ExpressionListNode,
		userDefinedFunctions map[string]*datamodels.DefunFunction,
		programStack *datamodels.ProgramItemStack,
		evalStack *datamodels.NodeStack,
	) (datamodels.Node, error)
}

type TopProgramItemCreator interface {
	CreateTopProgramItem(
		expressionListNode *datamodels.ExpressionListNode,
		variableMap map[string]datamodels.Node,
		programStack *datamodels.ProgramItemStack,
	)
}

type TopProgramItemUpdater interface {
	UpdateTopProgramItemToNextChild(
		programStack *datamodels.ProgramItemStack,
	)
}