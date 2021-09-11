package internal

import "lisp_lambda-go/internal/core/datamodels"

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

type AtomRootNodeAsserter interface {
	AssertAtomRootNode(atomNode datamodels.AtomNode) error
}

//counterfeiter:generate -o fakes/fake_built_in_function_evaluator.go . BuiltInFunctionEvaluator
type BuiltInFunctionEvaluator interface {
	EvaluateBuiltInFunction(
		functionName string,
		functionStack datamodels.NodeStack,
		top datamodels.ProgramItem,
		evalStack datamodels.NodeStack,
		programStack datamodels.ProgramItemStack,
	) error
}

type CondChildFunctionEvaluator interface {
	EvaluateCondChildFunction(
		top datamodels.ProgramItem,
		evalStack datamodels.NodeStack,
		programStack datamodels.ProgramItemStack,
	) error
}

type CondChildStackItemBuilder interface {
	BuildCondChildStackItems(
		condProgramItem datamodels.ProgramItem,
		programStack datamodels.ProgramItemStack,
	)
}

type CondFunctionEvaluator interface {
	EvaluateCondProgramItem(
		top datamodels.ProgramItem,
		programStack datamodels.ProgramItemStack,
	) error
}

type FinishedProgramItemEvaluator interface {
	EvaluateFinishedProgramItem(
		top datamodels.ProgramItem,
		userDefinedFunctions map[string]datamodels.DefunFunction,
		evalStack datamodels.NodeStack,
		programStack datamodels.ProgramItemStack,
	) error
}


//counterfeiter:generate -o fakes/fake_post_evaluation_stack_updater.go . PostEvaluationStackUpdater
type PostEvaluationStackUpdater interface {
	UpdateStacksAfterEvaluation(
		evaluatedNode datamodels.Node,
		variableMap map[string] datamodels.Node,
		evalStack datamodels.NodeStack,
		programStack datamodels.ProgramItemStack,
	)
}

type QuoteFunctionEvaluator interface{
	EvaluateQuoteFunction(
		top datamodels.ProgramItem,
		evalStack datamodels.NodeStack,
		programStack datamodels.ProgramItemStack,
	)
}

type RootNodeEvaluator interface {
	Evaluate(
		rootNode datamodels.ExpressionListNode,
		userDefinedFunctions map[string]datamodels.DefunFunction,
		programStack datamodels.ProgramItemStack,
		evalStack datamodels.NodeStack,
	) (datamodels.Node, error)
}

//counterfeiter:generate -o fakes/fake_stack_updater_determiner.go . StackUpdateDeterminer
type StackUpdateDeterminer interface {
	DetermineHowToUpdateStacks(
		node datamodels.Node,
		variableMap map[string]datamodels.Node,
		evalStack datamodels.NodeStack,
		programStack datamodels.ProgramItemStack,
	)
}

//counterfeiter:generate -o fakes/fake_top_program_item_creator.go . TopProgramItemCreator
type TopProgramItemCreator interface {
	CreateTopProgramItem(
		expressionListNode datamodels.ExpressionListNode,
		variableMap map[string]datamodels.Node,
		programStack datamodels.ProgramItemStack,
	)
}

//counterfeiter:generate -o fakes/fake_top_program_item_updater.go . TopProgramItemUpdater
type TopProgramItemUpdater interface {
	UpdateTopProgramItemToNextChild(
		programStack datamodels.ProgramItemStack,
	)
}

type UnfinishedProgramStackItemEvaluator interface {
	EvaluateUnfinishedProgramItem(
		top datamodels.ProgramItem,
		evalStack datamodels.NodeStack,
		programStack datamodels.ProgramItemStack,
	)
}

//counterfeiter:generate -o fakes/fake_user_defined_function_evaluator.go . UserDefinedFunctionEvaluator
type UserDefinedFunctionEvaluator interface {
	EvaluateUserDefinedFunction(
		userDefinedFunction datamodels.DefunFunction,
		variableMap map[string] datamodels.Node,
		functionStack datamodels.NodeStack,
		evalStack datamodels.NodeStack,
		programStack datamodels.ProgramItemStack,
	)
}