package impl

import (
	"lisp_lambda-go/internal/evaluator"
	"lisp_lambda-go/internal/function"
)

type singleton struct {
	programEvaluator evaluator.ProgramEvaluator
}

func NewSingleton(
	functionMap map[string]function.Function,
) *singleton {
	atomRootNodeAsserter := NewAtomRootNodeAsserter()
	topProgramItemCreator := NewTopProgramItemCreator()
	topProgramItemUpdater := NewTopProgramItemUpdater()
	condChildStackItemBuilder := NewCondChildStackItemBuilder(
		topProgramItemCreator,
	)
	condFunctionEvaluator := NewCondFunctionEvaluator(
		topProgramItemUpdater,
		condChildStackItemBuilder,
	)
	postEvaluationUpdater := NewPostEvaluationStackUpdater(topProgramItemUpdater)
	stackUpdateDeterminer := NewStackUpdateDeterminer(
		topProgramItemCreator,
		postEvaluationUpdater,
	)
	condChildFunctionEvaluator := NewCondChildFunctionEvaluator(
		stackUpdateDeterminer,
	)
	quoteFunctionEvaluator := NewQuoteFunctionEvaluator(
		postEvaluationUpdater,
	)
	unfinishedProgramItemEvaluator := NewUnfinishedProgramItemEvaluator(
		stackUpdateDeterminer,
	)
	builtInFunctionEvaluator := NewBuiltInFunctionEvaluator(
		functionMap,
		postEvaluationUpdater,
	)
	userDefinedFunctionEvaluator := NewUserDefinedFunctionEvaluator(
		stackUpdateDeterminer,
	)
	finishedProgramItemEvaluator := NewFinishedProgramItemEvaluator(
		postEvaluationUpdater,
		functionMap,
		builtInFunctionEvaluator,
		userDefinedFunctionEvaluator,
	)

	rootNodeEvaluator := NewRootNodeEvaluator(
		topProgramItemCreator,
		condFunctionEvaluator,
		condChildFunctionEvaluator,
		quoteFunctionEvaluator,
		unfinishedProgramItemEvaluator,
		finishedProgramItemEvaluator,
	)

	programEvaluator := NewProgramEvaluator(
		atomRootNodeAsserter,
		rootNodeEvaluator,
	)

	return &singleton{programEvaluator}
}

func (s *singleton) GetProgramEvaluator() evaluator.ProgramEvaluator {
	return s.programEvaluator
}
