package impl

import (
	"lisp_lambda-go/internal/asserter"
	"lisp_lambda-go/internal/evaluator"
	interpreter2 "lisp_lambda-go/internal/interpreter"
	parser2 "lisp_lambda-go/internal/parser"
	"lisp_lambda-go/internal/printer"
	tokenizer2 "lisp_lambda-go/internal/tokenizer"
	"lisp_lambda-go/internal/userdefined"
)

type singleton struct {
	interpreter interpreter2.Interpreter
}

func NewSingleton(
	tokenizer tokenizer2.Tokenizer,
	parser parser2.Parser,
	functionGenerator userdefined.FunctionGenerator,
	expressionListLengthAsserter asserter.ExpressionListLengthAsserter,
	programEvaluator evaluator.ProgramEvaluator,
	listNotationPrinter printer.ListNotationPrinter,
) *singleton {
	rootNodePartitioner := NewRootNodePartitioner()
	interpreter := NewInterpreter(
		tokenizer,
		parser,
		rootNodePartitioner,
		functionGenerator,
		expressionListLengthAsserter,
		programEvaluator,
		listNotationPrinter,
	)

	return &singleton{interpreter}
}

func (s *singleton) GetInterpreter() interpreter2.Interpreter {
	return s.interpreter
}
