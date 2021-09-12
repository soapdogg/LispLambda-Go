package internal

import (
	"lisp_lambda-go/internal/asserter"
	"lisp_lambda-go/internal/core/constants"
	"lisp_lambda-go/internal/evaluator"
	"lisp_lambda-go/internal/function"
	"lisp_lambda-go/internal/interpreter"
	parser2 "lisp_lambda-go/internal/parser"
	"lisp_lambda-go/internal/printer"
	tokenizer2 "lisp_lambda-go/internal/tokenizer"
	"lisp_lambda-go/internal/userdefined"
)

type Singleton struct{
	interpreter Interpreter
}

func (s *Singleton) Init() {
	wordTokenizer := tokenizer2.NewWordTokenizer()
	tokenizer := tokenizer2.NewTokenizer(
		wordTokenizer,
	)

	nodeParser := parser2.NewNodeParser()
	parser := parser2.NewParser(
		nodeParser,
	)

	rootNodePartitioner := interpreter.NewRootNodePartitioner()

	functionLengthDeterminer := asserter.NewFunctionLengthDeterminer()
	functionLengthAsserter := asserter.NewFunctionLengthAsserter(
		functionLengthDeterminer,
	)

	atomFunction := function.NewAtomFunction()
	nullFunction := function.NewNullFunction()

	functionMap := map[string]Function {
		constants.ATOM: atomFunction,
		constants.NULL: nullFunction,
	}

	invalidNames := map[string]bool{}
	for key := range functionMap {
		invalidNames[key] = true
	}
	invalidNames[constants.T] = true
	invalidNames[constants.NIL] = true

	functionNameAsserter := userdefined.NewFunctionNameAsserter(
		invalidNames,
	)
	formalParameterAsserter := userdefined.NewUserDefinedFormalParametersAsserter(
		invalidNames,
	)
	functionGenerator := userdefined.NewFunctionGenerator(
		functionLengthAsserter,
		functionNameAsserter,
		formalParameterAsserter,
	)

	functionLengthMap := map[string]int {
		constants.ATOM: constants.TWO,
		constants.NULL: constants.TWO,
	}

	minimumFunctionLengthMap := map[string]int {

	}

	expressionListLengthAsserter := asserter.NewExpressionListLengthAsserter(
		functionLengthAsserter,
		functionLengthMap,
		minimumFunctionLengthMap,
	)

	atomRootNodeAsserter := evaluator.NewAtomRootNodeAsserter()
	topProgramItemCreator := evaluator.NewTopProgramItemCreator()
	topProgramItemUpdater := evaluator.NewTopProgramItemUpdater()
	condChildStackItemBuilder := evaluator.NewCondChildStackItemBuilder(
		topProgramItemCreator,
	)
	condFunctionEvaluator := evaluator.NewCondFunctionEvaluator(
		topProgramItemUpdater,
		condChildStackItemBuilder,
	)
	postEvaluationUpdater := evaluator.NewPostEvaluationStackUpdater(topProgramItemUpdater)
	stackUpdateDeterminer := evaluator.NewStackUpdateDeterminer(
		topProgramItemCreator,
		postEvaluationUpdater,
	)
	condChildFunctionEvaluator := evaluator.NewCondChildFunctionEvaluator(
		stackUpdateDeterminer,
	)
	quoteFunctionEvaluator := evaluator.NewQuoteFunctionEvaluator(
		postEvaluationUpdater,
	)
	unfinishedProgramItemEvaluator := evaluator.NewUnfinishedProgramItemEvaluator(
	 	stackUpdateDeterminer,
	)
	builtInFunctionEvaluator := evaluator.NewBuiltInFunctionEvaluator(
		functionMap,
		postEvaluationUpdater,
	)
	userDefinedFunctionEvaluator := evaluator.NewUserDefinedFunctionEvaluator(
		stackUpdateDeterminer,
	)
	finishedProgramItemEvaluator := evaluator.NewFinishedProgramItemEvaluator(
		postEvaluationUpdater,
		functionMap,
		builtInFunctionEvaluator,
		userDefinedFunctionEvaluator,
	)

	rootNodeEvaluator := evaluator.NewRootNodeEvaluator(
		topProgramItemCreator,
		condFunctionEvaluator,
		condChildFunctionEvaluator,
		quoteFunctionEvaluator,
		unfinishedProgramItemEvaluator,
		finishedProgramItemEvaluator,
	)

	programEvaluator := evaluator.NewProgramEvaluator(
		atomRootNodeAsserter,
		rootNodeEvaluator,
	)

	listNotationPrinter := printer.NewListNotationPrinter()

	s.interpreter = interpreter.NewInterpreter(
		tokenizer,
		parser,
		rootNodePartitioner,
		functionGenerator,
		expressionListLengthAsserter,
		programEvaluator,
		listNotationPrinter,
	)
}

func (s *Singleton) Get() Interpreter {
	return s.interpreter
}